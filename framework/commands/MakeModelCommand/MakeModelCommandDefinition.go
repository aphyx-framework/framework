package MakeModelCommand

import (
	"github.com/aphyx-framework/framework/framework/cli"
	"github.com/aphyx-framework/framework/framework/logging"
	"github.com/aphyx-framework/framework/framework/utils"
)

func Definition(registry cli.Registry, logger logging.ApplicationLogger, utilities utils.BuiltinUtilities) {
	cmd := cli.Command{
		Command:     "make:model",
		Title:       "Make a model",
		Description: "Make a model inside of your model folder",
		Args: []cli.CommandArgument{
			{
				Name:        "name",
				Description: "The name of the model",
				Required:    true,
			},
		},
		ExampleUsage: map[string]string{
			"make:model User": "Creates a model called User",
		},
		Handler: func(c cli.CommandArgumentValue) {
			modelName := c.GetArgument("name", "")
			makeModel(modelName, logger, utilities.Strings, utilities.Modfile)
		},
	}
	registry.AddCommand(cmd)
}
