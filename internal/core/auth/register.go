// Register implementation

package auth

import (
	"context"
	"fmt"
	"github.com/RomanPlyazhnic/todolist/internal/core/users"
	"golang.org/x/crypto/bcrypt"

	"github.com/RomanPlyazhnic/todolist/internal/app/contracts"
	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

// Register checks if user credentials for a new user are valid and user is unique
// If credentials are valid - returns nil
// If credentials are not valid - returns error
func Register(_ context.Context, a *server.App, reg contracts.RegisterRequest) error {
	const op = "auth.Register"

	passHash, err := hashPassword(reg.Password)
	if err != nil {
		a.Logger.Info("failed to generate password hash", op, err)

		return fmt.Errorf("%s: %w", op, err)
	}

	err = users.Create(a, reg.Username, passHash)
	if err != nil {
		a.Logger.Info("failed to register user", op, err)

		return fmt.Errorf("failed to register user", op, err)
	}

	a.Logger.Info("register successful", op, true)

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}
