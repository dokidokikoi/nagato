package blank

import (
	"context"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/model"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type IBlankService interface {
	Create(ctx context.Context, example *model.Blank) error
	List(ctx context.Context, example *model.Blank, option *meta.ListOption) ([]*model.Blank, int64, error)
	Update(ctx context.Context, example *model.Blank) error
	UpdateBlank(ctx context.Context, example *model.Blank) error
	Save(ctx context.Context, example *model.Blank) error
	Del(ctx context.Context, example *model.Blank) error
	Get(ctx context.Context, example *model.Blank, option *meta.GetOption) (*model.Blank, error)
}

type blankSrv struct {
	store db.Store
}

func (b blankSrv) Create(ctx context.Context, example *model.Blank) error {
	return b.store.Blanks().Create(ctx, example, nil)
}

func (b blankSrv) List(ctx context.Context, example *model.Blank, option *meta.ListOption) ([]*model.Blank, int64, error) {
	res, err := b.store.Blanks().List(ctx, example, option)
	if err != nil {
		return nil, 0, err
	}
	total, _ := b.store.Blanks().Count(ctx, example, &option.GetOption)
	return res, total, nil
}

func (b blankSrv) Update(ctx context.Context, example *model.Blank) error {
	return b.store.Blanks().Update(ctx, example, nil)
}

func (b blankSrv) UpdateBlank(ctx context.Context, example *model.Blank) error {
	b.store.BlankMatters().Delete(ctx, &model.BlankMatter{BlankID: example.ID}, nil)
	return b.store.Blanks().Update(ctx, example, nil)
}

func (b blankSrv) Save(ctx context.Context, example *model.Blank) error {
	return b.store.Blanks().Save(ctx, example, nil)
}

func (b blankSrv) Del(ctx context.Context, example *model.Blank) error {
	return b.store.Blanks().Delete(ctx, example, nil)
}

func (b blankSrv) Get(ctx context.Context, example *model.Blank, option *meta.GetOption) (*model.Blank, error) {
	return b.store.Blanks().Get(ctx, example, option)
}

func NewBlankSrv(store db.Store) IBlankService {
	return &blankSrv{store: store}
}
