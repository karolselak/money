package color

import (
	"github.com/fatih/color"
)

const (
	g = color.FgGreen
	r = color.FgRed
	y = color.FgYellow
	c = color.FgCyan
	b = color.FgBlue
	m = color.FgMagenta
	w = color.FgWhite

	bo = color.Bold
	ul = color.Underline
)

var base [7]*color.Color
var C [7]*color.Color
var B [7]*color.Color
var U [7]*color.Color

func SetColors() {
	base[0] = color.New(r)
	base[1] = color.New(g)
	base[2] = color.New(b)

	base[3] = color.New(y)
	base[4] = color.New(c)
	base[5] = color.New(m)

	base[6] = color.New(w)
	var tmp color.Color
	for i := 0; i < 7; i++ {
		tmp = *base[i]
		C[i] = color.New()
		*C[i] = tmp
		B[i] = color.New()
		*B[i] = tmp
		B[i] = B[i].Add(bo)
		U[i] = color.New()
		*U[i] = tmp
		U[i] = U[i].Add(ul)
	}
}
