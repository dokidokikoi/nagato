package matter

import (
	"context"
	"io"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/es"
	"nagato/apiservice/internal/model"
)

type IMatterService interface {
	Upload(ctx context.Context, hash string, size int64, data io.Reader) error
	Download(ctx context.Context, hash string) (io.Reader, error)
	CreateLastestResource(ctx context.Context, name, hash string, size int64) error
	GetResourceMate(ctx context.Context, name string, version int) (*es.Resource, error)
	SearchResourceAllVersion(name string, from, size int) ([]*es.Resource, error)

	Create(ctx context.Context, example *model.Matter) error
}

type matterSrv struct {
	esCli es.IEsStore
	store db.Store
}

func (s matterSrv) Create(ctx context.Context, example *model.Matter) error {
	return s.store.Matters().Create(ctx, example, nil)
}

func NewMatterService(cli es.IEsStore, store db.Store) IMatterService {
	return &matterSrv{esCli: cli, store: store}
}
