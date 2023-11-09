package converter

import (
	"auth/internal/model"
	desc "auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreateUserInfoFromDesc(request *desc.CreateUserRequest) *model.CreateUserInfo {
	user := request.User

	return &model.CreateUserInfo{
		Name:            user.Name,
		Email:           user.Email,
		Role:            model.UserRole(user.Role.String()),
		Password:        user.Password,
		PasswordConfirm: user.PasswordConfirm,
	}
}

func ToUpdateUserInfoFromDesc(request *desc.UpdateUserRequest) *model.UpdateUserInfo {
	return &model.UpdateUserInfo{
		Id:    request.Id,
		Name:  request.Name.Value,
		Email: request.Email.Value,
		Role:  model.UserRole(request.Role.String()),
	}
}

func ToDescFromService(user *model.User) *desc.GetUserResponse {
	var updatedAt *timestamppb.Timestamp
	if user.UpdateAt != nil {
		updatedAt = timestamppb.New(*user.UpdateAt)
	}

	return &desc.GetUserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      desc.Role(desc.Role_value[string(user.Role)]),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
