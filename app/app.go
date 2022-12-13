package app

import (
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"gorm.io/gorm"
)

// Expose fx DI data to the rest of the application
var (
	Config configuration.Configuration
	DB     *gorm.DB
	Logger logging.ApplicationLogger
)
