package main

import (
	"log"
	"os"
	"time"

	money "github.com/mohfunk/money/internal"
	base "github.com/mohfunk/money/internal/base"
	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type Application struct {
	app    *cli.App
	log    *logrus.Logger
	config *money.Config
	wealth *money.Wealth
}

func (a *Application) executeAction() error {
	wfile, err := util.Open(a.config.DataFile)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Opening assets file failed")
	}
	defer util.Close(wfile)
	wbytes, err := util.Read(wfile)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Reading assets file failed")
	}
	err = util.Unmarshal(wbytes, a.wealth)
	if err != nil {
		a.log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Unmarshal assets file failed")
	}
	base.List(a.wealth)
	return nil
}

func (a *Application) action() {
	a.app.Action = func(c *cli.Context) error {
		return a.executeAction()
	}
	a.log.Info("App action executed")
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
	f, err := os.OpenFile("log.json", os.O_CREATE|os.O_WRONLY, 0666)
	_ = f
	if err != nil {
		a.log.Error("Cannot create file .log, logging to stderr instead")
	} else {
		file, err := os.OpenFile("log.json", os.O_APPEND|os.O_WRONLY, 0666)
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
	a.action()
}

func (a *Application) run() {
	err := a.app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
	a.log.Info("app running")
}
