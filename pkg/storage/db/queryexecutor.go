package db

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type PostGres struct {
	db *sql.DB
}

func NewPostGres(db *sql.DB) PostGres {
	return PostGres{
		db: db,
	}
}

func (p PostGres) Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.db.ExecContext(ctx, query, args...)
}

func (p PostGres) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return p.db.QueryRowContext(ctx, query, args...)
}

func (p PostGres) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.db.QueryContext(ctx, query, args...)
}
