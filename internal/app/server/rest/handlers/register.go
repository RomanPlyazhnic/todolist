package handlers

import (
	"net/http"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

func Register(a server.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "rest.Register"
	}
}
