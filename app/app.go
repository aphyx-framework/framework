package app

import (
	"github.com/rama-adi/RyFT-Framework/framework/caching"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"github.com/rama-adi/RyFT-Framework/framework/utils"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Expose fx DI data to the rest of the application
// This value is essential for your app to work properly.
// The value here is populated by the framework, but to avoid cyclic dependencies
// It must be pointed again in the app package.
var (
	Config     configuration.Configuration
	DB         *gorm.DB
	Logger     logging.ApplicationLogger
	Utilities  utils.BuiltinUtilities
	CacheTable map[string]caching.CacheTable
)

// Dependencies -  Get access to the framework dependency injection container
// You're free to add your own dependencies here
var Dependencies = fx.Options()
