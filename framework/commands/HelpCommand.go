package commands

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/aphyx-framework/framework/framework/cli"
	"github.com/gosuri/uitable"
)

func HelpCommand(registry cli.Registry) {
	cmd := cli.Command{
		Command:     "help",
		Title:       "CLI Help",
		Description: "Displays all of the available commands",
		Args:        []cli.CommandArgument{},
		ExampleUsage: map[string]string{
			"help": "Displays all of the available commands",
		},
		Handler: func(arg ...string) {
			helpCommandHandler(registry)
		},
	}
	registry.AddCommand(cmd)
}

func helpCommandHandler(registry cli.Registry) {
	println(color.CyanBackground + color.Black + " Apyhx Framework Commands: " + color.Reset)
	println("Info: add --help to the command for more information")
	println()

	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true // wrap columns

	for _, command := range registry.GetCommands() {

		table.AddRow(color.GreenBackground+color.Black+command.Command, color.Reset+"   "+command.Title)
		table.AddRow("Description:", "   "+command.Description)
		table.AddRow("") // blank
	}

	fmt.Println(table)
}
