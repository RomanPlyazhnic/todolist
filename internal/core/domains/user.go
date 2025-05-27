package domains

import (
	"fmt"
	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

func CreateUser(a *server.App, username, password string) (err error) {
	const op = "domains.CreateUser"

	sqlStatement := `
INSERT INTO users (username, password)
VALUES ($1, $2)`

	res, err := a.DB.Exec(sqlStatement, username, password)
	if err != nil {
		a.Logger.Info(fmt.Sprintf("%s: %w", op, err))
		return fmt.Errorf("failed to create user", op, err)
	}

	id, _ := res.LastInsertId()
	a.Logger.Info(fmt.Sprintf("successfully created user, id: %d", id), op, true)

	return nil
}
