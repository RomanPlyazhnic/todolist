// Register implementation

package auth

import (
	"fmt"
	"github.com/RomanPlyazhnic/todolist/internal/core/domains"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

// Register checks if user credentials for a new user are valid and user is unique
// If credentials are valid - returns nil
// If credentials are not valid - returns error
func Register(a *server.App, username, password string) error {
	const op = "auth.Register"

	err := domains.CreateUser(a, username, password)
	if err != nil {
		a.Logger.Info("failed to register user", op, err)

		return fmt.Errorf("failed to register user", op, err)
	}

	a.Logger.Info("register successful", op, true)

	return nil
}
