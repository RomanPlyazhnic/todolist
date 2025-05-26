// Register handler

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomanPlyazhnic/todolist/internal/app/contracts"
	"github.com/RomanPlyazhnic/todolist/internal/app/server"
	"github.com/RomanPlyazhnic/todolist/internal/core/auth"
)

// Register returns http handler
// Accepts JSON RegisterRequest body
// Checks if login and password are correct and responds with corresponding status
func Register(a *server.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "rest.Register"

		registerRequest := contracts.RegisterRequest{}

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&registerRequest)
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
			a.Logger.Info("failed to decode registerRequest", op, err)
			http.Error(w, "invalid request", http.StatusBadRequest)

			return
		}

		validateResult, err := registerRequest.Validate()
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
			a.Logger.Info("failed to validate registerRequest", op, err)
			http.Error(w, "invalid request", http.StatusBadRequest)

			enc := json.NewEncoder(w)
			err = enc.Encode(validateResult)
			if err != nil {
				a.Logger.Info("failed to write response", r.Method, fmt.Errorf("%s: %w", op, err))
			}

			return
		}

		err = auth.Register(a, registerRequest.Username, registerRequest.Password)
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
			a.Logger.Info("failed to register", op, err)
			http.Error(w, "failed to register", http.StatusUnauthorized)

			return
		}

		a.Logger.Info("registered successfully", op, true)

		_, err = w.Write([]byte("Registered successfully"))
		if err != nil {
			a.Logger.Info("failed to write response", r.Method, fmt.Errorf("%s: %w", op, err))
		}
	}
}
