package redis_cache

import (
	"context"
	"fmt"
	redisClient "github.com/redis/go-redis/v9"
)

type RedisCache struct {
	redis *redisClient.Client
}

func New(ctx context.Context, addr, pass string, DBNum int) (*RedisCache, error) {
	rdb := redisClient.NewClient(&redisClient.Options{
		Addr:     addr,
		Password: pass,
		DB:       DBNum,
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("redis connect err: %w", err)
	}

	return &RedisCache{redis: rdb}, nil
}
