package main

// Wealth an object containing assets
type Wealth struct {
	Wealth []Assets
}
type Assets struct {
	Type   string
	Assets []Asset
}

// Asset stuct contain 1 asset type
type Asset struct {
	Name    string
	Symbol  string
	Holding float64
}

func NewWealth() *Wealth {
	return &Wealth{}
}
