package GenerateKeyCommand

import (
	"encoding/base64"
	"github.com/aphyx-framework/framework/framework/cli"
	"github.com/aphyx-framework/framework/framework/logging"
	"math/rand"
)

func Definition(registry cli.Registry, logger logging.ApplicationLogger) {
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
			b := make([]byte, 32)
			_, err := rand.Read(b)
			if err != nil {
				logger.ErrorLogger.Fatalln(err)
				return
			}

			// Encode the bytes to base64
			key := base64.StdEncoding.EncodeToString(b)
			logger.InfoLogger.Println("Generated key: ", key)
		},
	}
	registry.AddCommand(cmd)
}
