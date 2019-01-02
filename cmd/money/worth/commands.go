package worth

import (
	"github.com/mohfunk/money/pkg/data"
	"github.com/mohfunk/money/pkg/util"
	log "github.com/sirupsen/logrus"
)

func List(Forte *data.Wealth) error {
	var sum float64
	listHead()

	listCat("Fiat")
	flen := len(Forte.Fiat)
	clen := len(Forte.Crypto)
	mlen := len(Forte.Metals)
	log.WithFields(log.Fields{
		"fiat": flen,
		"crpt": clen,
		"mlen": mlen,
	}).Info("Assets Lengths")
	for i := 0; i < flen; i++ {
		sym := Forte.Fiat[i].Symbol
		hld := Forte.Fiat[i].Holding
		dig := util.NumDig(hld)
		sum += (hld)
		listItem(sym, hld, hld, dig)
	}

	listCat("Cryptocurrencies")
	for i := 0; i < clen; i++ {
		sym := Forte.Crypto[i].Symbol
		price := util.GetPrice(sym)
		hld := Forte.Crypto[i].Holding
		dig := util.NumDig(hld)
		sum += (price * hld)
		listItem(sym, hld, price, dig)
	}

	listCat("Metals")
	for i := 0; i < mlen; i++ {
		sym := Forte.Metals[i].Symbol
		hld := Forte.Metals[i].Holding
		dig := util.NumDig(hld)
		sum += (hld)
		listItem(sym, hld, hld, dig)
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
