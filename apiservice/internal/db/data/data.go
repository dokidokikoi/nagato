package data

import (
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/db/data/mongo"
	"nagato/apiservice/internal/db/data/postgres"
	"nagato/apiservice/internal/db/data/redis"
	"sync"
)

var (
	datacFactory db.Store
	once         sync.Once
)

type dataCenter struct {
	pg    *postgres.Store
	redis *redis.Store
	mongo *mongo.Store
}

func (d dataCenter) Blanks() db.IBlankStore {
	return d.pg.Blanks()
}

func (d dataCenter) BlankMatters() db.IBlankMatterStore {
	return d.pg.BlankMatters()
}

func (d dataCenter) Installs() db.IInstallStore {
	return d.pg.Installs()
}

func (d dataCenter) Logs() db.ILogStore {
	return d.mongo.Logs()
}

func (d *dataCenter) Matters() db.IMatterStore {
	return newMatters(d)
}

func (d dataCenter) Shares() db.IShareStore {
	return d.pg.Shares()
}

func (d dataCenter) ShareMatters() db.IShareMatterStore {
	return d.pg.ShareMatters()
}

func (d dataCenter) SmallFileCaches() db.ISmallFileCacheStore {
	return d.pg.SmallFileCaches()
}

func (d dataCenter) Tags() db.ITagStore {
	return d.pg.Tags()
}

func (d dataCenter) Users() db.IUserStore {
	return d.pg.Users()
}

func (d *dataCenter) Transaction() db.ITransaction {
	return newTransaction(d)
}

func GetStoreDBFactory() (db.Store, error) {
	once.Do(func() {
		pg, err := postgres.GetPGFactory()
		if err != nil {
			panic(err)
		}

		redisCli, err := redis.GetRedisFactory()
		if err != nil {
			panic(err)
		}

		datacFactory = &dataCenter{
			pg:    pg,
			redis: redisCli,
		}
	})

	return datacFactory, nil
}

func SetStoreDBFactory() {
	var err error
	datacFactory, err = GetStoreDBFactory()
	if err != nil {
		return
	}
}
