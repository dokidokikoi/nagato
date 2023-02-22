package data

import (
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/db/data/mongo"
	"nagato/apiservice/internal/db/data/postgres"
	"nagato/apiservice/internal/db/data/redis"
)

type transaction struct {
	mongo *mongo.Store
	pg    *postgres.Store
	redis *redis.Store
}

func (t *transaction) TransactionBegin() db.Store {
	return &dataCenter{mongo: t.mongo, redis: t.redis, pg: t.pg.Transaction().TransactionBegin()}
}

func (t *transaction) TransactionRollback() {
	t.pg.Transaction().TransactionRollback()
}

func (t *transaction) TransactionCommit() {
	t.pg.Transaction().TransactionCommit()
}

var _ db.ITransaction = (*transaction)(nil)

func newTransaction(center *dataCenter) *transaction {
	return &transaction{
		pg:    center.pg,
		redis: center.redis,
		mongo: center.mongo,
	}
}
