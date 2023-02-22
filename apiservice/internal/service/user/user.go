package user

import (
	"context"
	"nagato/apiservice/internal/db"
	"nagato/apiservice/internal/model"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type IUserService interface {
	Create(ctx context.Context, example *model.User) error
	Get(ctx context.Context, example *model.User, option *meta.GetOption) (*model.User, error)
}

type userSrv struct {
	store db.Store
}

func (u userSrv) Create(ctx context.Context, example *model.User) error {
	return u.store.Users().Create(ctx, example, nil)
}

func (u userSrv) Get(ctx context.Context, example *model.User, option *meta.GetOption) (*model.User, error) {
	return u.store.Users().Get(ctx, example, option)
}

func NewUserService(store db.Store) IUserService {
	return &userSrv{store: store}
}
