package examples

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"testing"
)

func Download(uuid string, start, end int64, fileName string) (int64, error) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:10000/api/file/"+uuid, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("Authorization", authorization)
	if end == 0 {
		req.Header.Set("range", fmt.Sprintf("bytes=%d-", start))
	} else {
		req.Header.Set("range", fmt.Sprintf("bytes=%d-%d", start, end))
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	err = os.MkdirAll("./files", 0777)
	if err != nil {
		return 0, err
	}

	if start == 0 {
		_, err := os.Create("./files/" + fileName)
		if err != nil {
			return 0, err
		}
	}

	f, err := os.OpenFile("./files/"+fileName, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		return 0, err
	}

	n, err := io.Copy(f, resp.Body)

	fmt.Println(n)
	return n, err
}

// 分片下载
func TestDownload(t *testing.T) {
	uuid := "9254C1EB-454B-4187-A90C-341745C2EEC7"

	req, err := http.NewRequest(http.MethodHead, "http://localhost:10000/api/file/"+uuid, nil)
	req.Header.Set("Authorization", authorization)
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	size, err := strconv.ParseInt(resp.Header.Get("content-length"), 10, 64)
	if err != nil {
		panic(err)
	}
	fileName := resp.Header.Get("file-name")

	fmt.Printf("fileName: %s, size: %d\n", fileName, size)

	current := int64(0)
	for current < size {
		n, err := Download(uuid, current, current+1<<20, fileName)
		current += n
		if err != nil {
			panic(err)
		}
	}
}
