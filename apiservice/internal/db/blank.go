package db

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type IBlankStore interface {
	db.BasicCURD[model.Blank]
}
