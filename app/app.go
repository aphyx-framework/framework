package app

import (
	"RyftFramework/framework/configuration"
	"RyftFramework/framework/logging"
	"gorm.io/gorm"
)

// Expose fx DI data to the rest of the application
var (
	Config configuration.Configuration
	DB     *gorm.DB
	Logger logging.ApplicationLogger
)
