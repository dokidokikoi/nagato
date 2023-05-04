package api

import (
	"context"
	"nagato/apiservice/internal/model"
	"nagato/apiservice/internal/service"

	meta "github.com/dokidokikoi/go-common/meta/option"
)

type User = model.User

func GetUserByEmail(ctx context.Context, email string, option *meta.GetOption) (*User, error) {
	return service.NewService().User().Get(ctx, &model.User{Email: email}, option)
}
