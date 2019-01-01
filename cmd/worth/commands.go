package worth

func list(Forte Assets) error {
	var sum float64
	leng := len(Forte.Assets)

	if leng == 0 {
		println("You have no assets, run networth add <SYMBOL> <NAME>")
		return nil
	}
	listHead()
	for i := 0; i < leng; i++ {
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
	writeJSON(*Forte)
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
