package main

import (
	"fmt"

	"github.com/fatih/color"
)

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

// listHead prints table Header
func listHead() {
	y := color.New(Y).Add(BL)
	y.Print("Asset")
	y.Print("        ")
	y.Print("Holding")
	y.Print("        ")
	y.Print("Worth")
	fmt.Println()
}

// listItem prints asset  row
func listItem(sym string, hld float64, price float64, dig int) {

	g := color.New(G).Add(BL)
	b := color.New(B).Add(BL)
	w := color.New(W).Add(BL)
	b.Print(sym)
	b.Print("          ")
	color.Set(W, BL)
	fmt.Printf("%.2f", hld)
	for i := (5 - dig); i > 0; i-- {
		fmt.Print(" ")
	}
	color.Unset()
	w.Print("       ")
	g.Print(int(price * hld))
	fmt.Println()

}

// listFoot prints table footer
func listFoot(sum float64) {
	y := color.New(Y).Add(BL)
	g := color.New(G).Add(BL)
	y.Print("Net Worth: ")
	g.Printf("%.2f", sum)
	y.Println(" USD")
}

// printMod prints asset after modification
func printMod(n string, price float64, preHold float64, curHold float64) {
	g := color.New(G).Add(BL)
	w := color.New(W).Add(BL)
	w.Print(n)
	w.Print(" ")
	color.Set(R, BL)
	fmt.Printf("%.2f", preHold)
	w.Print(" ")
	g.Print(int(price * preHold))
	w.Print(" USD")
	color.Unset()
	w.Print("  ->  ")
	color.Set(Y, BL)
	fmt.Printf("%.2f", curHold)
	w.Print(" ")
	g.Print(int(price * curHold))
	w.Println(" USD")
	color.Unset()
}

func printConfirm(s string) {
	b := color.New(B).Add(BL)
	y := color.New(W).Add(BL)
	b.Print(s)
	y.Println(" Added!")
}
