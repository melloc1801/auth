package pg

import (
	"auth/internal/client/db"
	"auth/internal/client/db/prettier"
	"context"
	"fmt"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type pg struct {
	pool *pgxpool.Pool
}

func NewPg(ctx context.Context, dbDsn string) (db.DB, error) {
	pool, err := pgxpool.Connect(ctx, dbDsn)
	if err != nil {
		return nil, err
	}
	return &pg{
		pool: pool,
	}, nil
}

func (pg *pg) Exec(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q, args...)

	res, err := pg.pool.Exec(ctx, q.QueryString, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (pg *pg) QueryOneRow(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, q, args...)

	return pg.pool.QueryRow(ctx, q.QueryString, args...)
}

func (pg *pg) QueryAllRows(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q, args...)

	return pg.pool.Query(ctx, q.QueryString, args)
}

func (pg *pg) ScanOneRow(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	logQuery(ctx, q, args...)

	rows, err := pg.pool.Query(ctx, q.QueryString, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanOne(dest, rows)
}

func (pg *pg) ScanAllRows(ctx context.Context, dest interface{}, q db.Query, args ...interface{}) error {
	logQuery(ctx, q, args...)

	rows, err := pg.pool.Query(ctx, q.QueryString, args...)
	if err != nil {
		return err
	}

	return pgxscan.ScanAll(dest, rows)
}

func (pg *pg) Ping(ctx context.Context) error {
	return pg.pool.Ping(ctx)
}

func (pg *pg) Close() {
	defer pg.Close()
}

func logQuery(ctx context.Context, q db.Query, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryString, prettier.PlaceholderDollar, args...)
	log.Println(
		ctx,
		fmt.Sprintf("sql: %s", q.Name),
		fmt.Sprintf("query: %s", prettyQuery),
	)
}
