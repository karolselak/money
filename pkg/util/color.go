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

var Clr []*color.Color
var Bld []*color.Color
var Und []*color.Color

func setClr() {
	Clr[0] = color.New(R)
	Clr[1] = color.New(G)
	Clr[2] = color.New(B)
	Clr[3] = color.New(Y)
	Clr[4] = color.New(C)
	Clr[5] = color.New(M)
	Clr[6] = color.New(W)
	for i := 0; i < 7; i++ {
		Bld[i] = Clr[i].Add(BO)
		Und[i] = Clr[i].Add(UL)
	}
}

func bldSet(c color.Atrribute) {
	color.Set(c, BO)
}
func unset() {
	color.Unset()
}
