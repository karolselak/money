package money

// Trades carries an array for open trades, and an array for closed trades
type Trades struct {
	Open   []Trade `json:"open"`
	Closed []Trade `json:"closed"`
}

// Trade carries trade info
type Trade struct {
	Pair    string  `json:"pair"`
	Cost    float64 `json:"cost"`
	Amount  float64 `json:"amount"`
	Buy     float64 `json:"buy"`
	Sell    float64 `json:"sell"`
	Profit  float64 `json:"profit"`
	Current float64 `json:"current"`
}

func (t *Trades) Update() {
	t.Open[0].Cost = 4
}
