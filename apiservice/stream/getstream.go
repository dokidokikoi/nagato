package stream

import (
	"fmt"
	"io"
	"net/http"
)

type GetStream struct {
	reader io.Reader
}

func (r *GetStream) Read(p []byte) (n int, err error) {
	return r.reader.Read(p)
}

func newGetStream(url string) (*GetStream, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("dataServer return http code %d", resp.StatusCode)
	}

	return &GetStream{
		resp.Body,
	}, nil
}

func NewGetStream(server, hash string) (*GetStream, error) {
	if server == "" || hash == "" {
		return nil, fmt.Errorf("invalid server %s matter %s", server, hash)
	}

	return newGetStream("http://" + server + "/objects/" + hash)
}
