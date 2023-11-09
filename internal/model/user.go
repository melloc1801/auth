package model

import (
	"time"
)

type UserRole string

const (
	USER_ROLE  UserRole = "USER"
	ADMIN_ROLE UserRole = "ADMIN"
)

type CreateUserInfo struct {
	Name            string
	Email           string
	Role            UserRole
	Password        string
	PasswordConfirm string
}

type User struct {
	Id        int64
	Name      string
	Email     string
	Role      UserRole
	CreatedAt time.Time
	UpdateAt  *time.Time
}

type UpdateUserInfo struct {
	Id    int64
	Name  string
	Email string
	Role  UserRole
}
