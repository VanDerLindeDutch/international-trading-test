package redis_cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func (c *RedisCache) Set(ctx context.Context, key string, in interface{}, ttl time.Duration) error {
	marshal, err := json.Marshal(in)
	if err != nil {
		return err
	}
	err = c.redis.Set(ctx, key, marshal, ttl).Err()
	if err != nil {
		return fmt.Errorf("set redis token error: %w", err)
	}
	return nil
}
