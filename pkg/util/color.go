package util

import "github.com/fatih/color"

const (
	G  = color.FgGreen
	R  = color.FgRed
	Y  = color.FgYellow
	C  = color.FgCyan
	B  = color.FgBlue
	M  = color.FgMagenta
	W  = color.FgWhite
	BO = color.Bold
	UL = color.Underline
	BG = color.BgBlack
	FG = color.BgWhite
)

var Clr [7]*color.Color
var Bld [7]*color.Color
var Und [7]*color.Color

func SetClr() {
	Clr[0] = color.New(R)
	Bld[0] = color.New(R).Add(BO)
	Und[0] = color.New(R).Add(UL)
	Clr[1] = color.New(G)
	Bld[1] = color.New(G).Add(BO)
	Und[1] = color.New(G).Add(UL)
	Clr[2] = color.New(B)
	Bld[2] = color.New(B).Add(BO)
	Und[2] = color.New(B).Add(UL)
	Clr[3] = color.New(Y)
	Bld[3] = color.New(Y).Add(BO)
	Und[3] = color.New(Y).Add(UL)
	Clr[4] = color.New(C)
	Bld[4] = color.New(C).Add(BO)
	Und[4] = color.New(C).Add(UL)
	Clr[5] = color.New(M)
	Bld[5] = color.New(M).Add(BO)
	Und[5] = color.New(M).Add(UL)
	Clr[6] = color.New(W)
	Bld[6] = color.New(W).Add(BO)
	Und[6] = color.New(W).Add(UL)
}

func bldSet(c color.Attribute) {
	color.Set(c, BO)
}
func unset() {
	color.Unset()
}
