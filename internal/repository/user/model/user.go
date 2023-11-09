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
	Id        int64        `db:"id"`
	Name      string       `db:"name"`
	Email     string       `db:"email"`
	Role      UserRole     `db:"role"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
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
