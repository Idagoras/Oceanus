package valid

import (
	"bluesell/src/common"
	"github.com/go-playground/validator/v10"
)

func CurrencyValid(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return common.IsSupportedCurrency(currency)
	}
	return false
}
