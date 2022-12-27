package commands

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/aphyx-framework/framework/framework/cli"
	"github.com/gosuri/uitable"
	"strings"
)

func HelpCommand(registry cli.Registry) {
	cmd := cli.Command{
		Command:     "help",
		Title:       "CLI Help",
		Description: "Displays all of the available commands",
		Args:        map[string]string{},
		Handler: func(arg ...string) {
			helpCommandHandler(registry)
		},
	}
	registry.AddCommand(cmd)
}

func helpCommandHandler(registry cli.Registry) {
	println(color.CyanBackground + color.Black + " Apyhx Framework Commands: " + color.Reset)
	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true // wrap columns

	for _, command := range registry.GetCommands() {

		joinedArgs := "[No arguments needed]"

		if len(command.Args) > 0 {
			args := make([]string, 0, len(command.Args))
			for k := range command.Args {
				args = append(args, k)
			}

			// Join the slice with commas
			joinedArgs = strings.Join(args, ", ")
		}

		table.AddRow(color.GreenBackground+color.Black+command.Command, color.Reset+"   "+command.Title)
		table.AddRow("Description:", "   "+command.Description)
		table.AddRow("Arguments:", "   "+joinedArgs)
		table.AddRow("") // blank
	}

	fmt.Println(table)
}
