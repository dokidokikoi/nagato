package matter

import (
	"context"
	"nagato/apiservice/internal/es"
)

func (s matterSrv) CreateLastestResource(ctx context.Context, name, hash string, size int64) error {
	return s.esCli.ResourceEs().CreateLastestResource(name, hash, size)
}

func (s matterSrv) GetResourceMate(ctx context.Context, name string, version int) (*es.Resource, error) {
	return s.esCli.ResourceEs().GetResourceMate(name, version)
}

func (s matterSrv) SearchResourceAllVersion(name string, from, size int) ([]*es.Resource, error) {
	return s.esCli.ResourceEs().SearchResourceAllVersion(name, from, size)
}
