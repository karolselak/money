package main

import (
	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type ICommand interface {
	info(string, string, []string)
	action()
	execute(*cli.Context) error
}

type Command struct {
	cmd cli.Command
	act money.Fn
	w   *money.Wealth
	fp  string
	log *logrus.Logger
}

func (c *Command) info(n, usg string, ali []string) {
	c.cmd.Name = n
	c.cmd.Aliases = ali
	c.cmd.Usage = usg
}

func (c *Command) action() {
	c.cmd.Action = func(cntxt *cli.Context) error {
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

	// call the the command's function on the given interface
	wrt, err = c.act(c.w, c.log)

	// if the functon has written to the interface, the marshel it.
	if wrt == true {
		wbyte, err := util.Marshal(c.w)
		if err != nil {
			c.log.WithFields(logrus.Fields{
				"error": err,
			}).Fatal("Cannot write to json")
		}
		util.Write(wbyte, c.fp)
	}
	return nil
}
