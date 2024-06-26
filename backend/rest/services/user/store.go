package user

import (
	"database/sql"
	"fmt"

	"github.com/edgarbagagem/gochat/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetUserByUsername(username string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) UpdateUser(user *types.User) error {
	_, err := s.db.Exec("UPDATE users SET  photo = ? WHERE id = ?", user.Photo.String, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Photo,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
