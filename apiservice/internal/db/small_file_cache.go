package db

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type ISmallFileCacheStore interface {
	db.BasicCURD[model.SmallFileCache]
}
