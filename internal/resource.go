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

// Prices
var Prices []float64
var Currencies []string

// FetchPrices updates prices in Prices
func FetchPrices() {
	var ch []chan float64
	var p []float64
	ln := len(Currencies)
	for i := 0; i < ln; i++ {
		ch = append(ch, make(chan float64, 1))
		go cmcAPI(Currencies[i], ch[i])
	}
	for i := 0; i < ln; i++ {
		p = append(p, <-ch[i])
	}
}

func cmcAPI(sym string, c chan float64) {
	c <- util.GetPrice(sym)
	time.Sleep(time.Millisecond * 100)
}
