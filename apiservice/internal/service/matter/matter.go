package matter

import (
	"context"
	"io"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/model"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type IMatterService interface {
	Upload(ctx context.Context, example *model.Matter, hash string, size int64, data io.Reader) error
	Download(ctx context.Context, hash string) (io.Reader, error)

	Create(ctx context.Context, example *model.Matter) error
	Update(ctx context.Context, example *model.Matter) error
	Del(ctx context.Context, example *model.Matter, option *meta.DeleteOption) error
	UpdateByWhereNode(ctx context.Context, node *meta.WhereNode, example *model.Matter) error
	Get(ctx context.Context, example *model.Matter, option *meta.GetOption) (*model.Matter, error)

	CreateResource(ctx context.Context, example *model.Matter) error
	GetResourceMate(ctx context.Context, name string, version int) (*model.Matter, error)
	SearchResourceAllVersion(ctx context.Context, name string, from, size int) ([]*model.Matter, error)
}

type matterSrv struct {
	store db.Store
}

func (s matterSrv) Create(ctx context.Context, example *model.Matter) error {
	return s.store.Matters().Create(ctx, example, nil)
}

func (s matterSrv) Update(ctx context.Context, example *model.Matter) error {
	return s.store.Matters().Update(ctx, example, nil)
}

func (s matterSrv) Del(ctx context.Context, example *model.Matter, option *meta.DeleteOption) error {
	return s.store.Matters().Delete(ctx, example, option)
}

func (s matterSrv) UpdateByWhereNode(ctx context.Context, node *meta.WhereNode, example *model.Matter) error {
	return s.store.Matters().UpdateByWhere(ctx, node, example, nil)
}

func (s matterSrv) Get(ctx context.Context, example *model.Matter, option *meta.GetOption) (*model.Matter, error) {
	return s.store.Matters().Get(ctx, example, option)
}

func NewMatterService(store db.Store) IMatterService {
	return &matterSrv{store: store}
}
