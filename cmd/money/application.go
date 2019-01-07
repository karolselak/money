package main

import (
	"log"
	"os"
	"time"

	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/internal/base"
	"github.com/mohfunk/money/internal/trade"
	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Application holds all relevant structs to the main application
type Application struct {
	app    *cli.App
	cmd    *[]Command
	log    *logrus.Logger
	config *money.Config
	wealth money.Resource
	trades money.Resource
}

func (a *Application) info() {
	a.app.Name = "money"
	a.app.Usage = "track your finances"
	a.app.Version = "0.0.5"
	a.app.Compiled = time.Now()
	a.app.Copyright = "(c) MIT 2019"
	a.app.Authors = []cli.Author{
		{
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
	a.wealth = &money.Wealth{}
	a.trades = &money.Trades{}
	a.setLog()
	a.log.Info(" Log set ")
	a.info()
	a.cmd = a.register()
	file, err := util.Open(a.config.DataFile)
	defer util.Close(file)
	ibytes, err := util.Read(file)
	w := &money.Wealth{}
	err = util.Unmarshal(ibytes, w)
	for i := 0; i < len(w.Wealth[1].Assets); i++ {
		money.Currencies = append(money.Currencies, w.Wealth[1].Assets[i].Name)
		money.Symbols = append(money.Symbols, w.Wealth[1].Assets[i].Symbol)
	}
	money.FetchPrices()
	if err != nil {
		a.log.Fatal(err)
	}
}

func (a *Application) run() {
	err := a.app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
	a.log.Info("app running")
}
func (a *Application) register() *[]Command {
	ls := &Command{}
	ls.fp = a.config.DataFile
	ls.log = a.log
	ls.res = a.wealth
	ls.act = base.List

	ad := &Command{}
	ad.fp = a.config.DataFile
	ad.log = a.log
	ls.res = a.wealth
	ad.act = base.Add

	md := &Command{}
	md.fp = a.config.DataFile
	md.log = a.log
	md.res = a.wealth
	md.act = base.Modify

	tr := &Command{}
	tr.fp = a.config.TradeFile
	tr.log = a.log
	tr.res = a.trades
	tr.act = trade.List

	tm := &Command{}
	tm.fp = a.config.TradeFile
	tm.log = a.log
	tm.res = a.trades
	tm.act = trade.Mod

	ta := &Command{}
	ta.fp = a.config.TradeFile
	ta.log = a.log
	ta.res = a.trades
	ta.act = trade.Add

	ls.info("ls", "lists all assets", []string{"l"})
	ls.flag(false)
	ls.action()

	ad.info("ad", "add an asset type", []string{"a"})
	ad.flag(true, "type, t", "c", "specifies asset type")
	ad.action()

	md.info("md", "mod an asset", []string{"m"})
	md.flag(false)
	md.action()

	tr.info("ls-trade", "mod an asset", []string{"tl"})
	tr.flag(false)
	tr.action()

	tm.info("trmod", "mod an asset", []string{"tm"})
	tm.flag(false)
	tm.action()

	ta.info("tradd", "mod an asset", []string{"ta"})
	ta.flag(false)
	ta.action()

	c := &[]Command{*ls, *ad, *md, *tr, *tm, *ta}
	a.log.Info("Commands registered")
	a.app.Commands = cli.Commands{ls.cmd, ad.cmd, md.cmd, tr.cmd, tm.cmd, ta.cmd}
	a.log.Info("cli.Commands registered")
	return c
}
