package storage

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/swethabhageerath/storage/internal/storage/aws/secrets"
)

type AwsSecretsManager struct {
	redisManager RedisManager
}

func NewAwsSecretsManager(redisManager RedisManager) AwsSecretsManager {
	return AwsSecretsManager{
		redisManager: redisManager,
	}
}

type AwsSecretsManagerResponse struct {
	Data  string
	Error error
}

func (a AwsSecretsManager) getConnectionStringRedisExpirationTime() (int64, error) {
	expirationAsString := os.Getenv("REDIS_CONNSTRING_EXPIRATION")
	s, err := strconv.ParseInt(expirationAsString, 10, 64)
	if err != nil {
		return 0, err
	}
	return s, nil
}

func (a AwsSecretsManager) GetValueString(ctx context.Context, secretName string, out chan AwsSecretsManagerResponse) {
	t, err := a.getConnectionStringRedisExpirationTime()

	if err != nil {
		out <- AwsSecretsManagerResponse{
			Data:  "",
			Error: err,
		}
		return
	}

	i, err := a.redisManager.Get(secretName)

	response := AwsSecretsManagerResponse{}

	if err != nil || i == nil {
		v, err := secrets.Secrets{}.GetValueString(ctx, secretName)

		a.redisManager.Set(secretName, v, time.Duration(t))

		if err != nil {
			response.Data = ""
			response.Error = err
		} else {
			response.Data = v
			response.Error = nil
		}
	} else {
		response.Data = i.(string)
		response.Error = err
	}

	out <- response
}
