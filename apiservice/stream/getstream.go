package stream

import (
	"fmt"
	"io"
	"net/http"
)

type GetStream struct {
	reader io.Reader
}

func newGetStream(url string) (*GetStream, error) {
	r, e := http.Get(url)
	if e != nil {
		return nil, e
	}
	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dataServer return http code %d", r.StatusCode)
	}

	return &GetStream{
		r.Body,
	}, nil
}

func NewGetStream(server, hashEncode string) (*GetStream, error) {
	if server == "" || hashEncode == "" {
		return nil, fmt.Errorf("invalid server %s object %s", server, hashEncode)
	}

	return newGetStream("http://" + server + "/data/file/" + hashEncode)
}

func (r *GetStream) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}
