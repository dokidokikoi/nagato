package stream

import (
	"bytes"
	"context"
	"fmt"

	dataRpc "nagato/apiservice/rpc/client/data"
)

// 用于上传临时文件
type TempPutStream struct {
	Total  uint
	Size   uint
	Server string
	Uuid   string
}

// 实现 Writer 接口
// 将数据通过 grpc 写入数据服务的磁盘
func (t *TempPutStream) Write(p []byte) (n int, err error) {
	// 获取数据服务
	dataClient, err := dataRpc.GetDataClient(t.Server)
	if err != nil {
		return 0, err
	}
	err = dataClient.UploadTempFile(context.Background(), t.Uuid, bytes.NewBuffer(p))
	if err != nil {
		return 0, err
	}

	t.Total += uint(len(p))
	fmt.Printf("write: %d bytes, total: %d bytes, size: %d bytes\n", len(p), t.Total, t.Size)

	return len(p), nil
}

// 临时文件转正
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
	// 获取 uuid
	uuid, err := dataClient.CreateTempInfo(context.Background(), hash, int64(size))
	if err != nil {
		return nil, err
	}

	return &TempPutStream{
		Size:   size,
		Server: server,
		Uuid:   string(uuid),
	}, nil
}
