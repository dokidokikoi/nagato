package db

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type IBlankMatterStore interface {
	db.BasicCURD[model.BlankMatter]
}
