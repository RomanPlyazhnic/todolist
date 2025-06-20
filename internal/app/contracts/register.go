// Validate register request

package contracts

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// Register represents request JSON body format
type Register struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

// Validate checks if Register is valid
// Returns Register and error if it is not valid
// Register contains fields with errors
func (l *Register) Validate() (validateResult Register, err error) {
	const op = "contracts.Login.Validate"

	validateResult = Register{}

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
