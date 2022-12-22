package storage

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/swethabhageerath/storage/internal/storage/db"
)

type DbConnectionManagerResponse struct {
	Db    *sql.DB
	Error error
}

type DbConnectionManager struct {
	awsSecretsManager AwsSecretsManager
}

func NewDbConnectionManager(awsSecretsManager AwsSecretsManager) DbConnectionManager {
	return DbConnectionManager{
		awsSecretsManager: awsSecretsManager,
	}
}

type DbConnectionStringKey int

const (
	POSTGRES DbConnectionStringKey = iota
)

func (d DbConnectionStringKey) String() string {
	switch d {
	case POSTGRES:
		return "PgConnection"
	default:
		return "PgConnection"
	}
}

func (d DbConnectionManager) getConnection(connectionStringData *string) (*db.ConnectionString, error) {
	cs := new(db.ConnectionString)

	err := json.Unmarshal([]byte(*connectionStringData), cs)

	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling connectionstring from AWS Secrets Manager")
	}
	return cs, nil
}

func (d DbConnectionManager) connect(connectionString string) (*sql.DB, error) {
	conn := db.Connection{}

	dc, err := conn.Connect(connectionString)

	if err != nil {
		return nil, err
	}

	return dc, nil
}

func (c DbConnectionManager) Connect(ctx context.Context, key DbConnectionStringKey, out chan DbConnectionManagerResponse) {
	awsSecretManagerResponseChannel := make(chan AwsSecretsManagerResponse)
	go c.awsSecretsManager.GetValueString(ctx, key.String(), awsSecretManagerResponseChannel)

	secretsManagerResponse := <-awsSecretManagerResponseChannel

	if secretsManagerResponse.Error != nil {
		out <- DbConnectionManagerResponse{
			Db:    nil,
			Error: secretsManagerResponse.Error,
		}
	} else {
		cs, err := c.getConnection(&secretsManagerResponse.Data)

		if err != nil {
			out <- DbConnectionManagerResponse{
				Db:    nil,
				Error: err,
			}
		}

		d, err := c.connect(cs.ConnectionString)

		if err != nil {
			out <- DbConnectionManagerResponse{
				Db:    nil,
				Error: err,
			}
		}

		out <- DbConnectionManagerResponse{
			Db:    d,
			Error: nil,
		}
	}
}
