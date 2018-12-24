package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
)

// Close closes open resource
func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

// G green
var G = color.FgGreen

// R red
var R = color.FgRed

// Y yellow
var Y = color.FgYellow

// C cyan
var C = color.FgCyan

// B blue
var B = color.FgBlue

// M magenta
var M = color.FgMagenta

// W white
var W = color.FgWhite

// BL bold
var BL = color.Bold

// UL underline
var UL = color.Underline

// BG background color black
var BG = color.BgBlack

// FG background color white
var FG = color.BgWhite

// NumDig returns the numbers of digits
func NumDig(n float64) int {
	return len(strconv.FormatFloat(n, 'f', 0, 64))
}

// OpenJSON returns a pointer to the json file
func OpenJSON() *os.File {

	usr, _ := user.Current()
	dir := usr.HomeDir
	file := filepath.Join(dir, "go/src/github.com/mohfunk/netWorth/data/assets.json")
	JSONFile, err := os.Open(file)
	if err != nil {
		log.Fatal("open failed")
	}
	return JSONFile
}

// ReadJSON unmarshels jsons and return a byte
func ReadJSON(JF *os.File) {
	byteValue, err := ioutil.ReadAll(JF)
	if err != nil {
		log.Fatalf("reading failed")
	}
	err = JSON.Unmarshal(byteValue, &Forte)
	if err != nil {
		log.Fatalf("Unmarshal failed")
	}
}
