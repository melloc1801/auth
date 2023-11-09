package api

import (
	desc "auth/pkg/user_v1"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteUserRequest) (*empty.Empty, error) {
	err := i.userService.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
