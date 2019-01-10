package main

import (
	"log"
	"os"
	"time"

	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/internal/budget"
	"github.com/mohfunk/money/internal/trades"
	"github.com/mohfunk/money/internal/wealth"
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
	budget money.Resource
}

func (a *Application) info() {
	a.app.Name = "money"
	a.app.Usage = "track your finances"
	a.app.Version = "0.0.7"
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
	a.budget = &money.Budget{}
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

	w := &Command{}
	t := &Command{}
	b := &Command{}

	w.fp = a.config.DataFile
	t.fp = a.config.TradeFile
	b.fp = a.config.BudgetFile

	w.log = a.log
	t.log = a.log
	b.log = a.log

	w.res = a.wealth
	t.res = a.trades
	b.res = a.budget

	w.act = wealth.List
	t.act = trades.List
	b.act = budget.List

	w.info("wealth", "lists all assets", []string{"w"})
	t.info("trades", "lists all trades", []string{"t"})
	b.info("budget", "lists budget", []string{"b"})

	w.action()
	t.action()
	b.action()

	ad := &Command{}
	md := &Command{}
	rm := &Command{}

	ad.fp = a.config.DataFile
	md.fp = a.config.DataFile
	rm.fp = a.config.DataFile

	ad.log = a.log
	md.log = a.log
	rm.log = a.log

	ad.res = a.wealth
	md.res = a.wealth
	rm.res = a.wealth

	ad.act = wealth.Add
	ad.info("add", "add an asset type", []string{"a"})
	ad.action()

	md.act = wealth.Modify
	md.info("modify", "modify an asset", []string{"m"})
	md.action()

	rm.act = wealth.Remove
	rm.info("remove", "removes an asset", []string{"r"})
	rm.action()
	w.cmd.Subcommands = cli.Commands{ad.cmd, md.cmd, rm.cmd}

	adt := &Command{}
	adt.fp = a.config.TradeFile
	adt.log = a.log
	adt.res = a.trades
	adt.act = trades.Add
	adt.info("add", "add a trade", []string{"a"})
	adt.action()

	cls := &Command{}
	cls.fp = a.config.TradeFile
	cls.log = a.log
	cls.res = a.trades
	cls.act = trades.Close
	cls.info("close", "Close a trade", []string{"c"})
	cls.action()
	t.cmd.Subcommands = cli.Commands{adt.cmd, cls.cmd}
	/*

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

		tc := &Command{}
		tc.fp = a.config.TradeFile
		tc.log = a.log
		tc.res = a.trades
		tc.act = trade.BList

		tac := &Command{}
		tac.fp = a.config.TradeFile
		tac.log = a.log
		tac.res = a.trades
		tac.act = trade.Close

		ls.info("ls", "lists all assets", []string{"l"})
		ls.action()



		tr.info("ls-trade", "mod an asset", []string{"tl"})
		tr.action()

		tm.info("trmod", "mod an asset", []string{"tm"})
		tm.action()

		ta.info("tradd", "mod an asset", []string{"ta"})
		ta.action()

		tc.info("lsctr", "mod an asset", []string{"lc"})
		tc.action()

		tac.info("tac", "mod an asset", []string{"tc"})
		tac.action()

	*/
	c := &[]Command{*w, *t, *b}
	a.log.Info("Commands registered")
	a.app.Commands = cli.Commands{w.cmd, t.cmd, b.cmd}
	a.log.Info("cli.Commands registered")
	return c
}
