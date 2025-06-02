package domains

import (
	"fmt"
	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

func CreateUser(a *server.App, username, password_hash string) (err error) {
	const op = "domains.CreateUser"

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

func UserHashPassword(a *server.App, username string) (string, error) {
	const op = "domains.hashPassword"

	var storedHash string
	err := a.DB.QueryRow("SELECT password_hash FROM users WHERE username = $1", username).Scan(&storedHash)
	if err != nil {
		a.Logger.Info(fmt.Sprintf("failed to retrieve password hash", op, err))

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return storedHash, nil
}
