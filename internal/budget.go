package money

// Budget carries an array for open trades, and an array for closed trades
type Budget struct {
	Expenses []Expense `json:"expenses"`
	Total    float64   `json:"total"`
}

type Expense struct {
	Name      string  `json:"name"`
	CostMonth float64 `json:"cost_month"`
}

// Update runs after commands that modify the struct fields
func (b *Budget) Update() {
	var cost float64
	for i := 0; i < len(b.Expenses); i++ {
		cost += b.Expenses[i].CostMonth
	}
	b.Total = cost

}
