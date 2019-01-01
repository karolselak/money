package data

// Assets an array of Asset
type Assets struct {
	Assets []Asset `json:"Assets"`
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

func InitAssets() *Assets {
	f := &Assets{
		Assets: []Asset{},
	}
	return f
}
