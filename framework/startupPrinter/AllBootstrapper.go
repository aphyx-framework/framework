package startupPrinter

import (
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
)

func AllBootstrapper(logger logging.ApplicationLogger, configuration configuration.Configuration) {
	printAsciiArt()
	checkSecurityConfig(logger, configuration)
	checkAuthenticationConfig(configuration, logger)
	PrintEnabledFeatures(configuration)
}
