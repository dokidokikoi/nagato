package matter

import (
	"context"
	"io"
	"nagato/apiservice/internal/es"
)

type IMatterService interface {
	Upload(ctx context.Context, hash string, size int64, data io.Reader) error
	Download(ctx context.Context, hash string) (io.Reader, error)
	CreateLastestResource(ctx context.Context, name, hash string, size int64) error
	GetResourceMate(ctx context.Context, name string, version int) (*es.Resource, error)
	SearchResourceAllVersion(name string, from, size int) ([]*es.Resource, error)
}

type matterSrv struct {
	esCli es.IEsStore
}

func NewMatterService(cli es.IEsStore) IMatterService {
	return &matterSrv{esCli: cli}
}
