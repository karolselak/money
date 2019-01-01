package data

import (
	"log"
	"os"
	"os/user"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
)


// Asset stuct
type Config struct {
	dataDir  string
	dataFile string
}

// initalizes configuration
func initConfig() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	ddir := filepath.Join(dir, ".networth")
	jfile := filepath.Join(ddir, "assets.json")
	Conf = &Config{
		dataDir:  ddir,
		dataFile: jfile,
	}
}

func configure() {
	initConfig()
	if _, err := os.Stat(Conf.dataDir); os.IsNotExist(err) {
		os.Mkdir(Conf.dataDir, 0700)
	}
	os.OpenFile(Conf.dataFile, os.O_RDONLY|os.O_CREATE, 0666)
	fi, err := os.Stat(Conf.dataFile)
	if err != nil {
		log.Fatal(err)
	}
	if fi.Size() == 0 {
		initAssets()
	}
}
