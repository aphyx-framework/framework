package cli

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/gosuri/uitable"
	"os"
)

func RunCommandFromConsole(registry Registry) {

	command := os.Args[1]
	args := os.Args[2:]
	helpMode := false

	// Get the command from the registry
	commandResult, err := registry.GetCommand(command)

	// If the command is not found, panic and exit
	if err != nil {
		panic(err)
	}

	// Checks if the arguments contains the help flag
	for _, arg := range args {
		if arg == "--help" {
			helpMode = true
			printHelp(commandResult)
		}
	}

	// Check if the arguments satisfies the requirements
	if len(args) < len(commandResult.Args) {
		println(color.RedBackground + color.White + " Error: " + color.Reset + "  Not enough arguments provided")
		println("Expected: " + fmt.Sprintf("%d", len(commandResult.Args)) + " arguments")
		println("Received: " + fmt.Sprintf("%d", len(args)) + " arguments")
		os.Exit(1)
	}

	if helpMode == false {
		commandResult.Handler(args...)
	}

	os.Exit(0)
}

func printHelp(command Command) {
	table := uitable.New()
	table.MaxColWidth = 80
	table.Wrap = true // wrap columns

	joinedArgs := "[No arguments needed]"
	argsTable := uitable.New()
	argsTable.MaxColWidth = 80
	argsTable.Wrap = true // wrap columns

	if len(command.Args) > 0 {

		for _, arg := range command.Args {
			argsTable.AddRow(arg.Name, "► "+arg.Description+" (Required: "+fmt.Sprintf("%t", arg.Required)+")")
		}

		// Join the slice with commas
		joinedArgs = argsTable.String()
	}

	table.AddRow(color.GreenBackground+color.Black+command.Command, color.Reset+"   "+command.Title)
	table.AddRow("Description:", "   "+command.Description)
	table.AddRow("Arguments:", "   "+joinedArgs)

	if len(command.ExampleUsage) < 1 {
		table.AddRow("Usages:", "   "+"[No usage example available]")
	} else {
		usageTable := uitable.New()
		usageTable.MaxColWidth = 80
		usageTable.Wrap = true // wrap columns

		for usage, description := range command.ExampleUsage {
			usageTable.AddRow("   - "+usage, "► "+description)
		}

		table.AddRow("Usages:", usageTable.String())
	}
	table.AddRow("") // blank

	fmt.Println(table)
}
