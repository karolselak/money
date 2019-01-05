package base

import (
	"fmt"

	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
)

func List(w *money.Wealth, log *logrus.Logger) (bool, error) {
	listFiat(w)
	listCrypto(w)
	prntTotal(fmt.Sprintf("%9.3f", w.Worth))
	return false, nil
}
func listFiat(w *money.Wealth) {
	var data [][]string
	var hold float64
	var worth float64
	for j := 0; j < len(w.Wealth[0].Assets); j++ {
		data = append(data, []string{})
		data[j] = append(data[j], w.Wealth[0].Assets[j].Symbol)
		hold = w.Wealth[0].Assets[j].Holding
		data[j] = append(data[j], fmt.Sprintf("%9.3f", hold))
		worth = w.Wealth[0].Assets[j].Worth
		data[j] = append(data[j], fmt.Sprintf("%9.3f", worth))
	}
	prnt(data, "Fiat")
}
func listCrypto(w *money.Wealth) {
	var data [][]string
	var hold float64
	var worth float64
	for j := 0; j < len(w.Wealth[1].Assets); j++ {
		data = append(data, []string{})
		data[j] = append(data[j], w.Wealth[1].Assets[j].Symbol)
		hold = w.Wealth[1].Assets[j].Holding
		data[j] = append(data[j], fmt.Sprintf("%9.3f", hold))
		worth = w.Wealth[1].Assets[j].Worth
		data[j] = append(data[j], fmt.Sprintf("%9.3f", worth))
	}
	prnt(data, "Cryptocurrencies")
}
func Update(w *money.Wealth, log *logrus.Logger) (bool, error) {
	var sum float64
	var hold float64
	var price float64
	var worth float64
	var sym string
	for j := 0; j < len(w.Wealth[1].Assets); j++ {
		hold = w.Wealth[1].Assets[j].Holding
		sym = w.Wealth[1].Assets[j].Symbol
		if sym == "DSH" {
			price = util.GetPrice("DASH")
		} else {
			price = util.GetPrice(sym)
		}
		worth = hold * price
		w.Wealth[1].Assets[j].Worth = worth
		sum += worth
	}
	hold = w.Wealth[0].Assets[0].Holding
	w.Wealth[0].Assets[0].Worth = hold * 0.75
	sum += hold * 0.75
	w.Worth = sum
	return true, nil
}
