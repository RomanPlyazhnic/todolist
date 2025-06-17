// Login implementation

package auth

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

// Login generates JWT token based on user credentials
// If credentials are not valid - returns error
func Login(_ context.Context, a *server.App, username, password string) (jwtToken string, err error) {
	const op = "auth.Login"

	userId, err := checkCredentials(a, username, password)
	if err != nil {
		a.Logger.Info("failed to check credentials", op, err)

		return "", fmt.Errorf("%s: %w", op, err)
	}

	jwtToken, err = CreateToken(a, userId)
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		a.Logger.Info("failed to generate token", op, err)

		return "", err
	}

	return jwtToken, nil
}

// checkCredentials finds user by username, checks password and returns user_id if the password is correct.
// Otherwise - return error
func checkCredentials(a *server.App, username, password string) (userId int, err error) {
	const op = "auth.checkCredentials"

	var storedHash string
	err = a.DB.QueryRow("SELECT id, password_hash FROM users WHERE username = $1", username).Scan(&userId, &storedHash)
	if err != nil {
		a.Logger.Info("failed to retrieve password hash", op, err)

		return userId, fmt.Errorf("%s: %w", op, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		a.Logger.Info("failed to compare password and password hash", op, err)

		return userId, fmt.Errorf("%s: %w", op, err)
	}

	return userId, nil
}
