package GenerateKeyCommand

import (
	"github.com/aphyx-framework/framework/framework/cli"
	"github.com/aphyx-framework/framework/framework/logging"
	"github.com/aphyx-framework/framework/framework/utils"
)

func Definition(registry cli.Registry, logger logging.ApplicationLogger, utilities utils.BuiltinUtilities) {
	cmd := cli.Command{
		Command:     "generate-key",
		Title:       "Generate a key for the application",
		Description: "This command will create a 32 character key for the application",
		Args:        []cli.CommandArgument{},
		ExampleUsage: map[string]string{
			"generate-key": "Generates a key for the application",
		},
		Handler: func(c cli.CommandArgumentValue) {
			logger.InfoLogger.Println("Generating key...")
			logger.InfoLogger.Println("Generated key: ", utilities.Strings.Random(32))
		},
	}
	registry.AddCommand(cmd)
}
