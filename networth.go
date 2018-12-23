package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	cmc "github.com/coincircle/go-coinmarketcap"
	"github.com/fatih/color"
)

type Assets struct {
	Assets []Asset `json:"Assets"`
}

type Asset struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Holding string `json:"holding"`
}

func list() error {
	jsonFile, err := os.Open("data/assets.json")
	if err != nil {
		log.Fatalf("Cant open json file \n")
	}
	defer Close(jsonFile)
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var assets Assets
	err = json.Unmarshal(byteValue, &assets)
	if err != nil {
		log.Fatal(err)
	}
	g := color.New(color.FgGreen).Add(color.Bold)
	b := color.New(color.FgBlue).Add(color.Bold)
	c := color.New(color.FgCyan).Add(color.Bold)
	y := color.New(color.FgYellow).Add(color.Bold)
	y.Println("Asset        Holding        Worth        ")
	for i := 0; i < len(assets.Assets); i++ {
		price, err := cmc.Price(&cmc.PriceOptions{
			Symbol:  assets.Assets[i].Symbol,
			Convert: "USD",
		})
		if err != nil {
			log.Fatal(err)
		}
		g.Print(assets.Assets[i].Symbol + "          ")
		b.Print(assets.Assets[i].Holding + "          ")
		h, err := strconv.ParseFloat(assets.Assets[i].Holding, 64)
		c.Print(int(price * h))
		println()
	}
	return nil
}

func add(n string, s string, q string) error {

	fmt.Println(n + " " + s + " " + q)
	return nil
}
