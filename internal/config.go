package money

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/mohfunk/money/pkg/util"
)

type Config struct {
	DataDir  string
	DataFile string
	LogFile  string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Configure() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	ddir := filepath.Join(dir, ".money")
	jfile := filepath.Join(ddir, "assets.json")
	bjfile := filepath.Join(dir, "go/src/github.com/mohfunk/money/base/assets.json")
	logfile := filepath.Join(dir, "go/src/github.com/mohfunk/money/log.json")

	c.DataDir = ddir
	c.DataFile = jfile
	c.LogFile = logfile
	if _, err := os.Stat(c.DataDir); os.IsNotExist(err) {
		os.Mkdir(c.DataDir, 0700)
	}
	if _, err := os.Stat(c.DataFile); os.IsNotExist(err) {
		util.Copy(bjfile, jfile)
	}
}
