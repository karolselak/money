package data

// Assets an array of Asset
type Wealth struct {
	Fiat   []Asset `json:fiat`
	Crypto []Asset `json:"crypto"`
	Metals []Asset `json:"metals"`
}

// Asset stuct
type Asset struct {
	Name    string  `json:"name"`
	Symbol  string  `json:"symbol"`
	Holding float64 `json:"holding"`
}

// createAsset returns an asset
func CreateAsset(n string, s string) *Asset {
	na := &Asset{
		Name:    n,
		Symbol:  s,
		Holding: 0.0,
	}
	return na
}

//func EmpAssets(a *Assets) bool {
//
//}
