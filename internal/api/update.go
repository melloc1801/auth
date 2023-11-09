package api

import (
	userServiceConverter "auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) Update(ctx context.Context, req *desc.UpdateUserRequest) (*empty.Empty, error) {
	err := i.userService.Update(ctx, userServiceConverter.ToUpdateUserInfoFromDesc(req))
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
