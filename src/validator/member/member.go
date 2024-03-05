package valid

import (
	"github.com/go-playground/validator/v10"
)

func NameValid(field validator.FieldLevel) bool {
	fieldValue := field.Field().String()
	if fieldValue == "admin" {
		return false
	}
	return true
}
