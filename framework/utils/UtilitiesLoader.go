package utils

import (
	"github.com/aphyx-framework/framework/framework/configuration"
	"github.com/aphyx-framework/framework/framework/logging"
)

type BuiltinUtilities struct {
	Strings Strings
	Crypto  Crypto
	Modfile Modfile
}

func InitializeFrameworkUtils(config configuration.Configuration, logger logging.ApplicationLogger) BuiltinUtilities {
	return BuiltinUtilities{
		Strings: Strings{},
		Crypto: Crypto{
			config: config,
		},
		Modfile: Modfile{
			logger: logger,
		},
	}
}
