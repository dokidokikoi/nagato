package matter

import (
	"context"
	"io"
)

type IMatterService interface {
	CreateTempFile(ctx context.Context, hashEncode string, uuid string, size int64) error
	WriteTempFile(ctx context.Context, uuid string, data io.Reader) error
	CommitMatter(ctx context.Context, uuid, hash string) error
	DelMatterTemp(ctx context.Context, uuid string)
	CheckTempFileHash(ctx context.Context, uuid string, hash string, length int64) error

	GetFilePath(ctx context.Context, hash string) (string, error)
}

type matterSrv struct {
}

func NewMatterService() IMatterService {
	return &matterSrv{}
}
