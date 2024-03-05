package valid

import (
	"github.com/go-playground/validator/v10"
	"oceanus/src/common"
)

func CurrencyValid(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return common.IsSupportedCurrency(currency)
	}
	return false
}
