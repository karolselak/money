package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	cmc "github.com/coincircle/go-coinmarketcap"
	"github.com/fatih/color"
)

type Assets struct {
	Assets []Asset `json:"Assets"`
}

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
		fmt.Printf("%5.2f", hld)
		for i := (5 - dig); i > 0; i-- {
			fmt.Print(" ")
		}
		color.Unset()
		w.Print("       ")
		g.Print(int(price * Forte.Assets[i].Holding))
		fmt.Println()
	}
	return nil
}

func newAsset(n string, s string) error {

	nasset := Asset{
		Name:    n,
		Symbol:  s,
		Holding: 0.0,
	}
	Forte.Assets = append(Forte.Assets, nasset)
	wjson, _ := json.Marshal(Forte)
	err := ioutil.WriteFile("data/assets.json", wjson, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
