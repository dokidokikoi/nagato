package share

import (
	"context"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/model"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type IShareService interface {
	Create(ctx context.Context, example *model.Share) error
	Update(ctx context.Context, example *model.Share) error
	Save(ctx context.Context, example *model.Share) error
	Del(ctx context.Context, example *model.Share) error
	Get(ctx context.Context, example *model.Share, option *meta.GetOption) (*model.Share, error)
}

func (s shareSrv) Create(ctx context.Context, example *model.Share) error {
	return s.store.Shares().Create(ctx, example, nil)
}

func (s shareSrv) Update(ctx context.Context, example *model.Share) error {
	return s.store.Shares().Update(ctx, example, nil)
}

func (s shareSrv) Save(ctx context.Context, example *model.Share) error {
	return s.store.Shares().Save(ctx, example, nil)
}

func (s shareSrv) Del(ctx context.Context, example *model.Share) error {
	return s.store.Shares().Delete(ctx, example, nil)
}

func (s shareSrv) Get(ctx context.Context, example *model.Share, option *meta.GetOption) (*model.Share, error) {
	return s.store.Shares().Get(ctx, example, option)
}

type shareSrv struct {
	store db.Store
}

func NewShareSrv(store db.Store) IShareService {
	return &shareSrv{store: store}
}
