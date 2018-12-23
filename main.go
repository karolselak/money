package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "networth"
	app.Usage = "track your networth"
	app.Action = func(c *cli.Context) error {
		jsonFile, err := os.Open("data/assets.json")
		if err != nil {
			log.Fatalf("Cant open json file \n")
		}
		defer Close(jsonFile)
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			log.Fatal(err)
		}
		var coins Coins
		err = json.Unmarshal(byteValue, &coins)
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(coins.Coins); i++ {
			fmt.Println(coins.Coins[i].Symbol + " " + coins.Coins[i].Holding)
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
