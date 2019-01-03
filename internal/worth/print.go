package worth

import (
	"fmt"

	"github.com/mohfunk/money/pkg/util"
)

// listHead prints table Header
func digSp(dig int) {
	for i := (5 - dig); i > 0; i-- {
		fmt.Print(" ")
	}
}
func listHead() {
	util.SetClr()
	util.Bld[0].Print("Asset  ")
	print("                ")
	util.Bld[0].Print("Holding")
	print("                ")
	util.Bld[0].Print("Worth")
	fmt.Println()
}

// listCat
func listCat(cat string) {
	print("       ")
	print("                ")
	util.Bld[5].Print(cat)
	println()

}

// listItem prints asset  row
func listItem(sym string, hld float64, price float64, dig int) {
	util.Bld[2].Print(sym)
	print("    ")
	print("                ")

	util.Bld[3].Printf("%.2f", hld)
	digSp(dig)
	print("               ")
	util.Bld[1].Print(int(price * hld))
	fmt.Println()

}

// listFoot prints table footer
func listFoot(sum float64) {

	fmt.Println()
	util.Und[3].Print("Net Worth: ")
	util.Bld[1].Printf("%.2f", sum)
	util.Und[3].Println(" CAD")
}

/*
// printMod prints asset after modification
func printMod(n string, price float64, preHold float64, curHold float64) {

	Bld[2].Print(n)

	print(" ")

	bldSet(R)
	fmt.Printf("%.2f", preHold)
	unset()

	print(" ")

	Bld[0].Print(int(price * preHold))
	print(" ")
	Bld[1].Print("USD")

	Bld[6].Print("  ->  ")

	BldSet(Y)
	fmt.Printf("%.2f", curHold)
	unset()

	print(" ")

	Bld[2].Print(int(price * curHold))
	print(" ")
	Bld[2].Println("USD")
}

func printConfirm(s string) {
	Und[2].Print(s)
	Uld[6].Println(" Added!")
}
*/
