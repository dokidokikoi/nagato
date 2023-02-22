package postgres

import (
	"gorm.io/gorm"
)

type transaction struct {
	DB *gorm.DB
}

func (t *transaction) TransactionBegin() *Store {
	db := t.DB.Begin()
	return &Store{
		db: db,
	}
}

func (t *transaction) TransactionRollback() {
	t.DB.Rollback()
}

func (t *transaction) TransactionCommit() {
	t.DB.Commit()
}

func newTransaction(ds *Store) *transaction {
	return &transaction{
		DB: ds.db,
	}
}
