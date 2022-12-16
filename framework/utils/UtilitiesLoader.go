package utils

import (
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
)

type BuiltinUtilities struct {
	Strings Strings
	Crypto  Crypto
}

func InitializeFrameworkUtils(config configuration.Configuration) BuiltinUtilities {
	return BuiltinUtilities{
		Strings: Strings{},
		Crypto:  Crypto{config: config},
	}
}
