package tag

import (
	"context"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/model"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type ITagService interface {
	Create(ctx context.Context, example *model.Tag) error
	CreateCollection(ctx context.Context, examples []*model.Tag) []error
	List(ctx context.Context, example *model.Tag, option *meta.ListOption) ([]*model.Tag, int64, error)
	Del(ctx context.Context, example *model.Tag) error
}

type tagSrv struct {
	store db.Store
}

func (t tagSrv) Create(ctx context.Context, example *model.Tag) error {
	return t.store.Tags().Create(ctx, example, nil)
}

func (t tagSrv) CreateCollection(ctx context.Context, examples []*model.Tag) []error {
	return t.store.Tags().CreateCollection(ctx, examples, nil)
}

func (t tagSrv) List(ctx context.Context, example *model.Tag, option *meta.ListOption) ([]*model.Tag, int64, error) {
	res, err := t.store.Tags().List(ctx, example, option)
	if err != nil {
		return nil, 0, err
	}
	total, _ := t.store.Tags().Count(ctx, example, &option.GetOption)
	return res, total, nil
}

func (t tagSrv) Del(ctx context.Context, example *model.Tag) error {
	return t.store.Tags().Delete(ctx, example, nil)
}

func NewTagSrv(store db.Store) ITagService {
	return &tagSrv{store: store}
}
