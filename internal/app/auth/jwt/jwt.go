package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

var InvalidToken = errors.New("invalid token")

type MyCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(srv server.Server, username string) (tokenString string, err error) {
	const op = "jwt.CreateToken"

	claims := MyCustomClaims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(srv.Config().JWT.TokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    srv.Config().Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(srv.Config().JWT.Secret))
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		srv.Logger().Info("failed to generate token", op, err)
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(srv server.Server, tokenString string) (token *jwt.Token, err error) {
	const op = "jwt.ValidateToken"

	token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(srv.Config().JWT.Secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		srv.Logger().Info("failed to validate token", op, err)
	}

	if !token.Valid {
		return token, fmt.Errorf("%s: %w", op, InvalidToken)
	}

	return token, err
}
