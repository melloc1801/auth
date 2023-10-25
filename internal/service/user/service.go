package user

import (
	"auth/internal/repository"
	"auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

var _ service.UserService = (*serv)(nil)

func NewService(
	userRepository repository.UserRepository,
) service.UserService {
	return &serv{
		userRepository,
	}
}
