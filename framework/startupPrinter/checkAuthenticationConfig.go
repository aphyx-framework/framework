package startupPrinter

import (
	"github.com/aphyx-framework/framework/framework/configuration"
	"github.com/aphyx-framework/framework/framework/logging"
)

func checkAuthenticationConfig(config configuration.Configuration, logger logging.ApplicationLogger) {

	if config.Authentication.Enabled == true && config.Authentication.AuthenticationUrl == "" {
		logger.ErrorLogger.Fatalln("Authentication URL is not set")
	}

	if config.Authentication.Enabled == true && config.Database.Enabled == false {
		logger.ErrorLogger.Fatalln("Database must be enabled to use authentication")
	}
}
