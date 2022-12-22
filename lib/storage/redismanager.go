package storage

import (
	"time"

	"github.com/swethabhageerath/storage/internal/storage/redis"
)

type RedisManager struct{}

func (r RedisManager) Set(key string, value interface{}, expiration time.Duration) error {
	return redis.Redis{}.Set(key, value, expiration)
}

func (r RedisManager) Get(key string) (interface{}, error) {
	return redis.Redis{}.Get(key)
}

func (r RedisManager) Remove(key string, expiration time.Duration) error {
	return redis.Redis{}.Remove(key, expiration)
}
