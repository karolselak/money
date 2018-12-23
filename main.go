package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "networth"
	app.Usage = "track your networth"
	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list all assets",
			Action: func(c *cli.Context) error {
				return list()
			},
		},
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "networth add <asset name> <asset symbol> <asset quantity>",
			Action: func(c *cli.Context) error {
				return add(c.Args().Get(0), c.Args().Get(1), c.Args().Get(2))
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
