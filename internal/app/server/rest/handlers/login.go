// Login handler

package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/RomanPlyazhnic/todolist/internal/app/auth/jwt"
	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

// LoginRequest represents request JSON body format
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var ErrInvalidCredentials = errors.New("invalid credentials")

// Login returns http handler
// Accepts JSON LoginRequest body
// Checks if login and password are correct and returns JWT token in header
func Login(a server.Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "rest.Login"

		w.Header().Set("Accept", "application/json")

		loginRequest := LoginRequest{}

		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&loginRequest)
		if err != nil {
			err = fmt.Errorf("%s: %w", op, err)
			a.Logger().Info("failed to decode loginRequest", op, err)
			http.Error(w, "invalid request", http.StatusBadRequest)

			return
		}

		if loginRequest.Username == "admin" || loginRequest.Password == "admin" {
			jwtToken, err := jwt.CreateToken(a, loginRequest.Username)
			if err != nil {
				err = fmt.Errorf("%s: %w", op, err)
				a.Logger().Info("failed to generate token", op, err)
				http.Error(w, "failed to login", http.StatusBadRequest)

				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "jwt",
				Value:    jwtToken,
				HttpOnly: true,
				Domain:   a.Config().Domain,
				SameSite: http.SameSiteDefaultMode,
			})

			_, err = w.Write([]byte("login successful"))
			if err != nil {
				a.Logger().Info("failed to write response", op, fmt.Errorf("%s: %w", op, err))
			}

			a.Logger().Info("login successful", op, true)

			return
		}

		a.Logger().Info("invalid credentials", op, fmt.Errorf("%s: %w", op, ErrInvalidCredentials))
		http.Error(w, "failed to login", http.StatusUnauthorized)
	}
}
