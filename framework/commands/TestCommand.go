package commands

import "github.com/aphyx-framework/framework/framework/cli"

func TestCommand(registry cli.Registry) {
	cmd := cli.Command{
		Command:     "test",
		Title:       "Command for CLI testing",
		Description: "Test CLI",
		Args:        map[string]string{},
		Handler:     aboutCommandHandler,
	}
	registry.AddCommand(cmd)
}

func aboutCommandHandler(...string) {
	println("Testing")
}
