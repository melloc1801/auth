package api

import (
	userServiceConverter "auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	userFromDb, err := i.userService.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return userServiceConverter.ToDescFromService(userFromDb), nil
}
