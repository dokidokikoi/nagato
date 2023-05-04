package api

import (
	"context"
	"nagato/apiservice/internal/db/data"
	"nagato/apiservice/internal/model"

	myJwt "github.com/dokidokikoi/go-common/jwt"
)

func Auth(token string) (*model.User, error) {
	claims, err := myJwt.VerifyToken(token, "test")

	if err != nil {
		return nil, err
	}

	store, err := data.GetStoreDBFactory()
	if err != nil {
		return nil, err
	}
	user, err := store.Users().Get(context.TODO(), &model.User{Email: claims.Emial}, nil)
	if err != nil {
		return nil, err
	}

	return user, nil
}
