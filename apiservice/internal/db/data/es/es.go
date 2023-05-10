package es

import (
	"errors"
	"nagato/apiservice/internal/config"
	commonEs "nagato/common/es"
	"os"
	"sync"

	"github.com/elastic/go-elasticsearch/v8"
)

type Store struct {
	cli *elasticsearch.Client
}

var (
	esFactory *Store
	once      sync.Once
)

func (s Store) Blanks() *blanks {
	return newBlanks(s.cli)
}

func (s Store) Resources() *resources {
	return newResources(s.cli)
}

func GetEsFactory() (*Store, error) {
	once.Do(func() {
		cert, _ := os.ReadFile(config.Config().EsConfig.Cert)
		cli, err := commonEs.NewEsClient(config.Config().EsConfig.Username, config.Config().EsConfig.Password, cert, config.Config().EsConfig.Url())
		if err != nil {
			panic(err)
		}
		esFactory = &Store{cli: cli}
	})
	return esFactory, nil
}

func GetEsClient() (*elasticsearch.Client, error) {
	if esFactory == nil {
		return nil, errors.New("数据层未初始化")
	}
	return esFactory.cli, nil
}
