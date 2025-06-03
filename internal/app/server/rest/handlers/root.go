package handlers

import (
	"fmt"
	"net/http"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

func Root(a *server.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "rest.Root"

		a.Logger.Info("current user_id is "+fmt.Sprintf("%d", r.Context().Value("user_id").(int)), op, true)

		_, err := w.Write([]byte("Hello!"))
		if err != nil {
			a.Logger.Info("failed to write response", r.Method, fmt.Errorf("%s: %w", op, err))
		}
	}
}
