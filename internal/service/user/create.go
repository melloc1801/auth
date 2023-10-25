package user

import (
	"auth/internal/model"
	"auth/internal/repository/user/converter"
	"context"
	"errors"
)

func (s serv) Create(ctx context.Context, createUserInfo *model.CreateUserInfo) (int64, error) {
	if createUserInfo.Password != createUserInfo.PasswordConfirm {
		return 0, errors.New("passwords should be equal")
	}

	userId, err := s.userRepository.Create(ctx, converter.ToCreateUserInfoFromService(createUserInfo))
	if err != nil {
		return 0, err
	}

	return userId, nil
}
