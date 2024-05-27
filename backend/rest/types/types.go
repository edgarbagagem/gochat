package types

type UserStore interface {
	GetUserByUsername(username string) (*User, error)
	GetUserById(id int) (*User, error)
	CreateUser(User) error
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Photo    string `json:"photo"`
}

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
