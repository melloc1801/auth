package user

import (
	"auth/internal/client/db"
	"auth/internal/repository"
	"auth/internal/repository/user/model"
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"time"
)

const (
	tableName = "\"user\""

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, userInfo *model.CreateUserInfo) (int64, error) {
	insertBuilder := squirrel.Insert(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Columns(nameColumn, emailColumn, roleColumn, passwordColumn).
		Values(userInfo.Name, userInfo.Email, userInfo.Role, userInfo.Password).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, errors.New("failed to build query")
	}

	var id int64

	ro, err := r.db.DB().Exec(ctx, db.Query{Name: "Insert user", QueryString: query}, args...)
	fmt.Println(ro)
	if err != nil {
		return 0, errors.New("failed to make query")
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	selectBuilder := squirrel.Select(
		idColumn,
		nameColumn,
		emailColumn,
		roleColumn,
		createdAtColumn,
		updatedAtColumn,
	).
		PlaceholderFormat(squirrel.Dollar).
		From(tableName).
		Where(squirrel.Eq{idColumn: id})

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, errors.New("failed to build query")
	}

	var user = &model.User{}
	err = r.db.DB().ScanOneRow(ctx, user, db.Query{Name: "user_repository.Get", QueryString: query}, args...)
	if err != nil {
		return nil, errors.New("failed to make query")
	}

	return user, nil
}

func (r *repo) Update(ctx context.Context, updateUserInfo *model.UpdateUserInfo) error {
	builderUpdate := squirrel.Update(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Set(emailColumn, updateUserInfo.Email).
		Set(nameColumn, updateUserInfo.Name).
		Set(roleColumn, updateUserInfo.Role).
		Set(updatedAtColumn, time.Now()).
		Where(squirrel.Eq{idColumn: updateUserInfo.Id})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return errors.New("failed to build query")
	}

	_, err = r.db.DB().Exec(ctx, db.Query{Name: "user_repository.Update", QueryString: query}, args...)
	if err != nil {
		return errors.New("failed to executed query")
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	deleteBuilder := squirrel.Delete(tableName).
		PlaceholderFormat(squirrel.Dollar).
		Where(squirrel.Eq{idColumn: id})

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return errors.New("failed to build query")
	}
	_, err = r.db.DB().Exec(ctx, db.Query{Name: "user_repository.Delete", QueryString: query}, args...)
	if err != nil {
		return errors.New("failed to execute")
	}

	return nil
}
