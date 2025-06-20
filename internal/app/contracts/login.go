// Validate login request
package contracts

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Login represents request JSON body format
type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Validate checks if Login is valid
// Returns Login and error if it is not valid
// Login contains fields with errors
func (l *Login) Validate() (validateResult Login, err error) {
	const op = "contracts.Login.Validate"

	validateResult = Login{}

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
