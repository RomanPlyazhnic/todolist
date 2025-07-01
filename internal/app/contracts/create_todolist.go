// Validate create todolist request

package contracts

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// TodoList represents a todolist JSON request body format
type TodoList struct {
	UserId     int         `json:"user_id" validate:"required"`
	Text       string      `json:"text"`
	Checkboxes []*Checkbox `json:"checkboxes"`
}

// Checkbox belongs to Todolist
type Checkbox struct {
	Checked bool   `json:"checked" validate:"required"`
	Text    string `json:"text" validate:"required"`
}

func (t *TodoList) Validate() (validateResult TodoList, err error) {
	const op = "contracts.TodoList.Validate"

	validateResult = TodoList{}

	// TODO: return validate result in JSON, which represents request
	validate := validator.New(validator.WithRequiredStructEnabled())
	err = validate.Struct(t)
	if err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				switch e.Field() {
				case "Text":
					if e.Tag() == "required" {
						validateResult.Text = "text is required"
					}
				}
			}

			return validateResult, fmt.Errorf("%s: %w", op, err)
		}
	}

	return validateResult, nil
}
