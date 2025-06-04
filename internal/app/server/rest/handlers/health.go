package handlers

import (
	"fmt"
	"net/http"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

func Health(a *server.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "rest.Health"

		a.Logger.Info("Health check started", op, true)

		_, err := w.Write([]byte("Health check"))
		if err != nil {
			a.Logger.Info("failed to write response", r.Method, fmt.Errorf("%s: %w", op, err))
		}

		a.Logger.Info("Health check completed", op, true)
	}
}
