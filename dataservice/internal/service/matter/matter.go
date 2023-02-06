package matter

import (
	"context"
	"io"
)

type IMatterService interface {
	CreateTempFile(ctx context.Context, name string, uuid string, size int64) error
	WriteTempFile(ctx context.Context, uuid string, data io.Reader) error
	CommitMatter(ctx context.Context, uuid string) error
	DelMatterTemp(ctx context.Context, uuid string)

	GetFilePath(ctx context.Context, hash string) string
}

type matterSrv struct {
}

func NewMatterService() IMatterService {
	return &matterSrv{}
}
