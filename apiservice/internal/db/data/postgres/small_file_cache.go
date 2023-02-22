package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type smallFileCaches struct {
	db.PgModel[model.SmallFileCache]
}

func newSmallFileCaches(d *Store) *smallFileCaches {
	return &smallFileCaches{db.PgModel[model.SmallFileCache]{DB: d.db}}
}
