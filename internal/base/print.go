package base

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func prnt(data [][]string, sum float64) {
	println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Asset", "Holding", "Worth"})

	for _, v := range data {
		table.Append(v)
	}
	d := fmt.Sprintf("%f", sum)
	table.SetFooter([]string{"", "Total", d})
	table.SetBorder(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlueColor})

	table.SetFooterColor(
		tablewriter.Colors{},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor})
	table.Render() // Send output
	println()
}
