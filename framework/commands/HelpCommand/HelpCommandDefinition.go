package HelpCommand

import (
	"github.com/aphyx-framework/framework/framework/cli"
)

func Definition(registry cli.Registry) {
	cmd := cli.Command{
		Command:     "help",
		Title:       "CLI Help",
		Description: "Displays all of the available commands",
		Args:        []cli.CommandArgument{},
		ExampleUsage: map[string]string{
			"help": "Displays all of the available commands",
		},
		Handler: func(c cli.CommandArgumentValue) {
			handler(registry)
		},
	}
	registry.AddCommand(cmd)
}
