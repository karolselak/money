package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

var Forte Assets
var JsonFile *os.File

func main() {
	app := cli.NewApp()
	app.Name = "networth"
	app.Usage = "track your networth"
	app.Version = "0.1"
	app.Compiled = time.Now()
	app.Copyright = "(c) MIT 2018"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Mohammed Alhaythm",
			Email: "moh@abstractum.io",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "list",
			Aliases:   []string{"l"},
			Usage:     "networth list",
			UsageText: "newtorth list - lists all assets",
			ArgsUsage: "",
			Action: func(c *cli.Context) error {
				JsonFile, err := os.Open("data/assets.json")
				if err != nil {
					log.Fatalf("Cant open json file \n")
				}
				defer Close(JsonFile)
				byteValue, err := ioutil.ReadAll(JsonFile)
				if err != nil {
					log.Fatal(err)
				}
				err = json.Unmarshal(byteValue, &Forte)
				if err != nil {
					log.Fatal(err)
				}
				return list()
			},
		},
		{
			Name:    "newAsset",
			Aliases: []string{"na"},
			Usage:   "networth add <asset name> <asset symbol> <asset quantity>",
			Action: func(c *cli.Context) error {
				JsonFile, err := os.Open("data/assets.json")
				if err != nil {
					log.Fatalf("Cant open json file \n")
				}
				defer Close(JsonFile)
				byteValue, err := ioutil.ReadAll(JsonFile)
				if err != nil {
					log.Fatal(err)
				}
				err = json.Unmarshal(byteValue, &Forte)
				if err != nil {
					log.Fatal(err)
				}
				return newAsset(c.Args().Get(0), c.Args().Get(1))
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
