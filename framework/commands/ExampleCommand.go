package commands

import "github.com/aphyx-framework/framework/framework/cli"

func ExampleCommand(registry cli.Registry) {
	cmd := cli.Command{
		Command:     "example",
		Title:       "Example Command",
		Description: "A command that does nothing but print passed arguments",
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
			"example example:hi":              "Specify the argument example as hi",
			"example example:hi example2:bye": "Specify the argument example as hi and example2 as bye",
		},
		Handler: func(c cli.CommandArgumentValue) {
			println("This is an example command")
			println("This command does nothing, but it prints some arguments!")
			println("The arguments are:")
			println("example: " + c.GetArgument("example", ""))
			println("example2: " + c.GetArgument("example2", "unspecified"))
		},
	}
	registry.AddCommand(cmd)
}
