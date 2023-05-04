package service

import (
	"context"
	commonEsModel "nagato/common/es/model"
	"nagato/searchservice/internal/db"
)

type IBlankService interface {
	Search(ctx context.Context, blankReq commonEsModel.BlankReq, resourceReq commonEsModel.ResourceReq) ([]commonEsModel.Blank, int64, error)
}

type blankSrv struct {
	store db.Store
}

func (b blankSrv) Search(ctx context.Context, blankReq commonEsModel.BlankReq, resourceReq commonEsModel.ResourceReq) ([]commonEsModel.Blank, int64, error) {
	resourceReq.Select = []string{"id"}
	res, err := b.store.Resource().SearchResource(0, resourceReq)
	if err != nil {
		return nil, 0, err
	}

	ids := make([]uint, len(res))
	for i := range res {
		ids[i] = res[i].Source.ID
	}

	blankReq.MatterIDs = ids
	result, err := b.store.Blank().SearchBlank(0, blankReq)
	if err != nil {
		return nil, 0, err
	}

	blanks := make([]commonEsModel.Blank, len(result))
	for i := range result {
		blanks[i] = result[i].Source
	}

	return blanks, 0, nil
}

func newBlankSrv(store db.Store) IBlankService {
	return &blankSrv{store: store}
}
