package main

import (
	"os"
	"time"

	"github.com/mohfunk/money/internal/worth"
	"github.com/mohfunk/money/pkg/data"
	"github.com/mohfunk/money/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type fn func(a ...string) error
type fnc func(d *data.Wealth) error

// Config global var
var Conf *data.Config

// Forte global var containing assets
var Forte *data.Wealth = &data.Wealth{}

// JSONFile pointer to the assets.json file
var JSONFile *os.File

func execute(f fnc) error {
	Conf = data.Configure()
	JSONFile = util.OpenJSON(Conf.DataFile)
	defer util.Close(JSONFile)
	util.ReadJSON(JSONFile, Forte)
	return f(Forte)
}
func main() {
	app := cli.NewApp()
	log.WithFields(log.Fields{
		"step": "init",
	}).Info("NewApp")
	app.Name = "money"
	app.Usage = "track your finances"
	app.Version = "0.4"
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
			Name:    "worth",
			Aliases: []string{"w"},
			Subcommands: []cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Action: func(c *cli.Context) error {
						return execute(worth.List)
					},
				},
				//	{
				//		Name:    "add",
				//		Aliases: []string{"a"},
				//		Action: func(c *cli.Context) error {
				//			return execute(add, c.Args().Get(0), c.Args().Get(1))
				//		},
				//	},
				//	{
				//		Name:    "modify",
				//		Aliases: []string{"mod", "m"},
				//		Action: func(c *cli.Context) error {
				//			return execute(mod, c.Args().Get(0), c.Args().Get(1), c.Args().Get(2))
				//		},
				//	},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("app.Run failed.")

	}
}
