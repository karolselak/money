package main

import (
	"log"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/urfave/cli"
)

// Forte global var containing assets
var Forte Assets

// JSONFile pointer to the assets.json file
var JSONFile *os.File

// JSON accessor to json-iterator lib
var JSON = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	app := cli.NewApp()
	app.Name = "networth"
	app.Usage = "track your networth"
	app.Version = "0.2"
	app.Compiled = time.Now()
	app.Copyright = "(c) MIT 2019"
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
				JSONFile = OpenJSON()
				defer Close(JSONFile)
				ReadJSON(JSONFile)
				return list()
			},
		},
		{
			Name:    "newAsset",
			Aliases: []string{"na"},
			Usage:   "networth add <asset name> <asset symbol> <asset quantity>",
			Action: func(c *cli.Context) error {
				JSONFile = OpenJSON()
				defer Close(JSONFile)
				ReadJSON(JSONFile)
				return newAsset(c.Args().Get(0), c.Args().Get(1))
			},
		},
		{

			Name:    "modify",
			Aliases: []string{"mod", "m"},
			Usage:   "networth modify <asset symbol> <sign> <asset quantity>",
			Action: func(c *cli.Context) error {
				JSONFile = OpenJSON()
				defer Close(JSONFile)
				ReadJSON(JSONFile)
				return fund(c.Args().Get(0), c.Args().Get(1), c.Args().Get(2))
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
