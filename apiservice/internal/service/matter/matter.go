package matter

import (
	"context"
	"io"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/model"
)

type IMatterService interface {
	Upload(ctx context.Context, hash string, size int64, data io.Reader) error
	Download(ctx context.Context, hash string) (io.Reader, error)
	Create(ctx context.Context, example *model.Matter) error

	CreateLastestResource(ctx context.Context, name, hash string, size int64) error
	GetResourceMate(ctx context.Context, name string, version int) (*Resource, error)
	SearchResourceAllVersion(ctx context.Context, name string, from, size int) ([]*Resource, error)
}

type matterSrv struct {
	store db.Store
}

func (s matterSrv) Create(ctx context.Context, example *model.Matter) error {
	return s.store.Matters().Create(ctx, example, nil)
}

func NewMatterService(store db.Store) IMatterService {
	return &matterSrv{store: store}
}
