package cache

import (
	"context"
	"fmt"
	"time"

	"W-chat/pkg/encrypt"

	"github.com/redis/go-redis/v9"
)

type JwtTokenCache struct {
	redis *redis.Client
}

func NewTokenSessionCache(redis *redis.Client) *JwtTokenCache {
	return &JwtTokenCache{redis}
}

func (s *JwtTokenCache) SetBlackList(ctx context.Context, token string, exp time.Duration) error {
	return s.redis.Set(ctx, s.name(token), 1, exp).Err()
}

func (s *JwtTokenCache) IsBlackList(ctx context.Context, token string) bool {
	return s.redis.Get(ctx, s.name(token)).Val() != ""
}

func (s *JwtTokenCache) name(token string) string {
	return fmt.Sprintf("jwt:blacklist:%s", encrypt.Md5(token))
}
