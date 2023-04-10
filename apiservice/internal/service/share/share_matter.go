package share

import (
	"context"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/model"
)

type IShareMatterService interface {
	Create(ctx context.Context, example *model.ShareMatter) error
}

type shareMatterSrv struct {
	store db.Store
}

func (s shareMatterSrv) Create(ctx context.Context, example *model.ShareMatter) error {
	return s.store.ShareMatters().Create(ctx, example, nil)
}

func NewShareMatterSrv(store db.Store) IShareMatterService {
	return &shareMatterSrv{store: store}
}
