package es

import (
	"github.com/elastic/go-elasticsearch/v8"
)

type Metadata struct {
	Name    string
	Version int
	Size    int64
	Hash    string
}

type EsClient[T any] struct {
	Client *elasticsearch.Client
}

func NewEsClient(username, password string, cert []byte, esServers ...string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: esServers,
		Username:  username,
		Password:  password,
		CACert:    cert,
	}
	cli, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return cli, nil
}
