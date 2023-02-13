package es

import "io"

func (cli *EsClient[T]) CreateIndex(index string, body io.Reader) error {
	_, err := cli.Client.Indices.Create(index, cli.Client.Indices.Create.WithBody(body))
	return err
}

func (cli *EsClient[T]) DelIndices(indices ...string) error {
	_, err := cli.Client.Indices.Delete(indices)
	return err
}
