package main

import (
	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// ICommand interface define the 3 basic function signtures for Commands
type ICommand interface {
	info(string, string, []string)
	flag(bool, ...string)
	action()
	execute(*cli.Context) error
}

// Command is a wraps cli.Command with useful structs and data
type Command struct {
	cmd cli.Command
	act money.Fn
	w   *money.Wealth
	fp  string
	log *logrus.Logger
}

func (c *Command) info(n, usg string, ali []string) {

	c.log.Info("Command.info, names")
	c.cmd.Name = n
	c.cmd.Aliases = ali
	c.cmd.Usage = usg
	c.cmd.Usage = usg
	c.log.Info("Command.info, names, done")
}
func (c *Command) flag(include bool, finfo ...string) {
	if include == true {
		f := []cli.Flag{
			cli.StringFlag{
				Name:  finfo[0],
				Usage: finfo[1],
				Value: finfo[2],
			},
		}
		c.cmd.Flags = f
	}
}
func (c *Command) action() {
	c.log.Info("Command.action")
	c.cmd.Action = func(cntxt *cli.Context) error {
		c.log.Info("Command.action returning exec")
		return c.execute(cntxt)
	}
}

func (c *Command) execute(cntxt *cli.Context) error {
	var wrt bool
	// open the file
	file, err := util.Open(c.fp)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Opening assets file failed")
	}
	// close when done
	defer util.Close(file)

	// read the data in the file
	ibytes, err := util.Read(file)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Reading assets file failed")
	}

	// Unmarshle json
	err = util.Unmarshal(ibytes, c.w)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Unmarshal assets file failed")
	}

	c.log.WithFields(logrus.Fields{
		"Command": c.cmd.Name,
		"bytes":   string(ibytes),
		"json":    *c.w,
	}).Info("Unmarshal success, calling action")

	// call the the command's function on the given interface
	wrt, err = c.act(c.w, c.log, cntxt)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"Command": c.cmd.Name,
			"json":    *c.w,
			"write":   wrt,
		}).Fatal("Command Action failed")
	}
	c.log.WithFields(logrus.Fields{
		"Command": c.cmd.Name,
		"json":    *c.w,
		"write":   wrt,
	}).Info("Action success")

	// if the functon has written to the interface, the marshel it.
	if wrt == true {
		wbyte, err := util.Marshal(c.w)
		if err != nil {
			c.log.WithFields(logrus.Fields{
				"error": err,
			}).Fatal("Cannot write to json")
		}
		c.log.WithFields(logrus.Fields{
			"Command": c.cmd.Name,
			"json":    *c.w,
			"byte":    wbyte,
			"file":    c.fp,
		}).Info("Marshaling json success, writing now")
		err = util.Write(wbyte, c.fp)
		if err != nil {
			c.log.WithFields(logrus.Fields{
				"Command": c.cmd.Name,
				"json":    *c.w,
				"byte":    wbyte,
				"file":    c.fp,
				"error":   err,
			}).Fatal("Writing bytes to file failed")
		}
	}
	return nil
}
