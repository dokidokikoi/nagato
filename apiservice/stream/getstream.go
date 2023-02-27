package stream

import (
	"context"
	"fmt"
	"io"

	dataRpc "nagato/apiservice/rpc/client/data"
)

type GetStream struct {
	reader io.Reader
}

func newGetStream(server string, namePrefix string) (*GetStream, error) {
	dataRpc, err := dataRpc.GetDataClient(server)
	if err != nil {
		return nil, err
	}
	r, err := dataRpc.GetMatter(context.Background(), namePrefix)
	if err != nil {
		return nil, err
	}

	return &GetStream{
		r,
	}, nil
}

func NewGetStream(server, hashEncode string) (*GetStream, error) {
	if server == "" || hashEncode == "" {
		return nil, fmt.Errorf("invalid server %s object %s", server, hashEncode)
	}

	return newGetStream(server, hashEncode)
}

func (r *GetStream) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}
