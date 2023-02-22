package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type users struct {
	db.PgModel[model.User]
}

func newUsers(d *Store) *users {
	return &users{db.PgModel[model.User]{DB: d.db}}
}
