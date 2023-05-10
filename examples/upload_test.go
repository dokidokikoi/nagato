package examples

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
)

var authorization = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiIxMjNAZXhhbXBsZS5jb20iLCJyb2xlIjoiIiwiZXhwIjoxNzE0NDU3NjE2LCJpYXQiOjE2ODMwMDgwMTYsImlzcyI6ImhhcnVrYXplIiwibmJmIjoxNjgzMDA4MDE2fQ.c1VZqJEtNSKeekzfi1WSd113UobIhtdvvlKn7KYG4kg"

type response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func getUploadToken(filePath string) (*response, error) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileInfo, err := f.Stat()
	if err != nil {
		panic(err)
	}
	h := sha256.New()
	io.Copy(h, f)
	hash := base64.StdEncoding.EncodeToString(h.Sum(nil))
	body := map[string]interface{}{
		"name":    fileInfo.Name(),
		"sha256":  hash,
		"size":    fileInfo.Size(),
		"privacy": false,
		"path":    "/",
	}
	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:10000/api/file", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Authorization", authorization)
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	buff, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	res := &response{}
	json.Unmarshal(buff, res)

	return res, nil
}

func Upload(token string, filePath string, seek int64, cacheSize int) error {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	buff := make([]byte, cacheSize)
	current := 0

	if seek != 0 {
		buf := make([]byte, seek)
		io.ReadFull(f, buf)
	}
	for {
		n, err := io.ReadFull(f, buff)
		if err != nil {
			if err == io.EOF {
				fmt.Println(current)
				return nil
			}
			if err != io.ErrUnexpectedEOF {
				return err
			}
		}

		buff = buff[:n]

		req, err := http.NewRequest(http.MethodPut, "http://localhost:10000/api/file/temp/"+token, bytes.NewBuffer(buff))
		req.Header.Set("Authorization", authorization)
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", (int64(current)+seek)))
		if err != nil {
			return err
		}
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		resBytes, _ := io.ReadAll(resp.Body)
		fmt.Println(string(resBytes))

		current += n
	}

}

// 分片上传
func TestUpload(t *testing.T) {
	res, err := getUploadToken("./logo.png")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", res)

	req, err := http.NewRequest(http.MethodHead, "http://localhost:10000/api/file/temp/"+res.Data, nil)
	req.Header.Set("Authorization", authorization)
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	length, err := strconv.ParseInt(resp.Header.Get("content-length"), 10, 64)
	if err != nil {
		panic(err)
	}
	perSize, err := strconv.Atoi(resp.Header.Get("per-size"))
	if err != nil {
		panic(err)
	}

	err = Upload(res.Data, "./logo.png", length, perSize)
	fmt.Printf("%+v", err)
}
