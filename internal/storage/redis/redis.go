package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/pkg/errors"

	"github.com/go-redis/redis"
	r "github.com/go-redis/redis"
)

type Redis struct{}

func (f Redis) getClient() (*r.Client, error) {
	s, err := f.getAddress()

	if err != nil {
		return nil, err
	}

	return r.NewClient(&r.Options{
		Addr:     s,
		Password: "",
		DB:       0,
	}), nil
}

func (f Redis) getAddress() (string, error) {
	redisAddr := os.Getenv("REDIS_ADDR")

	if redisAddr == "" {
		return "", errors.New("REDIS_ADDR is not present in Environment Variable")
	}

	return redisAddr, nil
}

func (r Redis) Set(key string, value interface{}, expiration time.Duration) error {
	c, err := r.getClient()
	if err != nil {
		return err
	}

	s := c.Set(key, value, expiration)
	err = s.Err()
	if err != nil {
		return errors.Wrap(err, NewErrSetRedisCache().Message)
	}
	return nil
}

func (r Redis) Get(key string) (interface{}, error) {
	c, err := r.getClient()
	if err != nil {
		return nil, err
	}
	s := c.Get(key)
	result, err := s.Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.Wrap(err, fmt.Sprintf(NewErrKeyNotExists().Message, key))
		} else {
			return nil, errors.Wrap(err, fmt.Sprintf(NewErrRetrievingKey().Message, key))
		}
	}
	return result, nil
}

func (r Redis) Remove(key string, expiration time.Duration) error {
	c, err := r.getClient()
	if err != nil {
		return err
	}
	e := c.Expire(key, expiration)
	err = e.Err()
	if err != nil {
		if err == redis.Nil {
			return errors.Wrap(err, fmt.Sprintf(NewErrKeyNotExists().Message, key))
		} else {
			return errors.Wrap(err, fmt.Sprintf(NewErrRetrievingKey().Message, key))
		}
	}

	return err
}
