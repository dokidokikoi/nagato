package postgres

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type installs struct {
	db.PgModel[model.Install]
}

func newInstalls(d *Store) *installs {
	return &installs{db.PgModel[model.Install]{DB: d.db}}
}
