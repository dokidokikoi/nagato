package es

import (
	"io"
	commonErrors "nagato/common/errors"
)

func (cli *EsClient[T]) CreateIndex(index string, body io.Reader) error {
	resp, err := cli.Client.Indices.Create(index, cli.Client.Indices.Create.WithBody(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return commonErrors.ErrESCreateIndex
	}
	return nil
}

func (cli *EsClient[T]) DelIndices(indices ...string) error {
	_, err := cli.Client.Indices.Delete(indices)
	return err
}
