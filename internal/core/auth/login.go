package auth

import (
	"errors"
	"fmt"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

func Login(a server.Server, username, password string) (jwtToken string, err error) {
	const op = "auth.Login"

	if username != "admin" || password != "admin" {
		return "", ErrInvalidCredentials
	}

	jwtToken, err = CreateToken(a, username)
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		a.Logger().Info("failed to generate token", op, err)

		return "", err
	}

	return jwtToken, nil
}
