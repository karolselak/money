package util

import (
	"log"
	"os"

	"github.com/anaskhan96/soup"
	cmc "github.com/coincircle/go-coinmarketcap"
)

// getPrice fetches the price of a currency
func GetPrice(sym string) float64 {

	price, err := cmc.Price(&cmc.PriceOptions{
		Symbol:  sym,
		Convert: "USD",
	})
	if err != nil {
		log.Fatal(err)
	}
	return price
}

func GetPriceV2(name string) float64 {
	url := "https://coinmarketcap.com/currencies/" + name
	resp, err := soup.Get(url)
	if err != nil {
		os.Exit(1)
	}
	doc := soup.HTMLParse(resp)
	ptext := doc.Find("div", "class", "details-panel-item--price").Find("span", "class", "details-panel-item--price__value")
	price := Stf(ptext.Text())
	return price
}
