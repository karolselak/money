package money

// Trades carries an array for open trades, and an array for closed trades
type Trades struct {
	Pairs []Pair `json:"pairs"`
}

// Trade carries trade info
type Pair struct {
	Base  string  `json:"base"`
	Invs  string  `json:"invs"`
	Open  []Trade `json:"open"`
	Close []Trade `json:"close"`
}

// Tradec closed trades
type Trade struct {
	Status  bool    `json:"status"`
	Claim   bool    `json:"claim"`
	Cost    float64 `json:"cost"`
	Amount  float64 `json:"amount"`
	Buy     float64 `json:"buy"`
	Sell    float64 `json:"sell"`
	Percent float64 `json:"percent"`
	Current float64 `json:"current"`
	Profit  float64 `json:"profit"`
}

// Update runs after commands that modify the struct fields
func (t *Trades) Update() {
	var csym string
	var isym string
	var ccurr float64
	var iamnt float64
	var icurr float64
	var curr float64
	var cost float64
	for i := 0; i < len(t.Pairs); i++ {
		csym = t.Pairs[i].Base
		isym = t.Pairs[i].Invs
		for j := 0; j < len(t.Pairs[i].Open); j++ {
			iamnt = t.Pairs[i].Open[j].Amount
			cost = t.Pairs[i].Open[j].Cost
			ccurr = Prices[csym]
			icurr = Prices[isym]
			curr = ((icurr / ccurr) * iamnt)
			t.Pairs[i].Open[j].Current = curr
			t.Pairs[i].Open[j].Percent = (cost / curr) * 100
			t.Pairs[i].Open[j].Profit = (curr - cost) * ccurr
		}
	}
}
