package types

import "database/sql"

type UserStore interface {
	GetUserByUsername(username string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID       int            `json:"id"`
	Username string         `json:"username"`
	Password string         `json:"password"`
	Photo    sql.NullString `json:"photo"`
}

type RegisterUserPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
