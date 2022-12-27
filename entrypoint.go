package main

import (
	"github.com/aphyx-framework/framework/framework"
	"os"
)

func main() {
	// If you want to debug the DI container, set this to true
	const EnableFxLogger = true

	if len(os.Args) < 2 {
		framework.BoostrapKernel(EnableFxLogger, false) // If no argument is passed, start the server
	} else {
		framework.BoostrapKernel(EnableFxLogger, true) // Invoke the CLI application
	}
}
