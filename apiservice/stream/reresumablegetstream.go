package stream

import (
	"fmt"
	"io"
	"net/http"
)

type RSResumableGetStream struct {
	*RSGetStream
}

func NewRSResumableGetStream(dataServers []string, uuids []string, size uint) (*RSResumableGetStream, error) {
	readers := make([]io.Reader, ALL_SHARDS)
	for i := range dataServers {
		resp, err := http.Get(fmt.Sprintf("http://%s/data/file/temp/%s", dataServers[i], uuids[i]))
		if err != nil {
			return nil, err
		}

		readers[i] = resp.Body
	}

	doc := NewDecoder(readers, nil, size)

	return &RSResumableGetStream{
		RSGetStream: &RSGetStream{doc},
	}, nil
}
