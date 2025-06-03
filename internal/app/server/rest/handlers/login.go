// Login handler

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RomanPlyazhnic/todolist/internal/app/contracts"
	"github.com/RomanPlyazhnic/todolist/internal/app/server"
	"github.com/RomanPlyazhnic/todolist/internal/core/auth"
)

// Login returns http handler
// Accepts JSON LoginRequest body
// Checks if login and password are correct and returns JWT token in the header
func Login(a *server.App) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "rest.Login"

		w.Header().Set("Accept", "application/json")

		loginRequest := contracts.LoginRequest{}

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&loginRequest)
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
			a.Logger.Info("failed to decode loginRequest", op, err)
			http.Error(w, "invalid request", http.StatusBadRequest)

			return
		}

		validateResult, err := loginRequest.Validate()
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
			a.Logger.Info("failed to validate loginRequest", op, err)
			http.Error(w, "invalid request", http.StatusBadRequest)

			enc := json.NewEncoder(w)
			err = enc.Encode(validateResult)
			if err != nil {
				a.Logger.Info("failed to write response", r.Method, fmt.Errorf("%s: %w", op, err))
			}

			return
		}

		jwtToken, err := auth.Login(r.Context(), a, loginRequest.Username, loginRequest.Password)
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
			a.Logger.Info("failed to login", op, err)
			http.Error(w, "failed to login", http.StatusBadRequest)

			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "jwt",
			Value:    jwtToken,
			HttpOnly: true,
			Domain:   a.Config.Domain,
			SameSite: http.SameSiteDefaultMode,
		})

		a.Logger.Info("login successful", op, true)
		_, err = w.Write([]byte("Login successfully"))
		if err != nil {
			a.Logger.Info("failed to write response", r.Method, fmt.Errorf("%s: %w", op, err))
		}
	}
}
