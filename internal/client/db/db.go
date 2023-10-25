package db

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Query struct {
	Name        string
	QueryString string
}

type Client interface {
	DB() DB
	Close() error
}

type DB interface {
	Execers
	Pinger
	Close()
}

type Execers interface {
	Execer
	ExecerWithScan
}

type Execer interface {
	Exec(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryAllRows(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryOneRow(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type ExecerWithScan interface {
	ScanOneRow(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllRows(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

type Pinger interface {
	Ping(ctx context.Context) error
}
