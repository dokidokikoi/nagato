package matter

import (
	"context"
	"fmt"
	"io"
	"nagato/apiservice/internal/locate"
	"nagato/apiservice/stream"
)

func (s matterSrv) Download(ctx context.Context, hash string) (io.Reader, error) {
	server := locate.Locate(hash)
	if server == "" {
		return nil, fmt.Errorf("matter %s locate failed", hash)
	}

	return stream.NewGetStream(server, hash)
}
