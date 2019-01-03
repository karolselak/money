package main

// Wealth an object containing assets
type Wealth struct {
	Fiat   []Asset `json:fiat`
	Crypto []Asset `json:"crypto"`
	Metals []Asset `json:"metals"`
}

// Asset stuct contain 1 asset type
type Asset struct {
	Name    string  `json:"name"`
	Symbol  string  `json:"symbol"`
	Holding float64 `json:"holding"`
}
