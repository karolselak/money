package main

import (
	"log"
	"os"
	"time"

	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/internal/base"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type Application struct {
	app    *cli.App
	cmd    *[]Command
	log    *logrus.Logger
	config *money.Config
	wealth *money.Wealth
}

func (a *Application) info() {
	a.app.Name = "money"
	a.app.Usage = "track your finances"
	a.app.Version = "0.0.5"
	a.app.Compiled = time.Now()
	a.app.Copyright = "(c) MIT 2019"
	a.app.Authors = []cli.Author{
		cli.Author{
			Name:  "Mohammed Alhaythm",
			Email: "moh@abstractum.io",
		},
	}
	a.log.Info("app info registered")
}

func (a *Application) setLog() {
	a.log.SetFormatter(&logrus.JSONFormatter{})
	a.log.SetReportCaller(true)
	a.log.Out = os.Stdout
	f, err := os.OpenFile(a.config.LogFile, os.O_CREATE|os.O_WRONLY, 0666)
	_ = f
	if err != nil {
		a.log.Error("Cannot create file .log, logging to stderr instead")
	} else {
		file, err := os.OpenFile(a.config.LogFile, os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			a.log.Info("failed to log to file .log, logging to stderr instead")
		} else {
			a.log.Out = file
		}
	}
}

func (a *Application) init() {
	a.app = cli.NewApp()
	a.log = logrus.New()
	a.config = money.NewConfig()
	a.config.Configure()
	a.wealth = money.NewWealth()
	a.setLog()
	a.log.Info("\n Log set \n")
	a.info()
	a.cmd = a.register()
}

func (a *Application) run() {
	err := a.app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
	a.log.Info("app running")
}

func (a *Application) cmdbasic() *Command {
	c := &Command{}
	c.w = a.wealth
	c.fp = a.config.DataFile
	c.log = a.log

	return c
}
func (a *Application) register() *[]Command {
	list := a.cmdbasic()
	list.act = base.List
	list.info("list", "lists all assets", []string{"ls", "l"})
	list.action()
	c := &[]Command{*list}
	a.app.Commands = cli.Commands{list.cmd}
	return c
}
