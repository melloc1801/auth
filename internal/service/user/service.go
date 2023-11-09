package user

import (
	"auth/internal/client/db"
	"auth/internal/repository"
	"auth/internal/service"
)

type serv struct {
	txManager db.TxManager

	userRepository    repository.UserRepository
	userLogRepository repository.UserLogRepository
}

var _ service.UserService = (*serv)(nil)

func NewService(
	txManager db.TxManager,
	userRepository repository.UserRepository,
	userLogRepository repository.UserLogRepository,
) service.UserService {
	return &serv{
		txManager,
		userRepository,
		userLogRepository,
	}
}
