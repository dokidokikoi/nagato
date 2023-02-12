package db

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type ITagStore interface {
	db.BasicCURD[model.Tag]
}
