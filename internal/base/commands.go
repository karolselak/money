package base

import (
	"errors"
	"fmt"
	"time"

	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func List(w *money.Wealth, log *logrus.Logger, c *cli.Context) (bool, error) {
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
func Update(w *money.Wealth, log *logrus.Logger, c *cli.Context) (bool, error) {
	var sum float64
	var hold float64
	var sym string
	var p float64
	var wor float64
	var ch []chan float64

	ln := len(w.Wealth[1].Assets)

	for i := 0; i < ln; i++ {
		ch = append(ch, make(chan float64, 1))
		sym = w.Wealth[1].Assets[i].Name
		go cmcApi(sym, ch[i])
	}

	for i := 0; i < ln; i++ {

		log.WithFields(logrus.Fields{
			"index":  i,
			"wealth": w.Worth,
			"price":  p,
		}).Info("getting price")
		p = <-ch[i]
		hold = w.Wealth[1].Assets[i].Holding
		wor = p * hold
		w.Wealth[1].Assets[i].Worth = wor
	}
	log.WithFields(logrus.Fields{
		"woryh":  wor,
		"wealth": w.Worth,
	}).Info(" worth")

	hold = w.Wealth[0].Assets[0].Holding
	w.Wealth[0].Assets[0].Worth = hold * 0.75

	sum += hold * 0.75
	w.Worth = sum

	log.WithFields(logrus.Fields{
		"sum":    sum,
		"wealth": w.Worth,
	}).Info("Computed worth, command Update done")
	return true, nil
}

func Add(w *money.Wealth, log *logrus.Logger, c *cli.Context) (bool, error) {
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
	return Update(w, log, c)
}

func Modify(w *money.Wealth, log *logrus.Logger, c *cli.Context) (bool, error) {

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
	return Update(w, log, c)
}

func cmcApi(sym string, c chan float64) {
	c <- util.GetPrice(sym)
	time.Sleep(time.Millisecond * 100)
}
