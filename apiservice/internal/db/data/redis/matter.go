package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type matters struct {
	cli *redis.Client
}

func (m matters) Insert(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return m.cli.Set(ctx, key, value, expiration).Err()
}

func (m matters) Del(ctx context.Context, key string) error {
	return m.cli.Del(ctx, key).Err()
}

func (m matters) Get(ctx context.Context, key string) (string, error) {
	return m.cli.Get(ctx, key).Result()
}

func newMatters(cli *redis.Client) *matters {
	return &matters{cli: cli}
}
