package main

func list() error {
	var sum float64
	listHead()
	for i := 0; i < len(Forte.Assets); i++ {
		sym := Forte.Assets[i].Symbol
		price := getPrice(sym)
		hld := Forte.Assets[i].Holding
		dig := NumDig(hld)
		sum += (price * hld)
		listItem(sym, hld, price, dig)
	}
	listFoot(sum)
	return nil
}

func newAsset(n string, s string) error {
	nasset := createAsset(n, s)
	Forte.Assets = append(Forte.Assets, nasset)
	writeJSON(Forte)
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

	writeJSON(Forte)
	price := getPrice(n)
	printMod(n, price, preHold, curHold)
	return nil
}
