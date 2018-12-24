package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	cmc "github.com/coincircle/go-coinmarketcap"
	"github.com/fatih/color"
)

// Assets an array of Asset
type Assets struct {
	Assets []Asset `json:"Assets"`
}

// Asset stuct
type Asset struct {
	Name    string  `json:"name"`
	Symbol  string  `json:"symbol"`
	Holding float64 `json:"holding"`
}

func list() error {
	y := color.New(Y).Add(BL)
	g := color.New(G).Add(BL)
	b := color.New(B).Add(BL)
	w := color.New(W).Add(BL)
	y.Print("Asset")
	y.Print("        ")
	y.Print("Holding")
	y.Print("        ")
	y.Print("Worth")
	fmt.Println()
	var sum float64
	for i := 0; i < len(Forte.Assets); i++ {

		price, err := cmc.Price(&cmc.PriceOptions{
			Symbol:  Forte.Assets[i].Symbol,
			Convert: "USD",
		})
		if err != nil {
			log.Fatal(err)
		}
		sym := Forte.Assets[i].Symbol
		hld := Forte.Assets[i].Holding
		dig := NumDig(hld)
		b.Print(sym)
		b.Print("          ")
		color.Set(W, BL)
		fmt.Printf("%.2f", hld)
		for i := (5 - dig); i > 0; i-- {
			fmt.Print(" ")
		}
		color.Unset()
		w.Print("       ")
		g.Print(int(price * Forte.Assets[i].Holding))
		sum += (price * hld)
		fmt.Println()
	}
	y.Print("Net Worth: ")
	g.Printf("%.2f", sum)
	y.Println(" USD")
	return nil
}

func newAsset(n string, s string) error {

	nasset := Asset{
		Name:    n,
		Symbol:  s,
		Holding: 0.0,
	}
	Forte.Assets = append(Forte.Assets, nasset)

	wjson, _ := JSON.Marshal(Forte)
	err := ioutil.WriteFile("data/assets.json", wjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func fund(n string, q string) error {
	hld, err := strconv.ParseFloat(q, 64)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(Forte.Assets); i++ {
		if n == Forte.Assets[i].Symbol {
			Forte.Assets[i].Holding += hld
			break
		}
	}
	wjson, _ := JSON.Marshal(Forte)
	err = ioutil.WriteFile("data/assets.json", wjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
