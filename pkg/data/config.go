package data

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// Asset stuct
type Config struct {
	DataDir  string
	DataFile string
}

func Configure() *Config {
	usr, _ := user.Current()
	dir := usr.HomeDir
	ddir := filepath.Join(dir, ".networth")
	jfile := filepath.Join(ddir, "assets.json")
	c := &Config{
		DataDir:  ddir,
		DataFile: jfile,
	}
	if _, err := os.Stat(c.DataDir); os.IsNotExist(err) {
		os.Mkdir(c.DataDir, 0700)
	}
	os.OpenFile(c.DataFile, os.O_RDONLY|os.O_CREATE, 0666)
	fi, err := os.Stat(c.DataFile)
	if err != nil {
		log.Fatal(err)
	}
	if fi.Size() == 0 {
		return nil
	}
	return c
}
