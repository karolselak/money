package main

type Coins struct {
	Coins []Coin `json:"coins"`
}

type Coin struct {
	Name    string `json:"name"`
	Symbol  string `json:"sym"`
	Holding string `json:"holdings"`
}
