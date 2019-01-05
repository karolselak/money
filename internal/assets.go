package money

// Wealth an object containing assets
type Wealth struct {
	Wealth []Assets `json:"wealth"`
	Worth  float64  `json:"worth"`
}
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

func NewWealth() *Wealth {
	return &Wealth{}
}
