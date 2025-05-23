package auth

import (
	"errors"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

var ErrFailedRegister = errors.New("Failed to register user")

func Register(a server.Server, username, password string) error {
	const op = "auth.Register"

	if username != "admin" || password != "adminjjjj" {
		a.Logger().Info("failed to register user", op, ErrFailedRegister)

		return ErrFailedRegister
	}

	a.Logger().Info("register successful", op, true)

	return nil
}
