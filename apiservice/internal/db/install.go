package db

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type IInstallStore interface {
	db.BasicCURD[model.Install]
}
