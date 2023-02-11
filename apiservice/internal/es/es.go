package es

import (
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
	cert, _ := os.ReadFile("http_ca.crt")
	cli, err := commonEs.NewEsClient("elastic", "ADR*piFezssmbUhhN8*S", cert, "https://localhost:9200")
	if err != nil {
		panic(err)
	}

	return &esStore{
		cli: cli,
	}
}
