package service

import (
	"context"
	"time"
)

type Service struct {
	Cache Cache
}

func NewService(cache Cache) *Service {
	return &Service{
		Cache: cache,
	}
}

type Cache interface {
	Get(ctx context.Context, in string, dest interface{}) error
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
}
