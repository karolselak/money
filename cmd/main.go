package main

import (
	"log"
	"os"
	"time"

	"github.com/mohfunk/money/cmd/worth"
	"github.com/mohfunk/money/pkg/data"
	"github.com/urfave/cli"
)

type fn func(a ...string) error

// Config global var
var Conf *Config

// Forte global var containing assets
var Forte *Assets

// JSONFile pointer to the assets.json file
var JSONFile *os.File


func execute(f fn, a ...string) error {
	configure()
	JSONFile = OpenJSON()
	defer Close(JSONFile)
	ReadJSON(JSONFile)
	return f(a)
}
func main() {
	app := cli.NewApp()
	app.Name = "money"
	app.Usage = "track your finances"
	app.Version = "0.3"
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
			Name:	 "worth",
			Aliases: []string{"w"},
			Subcommands: []cli.Command{
				{
					Name:	 "list",
					Aliases: []string{"l"},
					Action:  
					func(c *cli.Context) error {
						return execute(worth.list)
					},
				},
				{
					Name:    "add",
					Aliases: []string{"a"},
					Action: 
					func(c *cli.Context) error {
						return execute(worth.add, c.Args().Get(0), c.Args().Get(1))
					},
				},
				{
					Name:    "modify",
					Aliases: []string{"mod", "m"},
					Action: 
					func(c *cli.Context) error {
						return return(worth.mod, c.Args().Get(0), c.Args().Get(1), c.Args().Get(2))
					},
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
