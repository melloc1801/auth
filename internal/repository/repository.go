package repository

import (
	"auth/internal/repository/user/model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, userInfo *model.CreateUserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, updateUserInfo *model.UpdateUserInfo) error
	Delete(ctx context.Context, id int64) error
}

type UserLogRepository interface {
	Create(ctx context.Context, message string) error
}
