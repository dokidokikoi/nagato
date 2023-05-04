package db

import (
	commonEs "nagato/common/es"
	commonEsModel "nagato/common/es/model"
)

type IResourceStore interface {
	CreateIndices(userID uint, indexReq string) error
	SearchResource(userID uint, req commonEsModel.ResourceReq) ([]commonEs.Result[commonEsModel.Resource], error)
	CreateDocWithID(userID uint, id string, req commonEsModel.Resource) error
}
