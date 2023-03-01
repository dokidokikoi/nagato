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

var authorization = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiZW1haWwiOiIxMjNAZXhhbXBsZS5jb20iLCJyb2xlIjoiIiwiZXhwIjoxNjc3NzIyMzIwLCJpYXQiOjE2Nzc2MzU5MjAsImlzcyI6ImhhcnVrYXplIiwibmJmIjoxNjc3NjM1OTIwfQ.PEBiv37ersSPv4z3-_V6pzBHbVp-3-UwY42-XOTrOUA"

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

func Upload(token string, filePath string, seek int64) error {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	size := 24000
	buff := make([]byte, size)
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

func TestUpload(t *testing.T) {
	res, err := getUploadToken("./logo.png")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", res)
	err = Upload(res.Data, "./logo.png", 0)
	println(err.Error())
}

func ResumeUpload(token string) error {
	req, err := http.NewRequest(http.MethodHead, "http://localhost:10000/api/file/temp/"+token, nil)
	req.Header.Set("Authorization", authorization)
	if err != nil {
		return err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	length, err := strconv.ParseInt(resp.Header.Get("content-length"), 10, 64)
	if err != nil {
		return err
	}

	fmt.Println(length)

	err = Upload(token, "./logo.png", length)

	return err
}

func TestResumeUpload(t *testing.T) {
	ResumeUpload("eyJOYW1lIjoibG9nbyIsIlNpemUiOjE1NDAwNjUsIkhhc2giOiJNeUZjSEE2SER5S01HNnR3TndrUTVGQUIvUVRrUjQzcXVyVDF2ZGF4K0FrPSIsIlNlcnZlcnMiOlsiOjEwMTUzIiwiOjEwMTUwIiwiOjEwMTUyIiwiOjEwMTUxIl0sIlV1aWRzIjpbIjE5ODczMTU5LTVBQkUtNDg0Mi04N0Q5LUI4NkU1NDE2RkEyMiIsIkNFMkM0QjY5LTRCMUQtNEQ5Mi04MkFGLUYxMDM5NEY3OTUwNiIsIjNDRUY3QzcxLTNDNTUtNEE1NS05RDAwLTQ0RDc3ODc4Q0EwNiIsIkQ0MTA3NDI1LTNDNjgtNDM1Ny1BNTVBLTAwMTRCQkNDQjJGOSJdfQ==")
}
