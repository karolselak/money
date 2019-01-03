package util

import (
	"log"

	cmc "github.com/coincircle/go-coinmarketcap"
)

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
