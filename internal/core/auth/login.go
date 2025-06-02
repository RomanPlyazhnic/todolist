// Login implementation

package auth

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
	"github.com/RomanPlyazhnic/todolist/internal/core/domains"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

// Login generates JWT token based on user credentials
// If credentials are not valid - returns error
func Login(a *server.App, username, password string) (jwtToken string, err error) {
	const op = "auth.Login"

	err = checkCredentials(a, username, password)
	if err != nil {
		a.Logger.Info(fmt.Sprintf("failed to check credentials", op, err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	jwtToken, err = CreateToken(a, username)
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		a.Logger.Info("failed to generate token", op, err)

		return "", err
	}

	return jwtToken, nil
}

func checkCredentials(a *server.App, username, password string) error {
	const op = "auth.checkCredentials"

	passHash, err := domains.UserHashPassword(a, username)
	if err != nil {
		a.Logger.Info(fmt.Sprintf("failed to retrieve password by username", op, err))

		return fmt.Errorf("%s: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password))
	if err != nil {
		a.Logger.Info(fmt.Sprintf("failed to compare password and password hash", op, err))

		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
