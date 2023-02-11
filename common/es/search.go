package es

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (cli *EsClient[T]) SearchDoc(index string, body io.Reader) (*SearchResult[T], error) {
	resp, err := cli.Client.Search(cli.Client.Search.WithIndex(index), cli.Client.Search.WithBody(body))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("es状态码不为200, code: %d", resp.StatusCode)
	}

	result, _ := io.ReadAll(resp.Body)

	res := SearchResult[T]{}
	err = json.Unmarshal(result, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
