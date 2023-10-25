package api

import (
	userServiceConverter "auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateUserRequest) (*desc.CreateUserResponse, error) {
	id, err := i.userService.Create(ctx, userServiceConverter.ToCreateUserInfoFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateUserResponse{
		Id: id,
	}, nil
}
