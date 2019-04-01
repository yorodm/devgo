package engine

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrAuth signals an authentication error
	ErrAuth = errors.New("Wrong user or password")
)

//Login makes a login attempt with the given user name and password
func (e *Engine) Login(ctx context.Context, username, password string) error {
	var hashedPassword string
	query := "select password from users where username = $1"

	err := e.db.QueryRowContext(ctx, query, username).Scan(&hashedPassword)
	if err != nil {
		// Check for ErrNoRow
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return ErrAuth
	}
	return nil
}
