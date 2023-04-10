package db

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type IShareMatterStore interface {
	db.BasicCURD[model.ShareMatter]
}
