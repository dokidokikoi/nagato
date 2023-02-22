package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type shares struct {
	db.PgModel[model.Share]
}

func newShares(d *Store) *shares {
	return &shares{db.PgModel[model.Share]{DB: d.db}}
}
