package db

type ITransaction interface {
	TransactionBegin() Store
	TransactionRollback()
	TransactionCommit()
}
