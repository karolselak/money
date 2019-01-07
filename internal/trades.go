package money

// Trades carries an array for open trades, and an array for closed trades
type Trades struct {
	Open   []Trade  `json:"open"`
	Closed []Tradec `json:"closed"`
}

// Trade carries trade info
type Trade struct {
	Base    string  `json:"base"`
	Invs    string  `json:"invs"`
	Claim   bool    `json:"claim"`
	Cost    float64 `json:"cost"`
	Amount  float64 `json:"amount"`
	Buy     float64 `json:"buy"`
	Current float64 `json:"current"`
}

type Tradec struct {
	Base    string  `json:"base"`
	Invs    string  `json:"invs"`
	Claim   bool    `json:"claim"`
	Cost    float64 `json:"cost"`
	Profit  float64 `json:"profit"`
	percent float64 `json:"percent"`
}

func (t *Trades) Update() {
	var buy bool
	var currcp float64
	var currip float64
	var curr float64
	for i := 0; i < len(t.Open); i++ {
		buy = t.Open[i].Claim
		if buy == true {
			currcp = Prices[t.Open[i].Base]
			currip = Prices[t.Open[i].Invs]
		} else {
			currcp = Prices[t.Open[i].Invs]
			currip = Prices[t.Open[i].Base]
		}
		curr = currip / currcp
		t.Open[i].Current = curr
	}
}
