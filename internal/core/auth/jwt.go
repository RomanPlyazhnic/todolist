// JWT auth implementation

package auth

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"

	"github.com/RomanPlyazhnic/todolist/internal/app/server"
)

var InvalidToken = errors.New("invalid token")

// MyCustomClaims represents JWT token format
type MyCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

// CreateToken generates JWT token based on user
func CreateToken(a *server.App, userId int) (tokenString string, err error) {
	const op = "jwt.CreateToken"

	claims := MyCustomClaims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(a.Config.JWT.TokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    a.Config.Name,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte(a.Config.JWT.Secret))
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		a.Logger.Info("failed to generate token", op, err)
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates JWT token and returns token
// If token is not valid - returns error
func ValidateToken(a *server.App, tokenString string) (tokenClaim *MyCustomClaims, err error) {
	const op = "jwt.ValidateToken"

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(a.Config.JWT.Secret), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	//token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(a.Config.JWT.Secret), nil
	//}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		err = fmt.Errorf("%s: %w", op, err)
		a.Logger.Info("failed to validate token", op, err)
	}

	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok {
		a.Logger.Info("failed to validate token", op, err)

		return nil, fmt.Errorf("%s: %w", op, InvalidToken)
	}

	if !token.Valid {
		a.Logger.Info("failed to validate token", op, err)

		return claims, fmt.Errorf("%s: %w", op, InvalidToken)
	}

	return claims, err
}
