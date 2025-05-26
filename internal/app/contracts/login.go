package contracts

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// LoginRequest represents request JSON body format
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// LoginResult represents fields and errors
type LoginResult struct {
	Username string
	Password string
}

// Validate checks if LoginRequest is valid
// Returns LoginResult and error if it is not valid
// LoginResult contains fields with errors
func (l *LoginRequest) Validate() (validateResult LoginResult, err error) {
	const op = "contracts.Login.Validate"

	validateResult = LoginResult{}

	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(l)
	if err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				switch e.Field() {
				case "Username":
					if e.Tag() == "required" {
						validateResult.Username = "username is required"
					}
				case "Password":
					if e.Tag() == "required" {
						validateResult.Password = "password is required"
					}
				}
			}

			return validateResult, fmt.Errorf("%s: %w", op, err)
		}
	}

	return validateResult, nil
}
