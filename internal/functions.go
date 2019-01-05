package money

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type Fn func(*Wealth, *logrus.Logger, *cli.Context) (bool, error)
