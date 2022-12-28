package cli

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/alexeyco/simpletable"
	"github.com/gosuri/uitable"
	"os"
)

var StyleThinUnicode = &simpletable.Style{
	Border: &simpletable.BorderStyle{
		TopLeft:            "┌",
		Top:                "─",
		TopRight:           "┐",
		Right:              "│",
		BottomRight:        "┘",
		Bottom:             "─",
		BottomLeft:         "└",
		Left:               "│",
		TopIntersection:    "┬",
		BottomIntersection: "┴",
	},
	Divider: &simpletable.DividerStyle{
		Left:         "├",
		Center:       "─",
		Right:        "┤",
		Intersection: "┼",
	},
	Cell: "│",
}

func RunCommand(registry Registry) {

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

	joinedArgs := "Arguments are not required"

	argsTable := simpletable.New()
	argsTable.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "Usage"},
			{Align: simpletable.AlignLeft, Text: "Description"},
			{Align: simpletable.AlignLeft, Text: "Required?"},
		},
	}

	if len(command.Args) > 0 {
		for _, arg := range command.Args {

			required := "X"

			if arg.Required {
				required = "✓"
			}

			row := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: arg.Name},
				{Align: simpletable.AlignLeft, Text: arg.Description},
				{Align: simpletable.AlignLeft, Text: required},
			}
			argsTable.Body.Cells = append(argsTable.Body.Cells, row)
		}

		joinedArgs = argsTable.String()
	}

	// Print the basic information
	table.AddRow(color.GreenBackground+color.Black+command.Command, color.Reset+"   "+command.Title)
	table.AddRow("Description:", "   "+command.Description)

	// If there are no arguments, add a row saying that
	if len(command.Args) < 1 {
		table.AddRow("Arguments:", "   "+joinedArgs)
	}

	// If there are no usage examples, add a row saying that
	if len(command.ExampleUsage) < 1 {
		table.AddRow("Usages:", "   "+"Command does not provide any example usages")
	}

	// Print the command table
	fmt.Println(table)

	// If the command has arguments, print the arguments table
	if len(command.Args) > 0 {
		println("Arguments:")
		fmt.Println(joinedArgs)
	}

	if len(command.ExampleUsage) > 0 || len(command.Args) > 0 {
		println(color.GreenBackground + color.Black + "      This command provides example usages or arguments.      " + color.Reset)
	}

	// If the command has example usages, print the example usages
	if len(command.ExampleUsage) > 0 {
		usageTable := simpletable.New()
		usageTable.SetStyle(StyleThinUnicode)
		usageTable.Header = &simpletable.Header{
			Cells: []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: "Usage"},
				{Align: simpletable.AlignLeft, Text: "Description"},
			},
		}

		for usage, description := range command.ExampleUsage {
			row := []*simpletable.Cell{
				{Align: simpletable.AlignLeft, Text: usage},
				{Align: simpletable.AlignLeft, Text: description},
			}
			usageTable.Body.Cells = append(usageTable.Body.Cells, row)
		}

		println("Usages:")
		fmt.Println(usageTable.String())
	}
}
