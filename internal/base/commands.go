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
func Update(w *money.Wealth) error {

	var sum float64
	var data [][]string
	var hold float64
	var price float64
	ind := 0
	for i := 0; i < 2; i++ {
		data = append(data, []string{})
		data[ind] = append(data[ind], "")
		data[ind] = append(data[ind], "")
		data[ind] = append(data[ind], "")
		for j := 0; j < len(w.Wealth[i].Assets); j++ {
			ind++
			data = append(data, []string{})
			data[ind] = append(data[ind], w.Wealth[i].Assets[j].Symbol)
			hold = w.Wealth[i].Assets[j].Holding
			if w.Wealth[i].Type == "Crypto" {
				price = util.GetPrice(w.Wealth[i].Assets[j].Symbol)
			} else {
				price = 1
			}
			data[ind] = append(data[ind], fmt.Sprintf("%9.2f", hold))
			data[ind] = append(data[ind], fmt.Sprintf("%9.2f", hold*price))
			sum += hold * price
		}
		ind++
	}
	return nil
}
