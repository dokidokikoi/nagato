package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type tags struct {
	db.PgModel[model.Tag]
}

func newTags(d *Store) *tags {
	return &tags{db.PgModel[model.Tag]{DB: d.db}}
}
