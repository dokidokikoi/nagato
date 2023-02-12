package db

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type IBridgeStore interface {
	db.BasicCURD[model.Bridge]
}
