package money

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// Fn is holds a generic Command signture
type Fn func(*Wealth, *logrus.Logger, *cli.Context) (bool, error)
