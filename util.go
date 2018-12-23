package main

import (
	"io"
	"log"
	"strconv"

	"github.com/fatih/color"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Fatal(err)
	}
}

var G = color.FgGreen
var R = color.FgRed
var Y = color.FgYellow
var C = color.FgCyan
var B = color.FgBlue
var M = color.FgMagenta
var W = color.FgWhite
var BL = color.Bold
var UL = color.Underline
var BG = color.BgBlack
var FG = color.BgWhite

func NumDig(n float64) int {
	return len(strconv.FormatFloat(n, 'f', 0, 64))
}
