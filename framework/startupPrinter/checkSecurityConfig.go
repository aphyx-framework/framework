package startupPrinter

import (
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
)

func checkSecurityConfig(logger logging.ApplicationLogger, configuration configuration.Configuration) {
	if configuration.Security.Key == "" {
		logger.ErrorLogger.Fatalln("Security key is not set")
	}

	if configuration.Security.DebugMode && configuration.Security.Production {
		logger.ErrorLogger.Println("Debug mode is enabled in production")
	}
}
