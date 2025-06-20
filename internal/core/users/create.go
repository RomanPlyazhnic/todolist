package users

import (
	"fmt"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

func Create(a *server.App, username, password_hash string) (err error) {
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
