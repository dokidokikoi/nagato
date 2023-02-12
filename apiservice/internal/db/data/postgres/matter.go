package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type matters struct {
	db.PgModel[model.Matter]
}

func newMatters(d *Store) *matters {
	return &matters{db.PgModel[model.Matter]{DB: d.DB}}
}
