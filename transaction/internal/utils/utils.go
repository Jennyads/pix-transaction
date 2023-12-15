package utils

import "github.com/shopspring/decimal"

func ToFloat(decimal decimal.Decimal) float64 {
	f, _ := decimal.Float64()
	return f
}
