package trades

import (
	"fmt"

	money "github.com/mohfunk/money/internal"
	"github.com/mohfunk/money/pkg/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func List(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {

	listOpen(r)
	listClose(r)
	return false, nil

}
func listOpen(r money.Resource) {
	t := r.(*money.Trades)
	var data [][]string
	var dir string
	var cost float64
	var buy float64
	var amnt float64
	var currp float64
	var profit float64
	var percent float64
	var curr float64
	ind := 0
	for i := 0; i < len(t.Pairs); i++ {
		for j := 0; j < len(t.Pairs[i].Open); j++ {
			if t.Pairs[i].Open[j].Claim == true {
				dir = "<-"
			} else {
				dir = "->"
			}
			cost = t.Pairs[i].Open[j].Cost
			amnt = t.Pairs[i].Open[j].Amount
			buy = t.Pairs[i].Open[j].Buy
			t.Update()
			curr = t.Pairs[i].Open[j].Current
			profit = t.Pairs[i].Open[j].Profit
			percent = t.Pairs[i].Open[j].Percent
			currp = curr / amnt
			data = append(data, []string{})
			data[ind] = append(data[ind], fmt.Sprintf("%d", j))
			data[ind] = append(data[ind], t.Pairs[i].Base+" "+dir+" "+t.Pairs[i].Invs)
			data[ind] = append(data[ind], fmt.Sprintf("%3.5f %s %3.5f", cost, dir, amnt))
			data[ind] = append(data[ind], fmt.Sprintf("%3.5f | %3.5f at %3.5f", buy, curr, currp))
			data[ind] = append(data[ind], fmt.Sprintf("%3.1f USD (%3.1f)", profit, percent))
			ind++
		}
	}
	prnt(data, "open")
}

func listClose(r money.Resource) {
	t := r.(*money.Trades)
	var data [][]string
	var dir string
	var cost float64
	var sell float64
	var buy float64
	var amnt float64
	var profit float64
	var percent float64
	ind := 0
	for i := 0; i < len(t.Pairs); i++ {
		for j := 0; j < len(t.Pairs[i].Close); j++ {
			if t.Pairs[i].Close[j].Claim == true {
				dir = "<-"
			} else {
				dir = "->"
			}
			cost = t.Pairs[i].Close[j].Cost
			amnt = t.Pairs[i].Close[j].Amount
			sell = t.Pairs[i].Close[j].Sell
			buy = t.Pairs[i].Close[j].Buy
			profit = t.Pairs[i].Close[j].Profit
			percent = t.Pairs[i].Close[j].Percent
			data = append(data, []string{})
			data[ind] = append(data[ind], fmt.Sprintf("%d", j))
			data[ind] = append(data[ind], t.Pairs[i].Base+" "+dir+" "+t.Pairs[i].Invs)
			data[ind] = append(data[ind], fmt.Sprintf("%3.5f %s %3.5f", cost, dir, amnt))
			data[ind] = append(data[ind], fmt.Sprintf("%3.5f | %3.5f", buy, sell))
			data[ind] = append(data[ind], fmt.Sprintf("%3.5f (%3.1f)", profit, percent))
			ind++
		}
	}
	prnt(data, "close")
}

/*
func Mod(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
	t := r.(*money.Trades)
	id := int(util.Stf(c.Args().Get(0)))
	narg := c.NArg()
	narg -= 1
	t.Pairs[id].Base = c.Args().Get(1)
	t.Pairs[id].Invs = c.Args().Get(2)
	t.Pairs[id].Cost = util.Stf(c.Args().Get(3))
	t.Pairs[id].Amount = util.Stf(c.Args().Get(4))
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
*/
func Add(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
	t := r.(*money.Trades)

	base := c.Args().Get(0)
	invs := c.Args().Get(1)
	clam := c.Args().Get(2)
	cost := c.Args().Get(3)
	amnt := c.Args().Get(4)
	exists := false
	id := 0
	for i := 0; i < len(t.Pairs); i++ {
		if base == t.Pairs[i].Base && invs == t.Pairs[i].Invs {
			exists = true
			println("exists")
			id = i + 1
			break
		}
	}

	if exists == false {
		println("does not exists")
		tr := []money.Trade{}
		ctr := []money.Trade{}
		pair := money.Pair{
			Base:  base,
			Invs:  invs,
			Open:  tr,
			Close: ctr,
		}
		t.Pairs = append(t.Pairs, pair)
		id = len(t.Pairs)
	}
	trade := &money.Trade{}
	if clam[:1] == "b" {
		trade.Claim = false
	} else {
		trade.Claim = true
	}
	trade.Cost = util.Stf(cost)
	trade.Amount = util.Stf(amnt)
	trade.Buy = util.Stf(cost) / util.Stf(amnt)
	t.Pairs[id-1].Open = append(t.Pairs[id-1].Open, *trade)
	return true, nil
}

/*
func Fill(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
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
	return true, nil
}
func BList(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
	t := r.(*money.Trades)
	var data [][]string
	var dir string
	var cost float64
	var prof float64
	var perc float64
	var sign float64
	for i := 0; i < len(t.Closed); i++ {
		data = append(data, []string{})
		if t.Closed[i].Claim == true {
			dir = "<-"
		} else {
			dir = "->"
		}

		data[i] = append(data[i], fmt.Sprintf("%d", i))
		data[i] = append(data[i], t.Closed[i].Base+" "+dir+" "+t.Closed[i].Invs)
		cost = t.Closed[i].Cost
		prof = t.Closed[i].Profit
		if cost > prof {
			sign = -1.0
		} else {
			sign = 1.0
		}
		perc = (((prof / cost) * 100) * sign)
		data[i] = append(data[i], fmt.Sprintf("%3.5f %s %3.5f", cost, dir, prof))
		data[i] = append(data[i], fmt.Sprintf("%3.5f", perc))
	}
	prntc(data, "Closed")
	return false, nil
}

*/

func Close(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
	t := r.(*money.Trades)

	base := c.Args().Get(0)
	invs := c.Args().Get(1)
	id := int(util.Stf(c.Args().Get(2)))
	amnt := util.Stf(c.Args().Get(3))
	sell := util.Stf(c.Args().Get(4))
	ind := 0
	trade := money.Trade{}
	change := 1.0
	pam := 0.0
	for i := 0; i < len(t.Pairs); i++ {
		if t.Pairs[i].Base == base && t.Pairs[i].Invs == invs {
			ind = i
			trade.Amount = t.Pairs[i].Open[id].Amount
			trade.Cost = t.Pairs[i].Open[id].Cost
			trade.Buy = t.Pairs[i].Open[id].Buy
			change = amnt / t.Pairs[i].Open[id].Amount
			t.Pairs[i].Open[id].Cost -= change * t.Pairs[i].Open[id].Cost
			t.Pairs[i].Open[id].Amount -= amnt
			if t.Pairs[i].Open[id].Amount == 0.0 {
				t.Pairs[i].Open = t.Pairs[i].Open[:id+copy(t.Pairs[i].Open[id:], t.Pairs[i].Open[id+1:])]
			}
		}
	}
	pam = trade.Amount
	change = amnt / pam
	trade.Cost = change * trade.Cost
	trade.Amount = amnt
	trade.Sell = sell
	trade.Profit = (trade.Sell - trade.Buy) * amnt
	trade.Percent = (trade.Sell / trade.Buy) * 100
	t.Pairs[ind].Close = append(t.Pairs[ind].Close, trade)
	return true, nil
}
