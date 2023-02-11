package stream

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type TempPutStream struct {
	Server string
	Uuid   string
}

func (t TempPutStream) Write(p []byte) (n int, err error) {
	req, err := http.NewRequest("PATCH", "http://"+t.Server+"/file/temp/"+t.Uuid, strings.NewReader(string(p)))
	if err != nil {
		return 0, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("dataServer return http code %d", resp.StatusCode)
	}

	return len(p), nil
}

func (t *TempPutStream) Commit(flag bool) {
	method := "DELETE"
	if flag {
		method = "PUT"
	}

	req, _ := http.NewRequest(method, "http://"+t.Server+"/file/temp/"+t.Uuid, nil)
	client := http.Client{}
	client.Do(req)
}

func NewTempPutStream(server string, hash string, size int64) (*TempPutStream, error) {
	// 将文件信息存到临时目录下文件名为生成的uuid
	req, err := http.NewRequest("POST", "http://"+server+"/file/temp/"+hash, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("size", fmt.Sprintf("%d", size))
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// 根据uuid将临时文件转正和删除
	uuid, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &TempPutStream{
		Server: server,
		Uuid:   string(uuid),
	}, nil
}
