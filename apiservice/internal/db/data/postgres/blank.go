package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type blanks struct {
	db.PgModel[model.Blank]
}

func newBlanks(d *Store) *blanks {
	return &blanks{db.PgModel[model.Blank]{DB: d.db}}
}
