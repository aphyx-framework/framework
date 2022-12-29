package MakeControllerCommand

import (
	"github.com/aphyx-framework/framework/framework/cli"
	"github.com/aphyx-framework/framework/framework/logging"
	"github.com/aphyx-framework/framework/framework/utils"
)

func Definition(registry cli.Registry, logger logging.ApplicationLogger, utilities utils.BuiltinUtilities) {
	cmd := cli.Command{
		Command: "make:controller",
		Title:   "Make Controller and its handler",
		Description: "This command will create a controller and its handler.\n" +
			"If the controller folder exists, it will be skipped and only the handler will be created",
		Args: []cli.CommandArgument{
			{
				Name:        "name",
				Description: "The name of the controller",
				Required:    true,
			},
			{
				Name:        "handler",
				Description: "The name of the handler",
				Required:    true,
			},
		},
		ExampleUsage: map[string]string{
			"make:controller name:UserController handler:UserHandler": "Creates a controller named UserController and a handler named UserHandler",
		},
		Handler: func(c cli.CommandArgumentValue) {
			controllerName := c.GetArgument("name", "")
			handlerName := c.GetArgument("handler", "")
			makeController(controllerName, handlerName, logger, utilities.Strings, utilities.Modfile)
		},
	}
	registry.AddCommand(cmd)
}
