package db

import (
	commonEs "nagato/common/es"
	commonEsModel "nagato/common/es/model"
)

type IBlankStore interface {
	CreateIndices(userID uint, indexReq string) error
	SearchBlank(userID uint, req commonEsModel.BlankReq) ([]commonEs.Result[commonEsModel.Blank], error)
	CreateDocWithID(userID uint, id string, req commonEsModel.Blank) error
}
