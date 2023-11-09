package service

import (
	"auth/internal/model"
	"context"
)

type UserService interface {
	Create(ctx context.Context, createUserInfo *model.CreateUserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, updateUserInfo *model.UpdateUserInfo) error
	Delete(ctx context.Context, id int64) error
}
