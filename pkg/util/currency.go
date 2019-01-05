package util

import (
	"fmt"

	"github.com/disiqueira/gocurrency"
	"github.com/shopspring/decimal"
)

// Convert provides currency conversion
func Convert(from, to string, amount float64) float64 {
	f := gocurrency.NewCurrency(from)
	t := gocurrency.NewCurrency(to)
	amt := decimal.NewFromFloat(amount)
	conv, _ := gocurrency.ConvertCurrency(f, t, amt)
	fmt.Println(conv)
	flt, _ := conv.Float64()
	println(flt)
	return flt
}
