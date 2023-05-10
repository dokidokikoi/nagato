package es

import (
	"fmt"
	"io"
	commonErrors "nagato/common/errors"
)

func (cli *EsClient[T]) CreateIndex(index string, body io.Reader) error {
	res, err := cli.Client.Indices.Exists([]string{index})
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// 索引已存在，直接返回nil
	if res.StatusCode != 404 && res.IsError() {
		return fmt.Errorf("index check error: %s", res.Status())
	}

	if res.StatusCode == 200 {
		return nil
	}

	// 建立索引
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
