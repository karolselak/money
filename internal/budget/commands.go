package budget

import (
	"fmt"

	money "github.com/mohfunk/money/internal"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func List(r money.Resource, log *logrus.Logger, c *cli.Context) (bool, error) {
	b := r.(*money.Budget)
	var data [][]string
	var cost float64
	for i := 0; i < len(b.Expenses); i++ {
		data = append(data, []string{})
		data[i] = append(data[i], b.Expenses[i].Name)
		cost = b.Expenses[i].CostMonth
		data[i] = append(data[i], fmt.Sprintf("%5.2f", cost))
	}

	b.Update()
	prnt(data, "Budget")
	prntTotal(fmt.Sprintf("%8.2f", b.Total))
	return false, nil

}
