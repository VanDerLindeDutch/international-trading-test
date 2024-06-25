package redis_cache

import (
	"context"
	"encoding/json"
	"errors"
	redisClient "github.com/redis/go-redis/v9"
	"international-trading-test/internal/pkg/cache_errors"
)

// Get returns nil if value doesn't exist
func (c *RedisCache) Get(ctx context.Context, key string, dest interface{}) error {
	cmd := c.redis.Get(ctx, key)
	if cmd.Err() != nil {
		if errors.Is(cmd.Err(), redisClient.Nil) {
			return cache_errors.NotFound
		}
		return cmd.Err()
	}
	res, err := cmd.Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(res), dest)
}
