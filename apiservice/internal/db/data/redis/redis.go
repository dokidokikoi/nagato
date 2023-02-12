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

func GetRedisFactory() (*Store, error) {
	once.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.RedisConfig.Host, config.RedisConfig.Port),
			Password: config.RedisConfig.Password,
			DB:       config.RedisConfig.DB,
		})
		redisFactory = &Store{cli: client}

	})
	return redisFactory, nil
}
