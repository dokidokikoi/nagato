package db

import (
	"io"
	"nagato/apiservice/internal/model"
	commonEs "nagato/common/es"
	commonEsModel "nagato/common/es/model"

	db "github.com/dokidokikoi/go-common/db/base"
)

type IBlankStore interface {
	db.BasicCURD[model.Blank]
	CreateIndices(userID uint, indexReq string) error
	SearchBlank(userID uint, req commonEsModel.BlankReq) ([]commonEs.Result[commonEsModel.Blank], error)
	CreateDocWithID(userID uint, id string, req commonEsModel.Blank) error
	UpdateDoc(userID uint, id string, body io.Reader) error
	DelDoc(userID uint, id string) error
}
