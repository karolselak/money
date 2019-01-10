package wealth

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func prnt(data [][]string, cap string) {
	println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Asset", "Holding", "Worth"})

	for _, v := range data {
		table.Append(v)
	}
	table.SetBorder(true)
	table.SetCaption(true, cap)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor})

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiCyanColor})

	table.Render() // Send output
	println()
}

func prntTotal(sum string) {

	println()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Total"})
	table.Append([]string{sum})
	table.SetBorder(true)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor})
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor})

	table.Render() // Send output
	println()
}
