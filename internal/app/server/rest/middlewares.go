// JWT auth middleware

package rest

import (
	"context"
	"errors"
	"net/http"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
	"github.com/RomanPlyazhnic/todolist/internal/core/auth"
)

// JWTAuth middleware checks if JWT token is valid
// If it is - it passes the request to the next handler
// If it is not - it returns 401 error
// If cookie is not present - it returns 401 error
// If cookie is present but token is invalid - it returns 401 error
// If cookie is present but token is valid, but it is not registered - it returns 401 error
// If cookie is present but token is valid, and it is registered, but it is not authorized - it returns 401 error
func JWTAuth(a *server.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			const op = "rest.JWTAuth"

			if r.URL.Path == "/Login" || r.URL.Path == "/Register" {
				next.ServeHTTP(w, r)

				return
			}

			c, err := r.Cookie("jwt")

			if err != nil {
				if errors.Is(err, http.ErrNoCookie) {
					a.Logger.Info("cookie is empty", op, true)
					http.Error(w, "unauthorized", http.StatusUnauthorized)

					return
				}

				a.Logger.Info("failed to retrieve cookie", op, err)
				http.Error(w, "internal error", http.StatusInternalServerError)

				return
			}

			claim, err := auth.ValidateToken(a, c.Value)
			if err != nil {
				if errors.Is(err, auth.InvalidToken) {
					a.Logger.Info("jwt token is invalid", op, err)
					http.Error(w, "unauthorized", http.StatusUnauthorized)

					return
				}

				a.Logger.Warn("failed to validate token", op, err)
				http.Error(w, "internal error", http.StatusInternalServerError)

				return
			}

			ctx := context.WithValue(r.Context(), "user_id", claim.UserId)

			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
