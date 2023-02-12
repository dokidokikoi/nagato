package db

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type IShareStore interface {
	db.BasicCURD[model.Share]
}
