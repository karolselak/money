package money

// Wealth an object containing assets
type Wealth struct {
	Wealth []Assets `json:"wealth"`
	Worth  float64  `json:"worth"`
}

// Assets contain an arrah of type Asset and a catgorie
type Assets struct {
	Type   string  `json:"type"`
	Assets []Asset `json:"assets"`
}

// Asset stuct contain 1 asset type
type Asset struct {
	Name    string  `json:"name"`
	Symbol  string  `json:"symbol"`
	Holding float64 `json:"holding"`
	Worth   float64 `json:"worth"`
}

// Update sets struct fields to the most recent Prices
func (w *Wealth) Update() {
	var sum float64
	var hold float64
	var wor float64
	for i := 0; i < len(Prices); i++ {
		hold = w.Wealth[1].Assets[i].Holding
		wor = hold * Prices[i]
		w.Wealth[1].Assets[i].Worth = wor
		sum += wor
	}
	hold = w.Wealth[0].Assets[0].Holding
	sum += hold * 0.75
	w.Worth = sum
}
