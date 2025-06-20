package todolist

import (
	"fmt"
	"github.com/RomanPlyazhnic/todolist/internal/app/contracts"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

// TODO: add check whether current user is able to create todolist for provided user_id
// TODO: add tests
// TODO: fill checkboxes with todolist id
func Create(a *server.App, todolist contracts.TodoList) (err error) {
	const op = "auth.createUser"

	const sqlStatementTodolist = `
INSERT INTO todolists (user_id, text)
VALUES ($1, $2)`
	const sqlStatementCheckboxes = `
INSERT INTO checkboxes (checked, text)
VALUES ($1, $2)`

	// TODO: do inside a transaction
	_, err = a.DB.Exec(sqlStatementTodolist, todolist.UserId, todolist.Text)
	if err != nil {
		a.Logger.Info(fmt.Sprintf("%s: %w", op, err))
		return fmt.Errorf("failed to create todolist", op, err)
	}

	for _, checkbox := range todolist.Checkboxes {
		_, err = a.DB.Exec(sqlStatementCheckboxes, checkbox.Checked, checkbox.Text)
		if err != nil {
			a.Logger.Info(fmt.Sprintf("%s: %w", op, err))
			return fmt.Errorf("failed to create todolist", op, err)
		}
	}

	return nil
}
