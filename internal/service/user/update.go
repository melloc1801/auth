package user

import (
	"auth/internal/model"
	"auth/internal/repository/user/converter"
	"context"
)

func (s serv) Update(ctx context.Context, updateUserInfo *model.UpdateUserInfo) error {
	err := s.userRepository.Update(ctx, converter.ToUpdateUserInfoFromService(updateUserInfo))

	return err
}
