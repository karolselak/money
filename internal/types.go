package money

import "github.com/sirupsen/logrus"

type Fn func(*Wealth, *logrus.Logger) (bool, error)
