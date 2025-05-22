package rest

import (
	"errors"
	"net/http"

	"github.com/RomanPlyazhnic/todolist/internal/app/auth/jwt"
	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

func JWTAuth(srv server.Server) func(http.Handler) http.Handler {
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
					srv.Logger().Info("cookie is empty", op, true)
					http.Error(w, "unauthorized", http.StatusUnauthorized)

					return
				}

				srv.Logger().Info("failed to retrieve cookie", op, err)
				http.Error(w, "internal error", http.StatusInternalServerError)

				return
			}

			_, err = jwt.ValidateToken(srv, c.Value)
			if err != nil {
				if errors.Is(err, jwt.InvalidToken) {
					srv.Logger().Info("jwt token is invalid", op, err)
					http.Error(w, "unauthorized", http.StatusUnauthorized)

					return
				}

				srv.Logger().Warn("failed to validate token", op, err)
				http.Error(w, "internal error", http.StatusInternalServerError)

				return
			}

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
