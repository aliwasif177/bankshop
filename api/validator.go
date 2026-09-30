package api

import "github.com/go-playground/validator/v10"

func ValidTransferParams(fl validator.FieldLevel) bool {
	req := fl.Parent().Interface().(TransferTxParams)

	return req.FromAccountID != req.ToAccountID
}

var supportedCurrencies = map[string]bool{
	"USD": true,
	"EUR": true,
}

func ValidCurrency(fl validator.FieldLevel) bool {

	currency := fl.Field().String()
	return supportedCurrencies[currency]
}
