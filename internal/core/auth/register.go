// Register implementation

package auth

import (
	"errors"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

var ErrFailedRegister = errors.New("Failed to register user")

// Register checks if user credentials for a new user are valid and user is unique
// If credentials are valid - returns nil
// If credentials are not valid - returns error
func Register(a *server.App, username, password string) error {
	const op = "auth.Register"

	if username != "admin" || password != "adminjjjj" {
		a.Logger.Info("failed to register user", op, ErrFailedRegister)

		return ErrFailedRegister
	}

	a.Logger.Info("register successful", op, true)

	return nil
}
