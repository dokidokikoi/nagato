package matter

import (
	"context"
	"io"
)

type IMatterService interface {
	Upload(ctx context.Context, hash string, size int64, data io.Reader) error
	Download(ctx context.Context, hash string) (io.Reader, error)
}

type matterSrv struct {
}

func NewMatterService() IMatterService {
	return &matterSrv{}
}
