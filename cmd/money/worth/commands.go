package worth

import (
	"github.com/mohfunk/money/pkg/data"
	"github.com/mohfunk/money/pkg/util"
)

func List(Forte *data.Assets) error {
	var sum float64
	leng := len(Forte.Assets)

	if leng == 0 {
		println("You have no assets, run networth add <SYMBOL> <NAME>")
		return nil
	}
	listHead()
	for i := 0; i < leng; i++ {
		sym := Forte.Assets[i].Symbol
		price := util.GetPrice(sym)
		hld := Forte.Assets[i].Holding
		dig := util.NumDig(hld)
		sum += (price * hld)
		listItem(sym, hld, price, dig)
	}
	listFoot(sum)
	return nil
}

/*
func add(n string, s string) error {
	nasset := data.CreateAsset(n, s)
	Forte.Assets = append(Forte.Assets, nasset)
	util.WriteJSON(Conf.DataFile, Forte)
	printConfirm(s)
	return nil
}

func mod(n string, s string, q string) error {
	hld := stf(q)
	var preHold, curHold float64
	for i := 0; i < len(Forte.Assets); i++ {
		if n == Forte.Assets[i].Symbol {
			preHold = Forte.Assets[i].Holding
			if s == "+" {
				Forte.Assets[i].Holding += hld
			} else if s == "-" {
				Forte.Assets[i].Holding -= hld
			}
			curHold = Forte.Assets[i].Holding
			break
		}
	}

	writeJSON(*Forte)
	price := getPrice(n)
	printMod(n, price, preHold, curHold)
	return nil
}
*/
