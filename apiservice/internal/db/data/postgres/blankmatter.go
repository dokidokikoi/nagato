package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type blankMatters struct {
	db.PgModel[model.BlankMatter]
}

func newBlankMatters(d *Store) *blankMatters {
	return &blankMatters{db.PgModel[model.BlankMatter]{DB: d.db}}
}
