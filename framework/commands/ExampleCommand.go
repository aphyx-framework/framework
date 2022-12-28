package commands

import "github.com/aphyx-framework/framework/framework/cli"

func ExampleCommand(registry cli.Registry) {
	cmd := cli.Command{
		Command:     "example",
		Title:       "Example Command",
		Description: "This is an example command",
		Args: []cli.CommandArgument{
			{
				Name:        "example",
				Description: "This is an example argument",
				Required:    true,
			},
			{
				Name:        "example2",
				Description: "This is an example argument that is not required",
				Required:    false,
			},
		},
		ExampleUsage: map[string]string{
			"example": "This is an example command",
		},
		Handler: func(arg ...string) {
			println("This is an example command")
			println("This command does nothing")
		},
	}
	registry.AddCommand(cmd)
}
