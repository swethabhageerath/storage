package storage

import (
	"context"
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

func (a AwsSecretsManager) GetValueString(ctx context.Context, secretName string, out chan AwsSecretsManagerResponse) {
	var d time.Duration = 100000000000
	i, err := a.redisManager.Get(secretName)

	if err != nil || i == nil {
		v, err := secrets.Secrets{}.GetValueString(ctx, secretName)

		a.redisManager.Set(secretName, v, time.Duration(d.Hours()))

		if err != nil {
			out <- AwsSecretsManagerResponse{
				Data:  "",
				Error: err,
			}
		} else {
			out <- AwsSecretsManagerResponse{
				Data:  v,
				Error: nil,
			}
		}
	} else {
		out <- AwsSecretsManagerResponse{
			Data:  i.(string),
			Error: nil,
		}
	}
}
