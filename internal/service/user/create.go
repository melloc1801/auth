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

	var userId int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		id, errTx := s.userRepository.Create(ctx, converter.ToCreateUserInfoFromService(createUserInfo))
		if errTx != nil {
			return errTx
		}
		userId = id

		return errors.New("asd")

		return s.userLogRepository.Create(ctx, "user_created")
	})
	if err != nil {
		return 0, err
	}

	return userId, nil
}
