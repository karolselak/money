package base

import (
	"errors"
	"fmt"

	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// List command prints all assets
func List(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
	w := r.(*money.Wealth)
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
	var sum float64
	for j := 0; j < len(w.Wealth[1].Assets); j++ {
		data = append(data, []string{})
		data[j] = append(data[j], w.Wealth[1].Assets[j].Symbol)
		hold = w.Wealth[1].Assets[j].Holding
		data[j] = append(data[j], fmt.Sprintf("%9.3f", hold))
		worth = w.Wealth[1].Assets[j].Worth
		data[j] = append(data[j], fmt.Sprintf("%9.3f", worth))
		sum += worth
	}
	prnt(data, "Cryptocurrencies")
	prntTotal(fmt.Sprintf("%9.3f", sum))
}

// Update command fetches prices and rates

// Add creates a new Asset type
func Add(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
	w := r.(*money.Wealth)
	name := c.Args().Get(0)
	sym := c.Args().Get(1)
	amnt := util.Stf(c.Args().Get(2))
	pr := util.GetPriceV2(name)
	wor := pr * amnt
	a := money.Asset{
		Name:    name,
		Symbol:  sym,
		Holding: amnt,
		Worth:   wor,
	}
	w.Wealth[1].Assets = append(w.Wealth[1].Assets, a)
	println("Asset %s Added!", a.Symbol)
	w.Update()
	return true, nil
}

// Modify changes the Asset.Holding filed of a given Asset
func Modify(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {

	w := r.(*money.Wealth)
	syml := c.Args().Get(0)
	sign := c.Args().Get(1)
	amnt := c.Args().Get(2)

	var sym string
	for i := 0; i < 2; i++ {
		for j := 0; j < len(w.Wealth[i].Assets); j++ {
			sym = w.Wealth[i].Assets[j].Symbol
			if sym == syml {
				if sign == "+" {
					w.Wealth[i].Assets[j].Holding += util.Stf(amnt)
				} else if sign == "-" {
					w.Wealth[i].Assets[j].Holding -= util.Stf(amnt)
				} else {
					return false, errors.New("wrong sign")
				}
			}
		}
	}
	w.Update()
	return true, nil
}
