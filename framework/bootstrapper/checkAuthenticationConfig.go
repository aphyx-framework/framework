package bootstrapper

import (
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
)

func checkAuthenticationConfig(config configuration.Configuration, logger logging.ApplicationLogger) {

	if config.Authentication.Enabled == true && config.Authentication.AuthenticationUrl == "" {
		logger.ErrorLogger.Fatalln("Authentication URL is not set")
	}

	if config.Authentication.Enabled == true && config.Database.Enabled == false {
		logger.ErrorLogger.Fatalln("Database must be enabled to use authentication")
	}
}
