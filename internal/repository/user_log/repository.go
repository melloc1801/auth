package user_log

import (
	"auth/internal/client/db"
	"auth/internal/repository"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

const (
	tableName     = "user_log"
	messageColumn = "message"
)

type repo struct {
	db db.Client
}

func NewUserLogRepository(db db.Client) repository.UserLogRepository {
	return &repo{
		db: db,
	}
}

func (r repo) Create(ctx context.Context, message string) error {
	insertBuilder := squirrel.Insert(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(messageColumn).
		Values(message).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return errors.Wrapf(err, "failed to build query: %s", err)
	}

	_, err = r.db.DB().Exec(ctx, db.Query{Name: "user_log_repository.Create", QueryString: query}, args...)
	if err != nil {
		return errors.Wrapf(err, "failed to make query: %s", err)
	}

	return nil
}
