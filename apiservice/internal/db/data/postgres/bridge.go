package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type bridges struct {
	db.PgModel[model.Bridge]
}

func newBridges(d *Store) *bridges {
	return &bridges{db.PgModel[model.Bridge]{DB: d.DB}}
}
