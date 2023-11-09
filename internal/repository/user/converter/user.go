package converter

import (
	userServiceModel "auth/internal/model"
	userRepoModel "auth/internal/repository/user/model"
	"time"
)

func ToCreateUserInfoFromService(userInfo *userServiceModel.CreateUserInfo) *userRepoModel.CreateUserInfo {
	return &userRepoModel.CreateUserInfo{
		Name:     userInfo.Name,
		Email:    userInfo.Email,
		Role:     userRepoModel.UserRole(userInfo.Role),
		Password: userInfo.Password,
	}
}

func ToUpdateUserInfoFromService(userInfo *userServiceModel.UpdateUserInfo) *userRepoModel.UpdateUserInfo {
	return &userRepoModel.UpdateUserInfo{
		Id:    userInfo.Id,
		Name:  userInfo.Name,
		Email: userInfo.Email,
		Role:  userRepoModel.UserRole(userInfo.Role),
	}
}

func ToUserFromRepo(user *userRepoModel.User) *userServiceModel.User {
	var updatedAt *time.Time
	if user.UpdatedAt.Valid {
		updatedAt = &user.UpdatedAt.Time
	}

	return &userServiceModel.User{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email,
		Role:      userServiceModel.UserRole(user.Role),
		CreatedAt: user.CreatedAt,
		UpdateAt:  updatedAt,
	}
}
