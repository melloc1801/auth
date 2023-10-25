package pg

import (
	"auth/internal/client/db"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type pgClient struct {
	masterDBC db.DB
}

func NewPgClient(ctx context.Context, dsn string) (db.Client, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, nil
	}

	return &pgClient{
		masterDBC: &pg{
			pool: pool,
		},
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
