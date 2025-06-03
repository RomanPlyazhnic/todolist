// Register implementation

package auth

import (
	"context"
	"fmt"
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

	err = createUser(a, reg.Username, passHash)
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

func createUser(a *server.App, username, password_hash string) (err error) {
	const op = "auth.createUser"

	const sqlStatement = `
INSERT INTO users (username, password_hash)
VALUES ($1, $2)`

	res, err := a.DB.Exec(sqlStatement, username, password_hash)
	if err != nil {
		a.Logger.Info(fmt.Sprintf("%s: %w", op, err))
		return fmt.Errorf("failed to create user", op, err)
	}

	id, _ := res.LastInsertId()
	a.Logger.Info(fmt.Sprintf("successfully created user, id: %d", id), op, true)

	return nil
}
