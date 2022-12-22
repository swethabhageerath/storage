package secrets

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	cfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/pkg/errors"
)

const (
	KEY_AWS_SECRETSMANAGER_REGION        = "AWS_SECRETSMANAGER_REGION"
	KEY_AWS_SECRETSMANAGER_VERSION_STAGE = "AWSCURRENT"
)

type Secrets struct{}

func (s Secrets) GetValueString(ctx context.Context, secretName string) (string, error) {
	client, err := s.getClient(ctx)

	if err != nil {
		return "", errors.Wrap(err, NewErrRetrievingAwsSecretsManagerClient().Message)
	}

	input := s.getRequestInput(secretName)

	output, err := client.GetSecretValue(ctx, input)

	if err != nil {
		return "", errors.Wrap(err, NewErrRetrievingSecretFromAwsSecretsManager().Message)
	}

	return *output.SecretString, nil
}

func (s Secrets) getConfig(ctx context.Context) (aws.Config, error) {
	region := s.getRegion(KEY_AWS_SECRETSMANAGER_REGION)

	if region == "" {
		return aws.Config{}, errors.New(NewErrRegionNotSpecifiedForSecretsManager().Message)
	}

	c, err := cfg.LoadDefaultConfig(ctx, cfg.WithRegion(region))
	if err != nil {
		er := errors.Wrap(err, NewErrLoadingConfiguringForAwsSecretsManager().Message)
		return aws.Config{}, er
	}

	return c, nil
}

func (s Secrets) getRequestInput(secretName string) *secretsmanager.GetSecretValueInput {
	return &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String(KEY_AWS_SECRETSMANAGER_VERSION_STAGE),
	}
}

func (s Secrets) getClient(ctx context.Context) (*secretsmanager.Client, error) {
	config, err := s.getConfig(ctx)
	if err != nil {
		return nil, err
	}
	return secretsmanager.NewFromConfig(config), nil
}

func (s Secrets) getRegion(key string) string {
	return os.Getenv(key)
}
