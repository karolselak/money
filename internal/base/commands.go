package base

import (
	"fmt"

	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/pkg/util"
)

func List(w *money.Wealth) error {
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
				price = util.GetPriceV2(w.Wealth[i].Assets[j].Name)
			} else {
				price = 1
			}
			data[ind] = append(data[ind], fmt.Sprintf("%f", hold))
			data[ind] = append(data[ind], fmt.Sprintf("%f", hold*price))
			sum += hold * price
		}
		ind++
	}
	prnt(data, sum)
	return nil
}
