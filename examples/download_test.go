package examples

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func Download(uuid string, start, end int64) (int64, error) {
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
		_, err := os.Create("./files/test.png")
		if err != nil {
			return 0, err
		}
	}

	f, err := os.OpenFile("./files/test.png", os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		return 0, err
	}

	n, err := io.Copy(f, resp.Body)

	fmt.Println(n)
	return n, err
}

// 分片下载
func TestDownload(t *testing.T) {
	size := int64(1540065)
	current := int64(0)
	for current < size {
		n, err := Download("F2C2E0FE-61D5-423F-B71B-2DD77CB5A956", current, current+1<<20)
		current += n
		if err != nil {
			panic(err)
		}
	}
}
