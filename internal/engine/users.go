package engine

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

// User data in the blog
type User struct {
	ID       int64  `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username"`
}

// UserProfile (metadata)
type UserProfile struct {
	User
	name string `json:"name"`
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CreateUser creates and stores a new User in the database
func (e *Engine) CreateUser(ctx context.Context, name, email, username, password string) error {
	var query = "insert into users (name, email, username, password) values ($1, $2, $3, $4)"
	hashed, err := hashPassword(password)
	if err != nil {
		return err
	}
	_, err = e.db.ExecContext(ctx, query, name, email, username, hashed)
	if err != nil {
		return err
	}
	return nil
}

// ListUsers returns all users in the database
func (e *Engine) ListUsers(ctx context.Context) (users []User, err error) {
	var query = "select id, email, username from users"
	rows, err := e.db.QueryContext(ctx, query)
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Email, &user.Username)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	return
}
