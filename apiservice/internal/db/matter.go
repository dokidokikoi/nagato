package db

import (
	"context"
	"io"
	"nagato/apiservice/internal/model"
	commonEs "nagato/common/es"
	commonEsModel "nagato/common/es/model"
	"time"

	db "github.com/dokidokikoi/go-common/db/base"
)

type IMatterStore interface {
	db.BasicCURD[model.Matter]
	Insert(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	DelCache(ctx context.Context, key string) error
	GetCache(ctx context.Context, key string) (string, error)

	CreateIndices(userID uint, indexReq string) error
	SearchResource(userID uint, req commonEsModel.ResourceReq) ([]commonEs.Result[commonEsModel.Resource], error)
	CreateDocWithID(userID uint, id string, req commonEsModel.Resource) error
	UpdateDoc(userID uint, id string, body io.Reader) error
	DelDoc(userID uint, id string) error
}
