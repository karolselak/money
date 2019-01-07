package main

import (
	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Command is a wraps cli.Command with useful structs and data
type Command struct {
	cmd cli.Command
	fp  string
	log *logrus.Logger
	act money.Fn
	res money.Resource
}

func (c *Command) info(n, usg string, ali []string) {

	c.log.Info("Command.info, names")
	c.cmd.Name = n
	c.cmd.Aliases = ali
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
	// close when done
	defer util.Close(file)

	// read the data in the file
	ibytes, err := util.Read(file)
	// Unmarshle json
	err = util.Unmarshal(ibytes, c.res)

	// call the the command's function on the given interface
	wrt, err = c.act(c.res, c.log, cntxt)

	if err != nil {
		c.log.WithFields(logrus.Fields{
			"Command": c.cmd.Name,
			"json":    c.res,
			"write":   wrt,
		}).Fatal("Command Action failed")
	}

	// if the functon has written to the interface, the marshel it.
	if wrt == true {
		wbyte, err := util.Marshal(c.res)

		err = util.Write(wbyte, c.fp)
		if err != nil {
			c.log.WithFields(logrus.Fields{
				"Command": c.cmd.Name,
				"json":    c.res,
				"byte":    wbyte,
				"file":    c.fp,
				"error":   err,
			}).Fatal("Writing bytes to file failed")
		}
	}
	return nil
}
