package bootstrapper

import (
	"RyftFramework/framework/configuration"
	"RyftFramework/framework/logging"
)

func AllBootstrapper(logger logging.ApplicationLogger, configuration configuration.Configuration) {
	printAsciiArt()
	checkSecurityConfig(logger, configuration)
	checkAuthenticationConfig(configuration, logger)
	PrintEnabledFeatures(configuration)
}
