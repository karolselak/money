package main

import (
	"io"
	"os"

	c "github.com/mohfunk/money/pkg/color"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func setHelpText() {
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		c.B[3].Println("Money")
		println()
		c.B[2].Println("COMMANDS")

		println()
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Command", "Alias", "Usage", "function"})
		table.Append([]string{"ls", "l", "money ls", "lists all assets"})
		table.Append([]string{"ad", "a", "money ad NAME SYMBOL", "adds a new asset type"})
		table.Append([]string{"md", "m", "money md SYMBOL [+/-] AMOUNT", "modifies holding"})
		table.SetBorder(true)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeaderColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor})

		table.SetColumnColor(
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
			tablewriter.Colors{tablewriter.Bold, tablewriter.FgCyanColor})

		table.Render() // Send output
		println()
	}
}

func main() {
	c.SetColors()
	setHelpText()
	app := &Application{}
	app.init()
	app.run()
}
