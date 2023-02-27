package stream

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	dataRpc "nagato/apiservice/rpc/client/data"
)

type TempPutStream struct {
	Total  uint
	Size   uint
	Server string
	Uuid   string
}

func (t *TempPutStream) Write(p []byte) (n int, err error) {
	dataClient, err := dataRpc.GetDataClient(t.Server)
	if err != nil {
		return 0, err
	}
	err = dataClient.UploadTempFile(context.Background(), t.Uuid, strings.NewReader(string(p)))
	if err != nil {
		return 0, err
	}

	t.Total += uint(len(p))
	fmt.Printf("write: %d bytes, total: %d bytes, size: %d bytes\n", len(p), t.Total, t.Size)

	return len(p), nil
}

func (t *TempPutStream) Commit(flag bool, hash string) {
	dataClient, err := dataRpc.GetDataClient(t.Server)
	if err != nil {
		return
	}

	if flag {
		dataClient.CommitTempFile(context.Background(), t.Uuid, hash)
	} else {
		dataClient.DeleteTempFile(context.Background(), t.Uuid)
	}
}

func NewTempPutStream(server string, hash string, size uint) (*TempPutStream, error) {
	// 将文件信息存到临时目录下文件名为生成的uuid
	dataClient, err := dataRpc.GetDataClient(server)
	if err != nil {
		return nil, err
	}
	uuid, err := dataClient.CreateTempInfo(context.Background(), url.PathEscape(hash), int64(size))
	if err != nil {
		return nil, err
	}

	return &TempPutStream{
		Size:   size,
		Server: server,
		Uuid:   string(uuid),
	}, nil
}
