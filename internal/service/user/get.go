package user

import (
	"auth/internal/model"
	"auth/internal/repository/user/converter"
	"context"
)

func (s serv) Get(ctx context.Context, id int64) (*model.User, error) {
	user, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(user), nil
}
