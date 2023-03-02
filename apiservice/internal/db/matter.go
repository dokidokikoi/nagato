package db

import (
	"context"
	"nagato/apiservice/internal/model"
	"time"

	db "github.com/dokidokikoi/go-common/db/base"
)

type IMatterStore interface {
	db.BasicCURD[model.Matter]
	Insert(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	DelCache(ctx context.Context, key string) error
	GetCache(ctx context.Context, key string) (string, error)
}
