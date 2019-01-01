package worth

import (
	"fmt"
)

// listHead prints table Header
func sepCol() {
	print("                ")
}
func digSp(dig int) {
	for i := (5 - dig); i > 0; i-- {
		fmt.Print(" ")
	}
}
func listHead() {
	Bld[0].Print("Asset")
	sepCol()
	Bld[0].Print("Holding")
	sepCol()
	Bld[0].Print("Worth")
	fmt.Println()
}

// listItem prints asset  row
func listItem(sym string, hld float64, price float64, dig int) {

	Bld[2].Print(sym)

	sepCol()

	bldSet(3)
	fmt.Printf("%.2f", hld)
	unset()

	sepCol()

	digSp(dig)
	sepCol()
	Bld[1].Print(int(price * hld))
	fmt.Println()

}

// listFoot prints table footer
func listFoot(sum float64) {
	Und[3].Print("Net Worth: ")
	Und[1].Printf("%.2f", sum)
	Und[3].Println(" USD")
}

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
