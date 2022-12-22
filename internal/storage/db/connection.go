package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

type Connection struct{}

func (c Connection) Connect(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Error connecting to PostGres database")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "Error pining PostGres Db")
	}

	return db, nil
}
