package data

import (
	"fmt"
	"io"
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
	ddir := filepath.Join(dir, ".money")
	jfile := filepath.Join(ddir, "assets.json")
	bjfile := filepath.Join(dir, "go/src/github.com/mohfunk/money/assets/assets.json")
	c := &Config{
		DataDir:  ddir,
		DataFile: jfile,
	}
	if _, err := os.Stat(c.DataDir); os.IsNotExist(err) {
		os.Mkdir(c.DataDir, 0700)
		copy(bjfile, jfile)
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

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}
