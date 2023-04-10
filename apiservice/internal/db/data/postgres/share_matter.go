package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type shareMatters struct {
	db.PgModel[model.ShareMatter]
}

func newShareMatters(d *Store) *shareMatters {
	return &shareMatters{db.PgModel[model.ShareMatter]{DB: d.db}}
}
