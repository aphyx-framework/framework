package HelpCommand

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/aphyx-framework/framework/framework/cli"
	"github.com/gosuri/uitable"
)

func handler(registry cli.Registry) {
	println(color.CyanBackground + color.Black + " Apyhx Framework Commands: " + color.Reset)
	println("Info: add --help to the command for more information")
	println()

	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true // wrap columns

	for _, command := range registry.GetCommands() {

		table.AddRow("'"+command.Command+"'", command.Title)
		table.AddRow("Description:", command.Description)
		table.AddRow("") // blank
	}

	fmt.Println(table)
}
