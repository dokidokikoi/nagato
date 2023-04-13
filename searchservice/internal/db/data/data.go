package data

import (
	"nagato/searchservice/internal/db"
	"nagato/searchservice/internal/db/data/es"
	"nagato/searchservice/internal/db/data/redis"
	"sync"
)

var (
	once         sync.Once
	datacFactory *dataCenter
)

type dataCenter struct {
	esCli *es.Store
	cli   *redis.Store
}

func (d dataCenter) Blank() db.IBlankStore {
	return newBlanks(d)
}

func (d dataCenter) Resource() db.IResourceStore {
	return newResources(d)
}

func GetDataCenter() *dataCenter {
	once.Do(func() {
		esCli, err := es.GetEsFactory()
		if err != nil {
			panic(err)
		}

		redisCli, err := redis.GetRedisFactory()
		if err != nil {
			panic(err)
		}

		datacFactory = &dataCenter{
			esCli: esCli,
			cli:   redisCli,
		}
	})

	return datacFactory
}
