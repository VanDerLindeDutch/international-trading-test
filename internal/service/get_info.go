package service

import (
	"context"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"international-trading-test/internal/pkg/cache_errors"
	"time"
)

func (s *Service) GetInfoFromSomeClient(ctx context.Context, in string) (*HashedValue, error) {
	var value HashedValue
	if err := s.Cache.Get(ctx, in, &value); err != nil {
		if !errors.Is(err, cache_errors.NotFound) {
			return nil, err
		}
	} else {

		return &value, nil
	}
	hashedValue := s.hash(in)
	err := s.Cache.Set(ctx, in, hashedValue, time.Second*10)
	if err != nil {
		return nil, err
	}
	return hashedValue, nil
}

func (s *Service) hash(in string) *HashedValue {
	//emulate work
	time.Sleep(time.Second * 2)
	mdHash := md5.New()
	shaHash := sha256.New()

	return &HashedValue{
		SHA256value: hex.EncodeToString(shaHash.Sum([]byte(in))),
		MD5value:    hex.EncodeToString(mdHash.Sum([]byte(in))),
	}
}
