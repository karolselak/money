package util

import (
	"io"
	"log"
	"strconv"

	cmc "github.com/coincircle/go-coinmarketcap"
)

// Close closes open resource
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// NumDig returns the numbers of digits
func NumDig(n float64) int {
	return len(strconv.FormatFloat(n, 'f', 0, 64))
}

// getPrice fetches the price of a currency
func GetPrice(sym string) float64 {

	price, err := cmc.Price(&cmc.PriceOptions{
		Symbol:  sym,
		Convert: "CAD",
	})
	if err != nil {
		log.Fatal(err)
	}
	return price
}

// stf converts string to float
func Stf(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal(err)
	}
	return f
}
