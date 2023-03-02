package redis

import (
	"fmt"
	"nagato/apiservice/internal/config"
	"sync"

	"github.com/redis/go-redis/v9"
)

type Store struct {
	cli *redis.Client
}

var (
	redisFactory *Store
	once         sync.Once
)

func (s Store) Matters() *matters {
	return newMatters(s.cli)
}

func GetRedisFactory() (*Store, error) {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Config().RedisConfig.Host, config.Config().RedisConfig.Port),
			Password: config.Config().RedisConfig.Password,
			DB:       config.Config().RedisConfig.DB,
		})
		redisFactory = &Store{cli: client}

	})
	return redisFactory, nil
}
