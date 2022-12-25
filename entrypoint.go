package main

import (
	"github.com/aphyx-framework/framework/framework"
	"github.com/aphyx-framework/framework/framework/cli"
	"os"
)

func main() {
	// If you want to debug the DI container, set this to true
	const EnableNopLogger = false

	// Bootstraps the web framework
	if len(os.Args) < 2 {
		framework.RunWebApplication(EnableNopLogger) // If no argument is passed, start the server
	} else {
		cli.RunCliApplication(EnableNopLogger) // Invoke the CLI application
	}
}
