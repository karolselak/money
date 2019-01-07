package trade

import (
	"fmt"

	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func List(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
	t := r.(*money.Trades)
	t.Update()
	var data [][]string
	var dir string
	var cost float64
	var amnt float64
	var curr float64
	for i := 0; i < len(t.Open); i++ {

		data = append(data, []string{})
		if t.Open[i].Claim == true {
			dir = "<-"
		} else {
			dir = "->"
		}

		data[i] = append(data[i], fmt.Sprintf("%d", i))
		data[i] = append(data[i], t.Open[i].Base+" "+dir+" "+t.Open[i].Invs)
		cost = t.Open[i].Cost
		amnt = t.Open[i].Amount
		curr = t.Open[i].Current
		data[i] = append(data[i], fmt.Sprintf("%3.5f %s %3.5f", cost, dir, amnt))
		data[i] = append(data[i], fmt.Sprintf("%3.5f", curr))
	}
	prnt(data, "open")
	return false, nil

}

func Mod(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
	t := r.(*money.Trades)
	id := int(util.Stf(c.Args().Get(0)))
	narg := c.NArg()
	narg -= 1

	var field string
	var value string
	for i := 1; i < narg-1; i++ {
		if c.Args().Get(i)[:1] == "-" {
			field = c.Args().Get(i)[1:]
			value = c.Args().Get(i + 1)
			if field == "pair" {
				field = "invs"
				value = c.Args().Get(i + 2)
				mod(t, id, field, value)
				field = "base"
				value = c.Args().Get(i + 1)
			}
			mod(t, id, field, value)
		}
	}

	t.Update()
	return true, nil
}

func mod(t *money.Trades, id int, field string, val string) {

	switch field {
	case "base":
		t.Open[id].Base = val
	case "invs":
		t.Open[id].Invs = val
	case "amount":
		t.Open[id].Amount = util.Stf(val)
	case "cost":
		t.Open[id].Cost = util.Stf(val)
	case "dir":
		if val == "buy" {
			t.Open[id].Claim = false
		} else {
			t.Open[id].Claim = true
		}

	}
}

func Add(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
	t := r.(*money.Trades)

	base := c.Args().Get(0)
	invs := c.Args().Get(1)
	clam := c.Args().Get(2)
	cost := c.Args().Get(3)
	amnt := c.Args().Get(4)

	trade := &money.Trade{}
	trade.Base = base
	trade.Invs = invs

	if clam[:1] == "b" {
		trade.Claim = false
	} else {
		trade.Claim = true
	}
	trade.Cost = util.Stf(cost)
	trade.Amount = util.Stf(amnt)
	trade.Buy = util.Stf(amnt) / util.Stf(cost)
	t.Open = append(t.Open, *trade)
	t.Update()
	return true, nil
}
