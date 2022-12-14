package utils

import (
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
)

type Util struct {
	Config configuration.Configuration
}

func NewUtil(config configuration.Configuration) Util {
	return Util{
		Config: config,
	}
}
