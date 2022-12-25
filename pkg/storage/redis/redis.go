package redis

import (
	"fmt"
	"time"

	"github.com/pkg/errors"

	r "github.com/go-redis/redis"
)

type Redis struct {
	client *r.Client
}

func New(client *r.Client) Redis {
	return Redis{
		client: client,
	}
}

// func (f Redis) getClient() (*r.Client, error) {
// 	s, err := f.getAddress()

// 	if err != nil {
// 		return nil, err
// 	}

// 	return r.NewClient(&r.Options{
// 		Addr:     s,
// 		Password: "",
// 		DB:       0,
// 	}), nil
// }

// func (f Redis) getAddress() (string, error) {
// 	redisAddr := os.Getenv("REDIS_ADDR")

// 	if redisAddr == "" {
// 		return "", errors.New("REDIS_ADDR is not present in Environment Variable")
// 	}

// 	return redisAddr, nil
// }

func (r Redis) Set(key string, value interface{}, expiration time.Duration) error {
	s := r.client.Set(key, value, expiration)
	err := s.Err()
	if err != nil {
		return errors.Wrap(err, NewErrSetRedisCache().Message)
	}
	return nil
}

func (re Redis) Get(key string) (interface{}, error) {
	s := re.client.Get(key)
	result, err := s.Result()
	if err != nil {
		if err == r.Nil {
			return nil, errors.Wrap(err, fmt.Sprintf(NewErrKeyNotExists().Message, key))
		} else {
			return nil, errors.Wrap(err, fmt.Sprintf(NewErrRetrievingKey().Message, key))
		}
	}
	return result, nil
}

func (re Redis) Remove(key string, expiration time.Duration) error {
	e := re.client.Expire(key, expiration)
	err := e.Err()
	if err != nil {
		if err == r.Nil {
			return errors.Wrap(err, fmt.Sprintf(NewErrKeyNotExists().Message, key))
		} else {
			return errors.Wrap(err, fmt.Sprintf(NewErrRetrievingKey().Message, key))
		}
	}

	return err
}
