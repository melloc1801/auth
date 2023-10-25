package model

import (
	"database/sql"
	"time"
)

type UserRole string

const (
	USER_ROLE  UserRole = "USER"
	ADMIN_ROLE UserRole = "ADMIN"
)

type User struct {
	Id        int64
	Name      string
	Email     string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type CreateUserInfo struct {
	Name     string
	Email    string
	Role     UserRole
	Password string
}

type UpdateUserInfo struct {
	Id    int64
	Name  string
	Email string
	Role  UserRole
}
