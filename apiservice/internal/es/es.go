package es

import (
	"nagato/apiservice/internal/config"
	commonEs "nagato/common/es"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
)

type IEsStore interface {
	ResourceEs() *resourceEs
}

type esStore struct {
	cli *elasticsearch.Client
}

func (s *esStore) ResourceEs() *resourceEs {
	return newResource(s)
}

func NewEsSore() *esStore {
	cert, _ := os.ReadFile(config.Config().EsConfig.Cert)
	cli, err := commonEs.NewEsClient(config.Config().EsConfig.Username, config.Config().EsConfig.Password, cert, config.Config().EsConfig.Url())
	if err != nil {
		panic(err)
	}

	return &esStore{
		cli: cli,
	}
}
