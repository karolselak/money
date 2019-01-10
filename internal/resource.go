package money

import (
	"time"

	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Resource interface
type Resource interface {
	Update()
}

// Fn function type
type Fn func(Resource, *logrus.Logger, *cli.Context) (bool, error)

// Prices map contains updated prices
var Prices map[string]float64

// Currencies contains money.Name for cryptocurrencies
var Currencies []string

// Symbols contains money.Symbol for cryptocurrencies
var Symbols []string

// FetchPrices updates prices in Prices
func FetchPrices() {
	var p float64
	var ch []chan float64
	Prices = make(map[string]float64)
	ln := len(Currencies)
	for i := 0; i < ln; i++ {
		ch = append(ch, make(chan float64, 1))
		go cmcAPI(Currencies[i], ch[i])
	}
	for i := 0; i < ln; i++ {
		p = <-ch[i]
		Prices[Symbols[i]] = p
	}
}

func cmcAPI(sym string, c chan float64) {
	c <- util.GetPrice(sym)
	time.Sleep(time.Millisecond * 250)
}
