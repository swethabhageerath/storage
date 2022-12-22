package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type postGres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) postGres {
	return postGres{
		db: db,
	}
}

func (p postGres) Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.db.ExecContext(ctx, query, args...)
}

func (p postGres) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return p.db.QueryRowContext(ctx, query, args...)
}

func (p postGres) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.QueryContext(ctx, query, args...)
}
