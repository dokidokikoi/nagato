package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type Matters struct {
	db.PgModel[model.Matter]
}

func newMatters(d *Store) *Matters {
	return &Matters{db.PgModel[model.Matter]{DB: d.db}}
}
