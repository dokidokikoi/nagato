package es

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	commonErrors "nagato/common/errors"
)

func (cli *EsClient[T]) GetDoc(index, docId string) (*T, error) {
	resp, err := cli.Client.GetSource(index, docId)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("es状态码不为200, code: %d", resp.StatusCode)
	}

	result, _ := io.ReadAll(resp.Body)

	res := new(T)
	err = json.Unmarshal(result, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (cli *EsClient[T]) CreateDocByID(index, id string, body io.Reader) error {
	resp, err := cli.Client.Create(index, id, body)
	if resp.StatusCode != 201 {
		return commonErrors.ErrESCreateDoc
	}
	return err
}

func (cli *EsClient[T]) UpdateDoc(index, id string, body io.Reader) error {
	_, err := cli.Client.Index(index, body, cli.Client.Index.WithDocumentID(id))
	return err
}

func (cli *EsClient[T]) DelDoc(index, id string) error {
	_, err := cli.Client.Get(index, id)
	return err
}

func (cli *EsClient[T]) BulkDoc(index string, body io.Reader) error {
	_, err := cli.Client.Bulk(body, cli.Client.Bulk.WithIndex(index))
	return err
}
