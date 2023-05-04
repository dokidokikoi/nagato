package service

import (
	"context"
	commonEsModel "nagato/common/es/model"
	"nagato/searchservice/internal/db"
)

type IResourceService interface {
	Search(ctx context.Context, resourceReq commonEsModel.ResourceReq) ([]commonEsModel.Resource, int64, error)
}

type resourceSrv struct {
	store db.Store
}

func (r resourceSrv) Search(ctx context.Context, resourceReq commonEsModel.ResourceReq) ([]commonEsModel.Resource, int64, error) {
	res, err := r.store.Resource().SearchResource(0, resourceReq)
	if err != nil {
		return nil, 0, err
	}

	resource := make([]commonEsModel.Resource, len(res))
	for i := range res {
		resource[i] = res[i].Source
	}

	return resource, 0, nil
}

func newResourceSrv(store db.Store) IResourceService {
	return &resourceSrv{store: store}
}
