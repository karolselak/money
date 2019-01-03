package main

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/mohfunk/money/pkg/util"
)

type Config struct {
	dataDir  string
	dataFile string
}

func (c *Config) configure() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	ddir := filepath.Join(dir, ".money")
	jfile := filepath.Join(ddir, "assets.json")
	bjfile := filepath.Join(dir, "go/src/github.com/mohfunk/money/assets/assets.json")

	c.dataDir = ddir
	c.dataFile = jfile
	if _, err := os.Stat(c.dataDir); os.IsNotExist(err) {
		os.Mkdir(c.dataDir, 0700)
		util.Copy(bjfile, jfile)
	}
}
