// Validate register request

package contracts

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// RegisterRequest represents request JSON body format
type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

// RegisterResult represents fields and errors
type RegisterResult struct {
	Username string
	Password string
}

// Validate checks if RegisterRequest is valid
// Returns RegisterRequest and error if it is not valid
// RegisterRequest contains fields with errors
func (l *RegisterRequest) Validate() (validateResult RegisterResult, err error) {
	const op = "contracts.Login.Validate"

	validateResult = RegisterResult{}

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
					switch e.Tag() {
					case "required":
						validateResult.Password = "password is required"
					case "min":
						validateResult.Password = "password must be at least 8"
					}
				}
			}

			return validateResult, fmt.Errorf("%s: %w", op, err)
		}
	}

	return validateResult, nil
}
