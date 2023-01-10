package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/mbaxamb3333/simplebank/util"
)

var validCurrency validator.Func = func(fieldLever validator.FieldLevel) bool {
	if currency, ok := fieldLever.Field().Interface().(string); ok {
		// check if currency is valid
		return util.IsSupportedCurrency(currency)
	}
	return false
}
