package handlers

import (
	"fmt"
	"net/http"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

func Get(a server.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "rest.Get"

		_, err := w.Write([]byte("Hello!"))
		if err != nil {
			a.Logger().Info("failed to write response", r.Method, fmt.Errorf("%s: %w", op, err))
		}
	}
}
