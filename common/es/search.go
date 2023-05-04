package es

import (
	"encoding/json"
	"fmt"
	"io"
)

func (cli *EsClient[T]) SearchDoc(index string, body io.Reader) (*SearchResult[T], error) {
	resp, err := cli.Client.Search(cli.Client.Search.WithIndex(index), cli.Client.Search.WithBody(body))
	if err != nil {
		return nil, err
	}

	result, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode > 400 {
		return nil, fmt.Errorf("es状态码不为200, error: %s", string(result))
	}

	res := SearchResult[T]{}
	err = json.Unmarshal(result, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
