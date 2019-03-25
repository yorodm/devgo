package engine

import (
	"bytes"
	"context"
	"crypto/sha256"
)

type User struct {
	ID       int64  `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username"`
}

type UserProfile struct {
	User
	name string `json:"name"`
}

// CreateUser creates and stores a new User in the database
func (e *Engine) CreateUser(ctx context.Context, email, username, password string) error {
	var query = "insert into users (email, username, password) values ($1, $2, $3)"
	hashed := bytes.NewBufferString(password).Bytes()
	_, err := e.db.ExecContext(ctx, query, email, username, sha256.Sum256(hashed))
	if err != nil {
		return err
	}
	return nil
}

// Users returns all users in the database
func (e *Engine) Users(ctx context.Context) ([]User, error) {

	return nil, nil
}
