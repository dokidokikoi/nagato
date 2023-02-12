package db

import (
	"nagato/apiservice/internal/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type IUserStore interface {
	db.BasicCURD[model.User]
}
