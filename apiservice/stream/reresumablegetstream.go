package stream

import (
	"context"
	"io"

	dataRpc "nagato/apiservice/rpc/client/data"
)

type RSResumableGetStream struct {
	*RSGetStream
}

func NewRSResumableGetStream(dataServers []string, uuids []string, size uint) (*RSResumableGetStream, error) {
	readers := make([]io.Reader, ALL_SHARDS)
	for i := range dataServers {
		dataRpc, err := dataRpc.GetDataClient(dataServers[i])
		if err != nil {
			return nil, err
		}
		r, err := dataRpc.GetTempFile(context.Background(), uuids[i])
		if err != nil {
			return nil, err
		}

		readers[i] = r
	}

	doc := NewDecoder(readers, nil, size)

	return &RSResumableGetStream{
		RSGetStream: &RSGetStream{doc},
	}, nil
}
