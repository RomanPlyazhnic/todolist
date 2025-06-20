// Create a todolist handler

package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/RomanPlyazhnic/todolist/internal/core/todolist"
	"net/http"

	"github.com/RomanPlyazhnic/todolist/internal/app/contracts"
	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

// CreateTodoList handles the creation of a new TodoList using data from the HTTP request body.
// Accepts JSON TodoList body
func CreateTodoList(a *server.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "rest.CreateTodoList"

		w.Header().Set("Accept", "application/json")

		request := contracts.TodoList{}

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&request)
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
			a.Logger.Info("failed to decode request", op, err)
			http.Error(w, "invalid request", http.StatusBadRequest)

			return
		}

		validateResult, err := request.Validate()
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
			a.Logger.Info("failed to validate request", op, err)
			http.Error(w, "invalid request", http.StatusBadRequest)

			enc := json.NewEncoder(w)
			err = enc.Encode(validateResult)
			if err != nil {
				a.Logger.Info("failed to write response", r.Method, fmt.Errorf("%s: %w", op, err))
			}
		}

		err = todolist.Create(a, request)
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
			a.Logger.Info("failed to create todolist", op, err)
			http.Error(w, "failed to create todolist", http.StatusBadRequest)

			return
		}

		a.Logger.Info("todolist is created successfully", op, true)
		_, err = w.Write([]byte("Todolist is created"))
		if err != nil {
			a.Logger.Info("failed to write response", r.Method, fmt.Errorf("%s: %w", op, err))
		}
	}
}
