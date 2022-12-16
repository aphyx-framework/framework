package app

import (
	"github.com/rama-adi/RyFT-Framework/app/utils"
	"github.com/rama-adi/RyFT-Framework/framework/caching"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Expose fx DI data to the rest of the application
var (
	Config     configuration.Configuration
	DB         *gorm.DB
	Logger     logging.ApplicationLogger
	Utilities  utils.Util
	CacheTable map[string]caching.CacheTable
)

var Dependencies = fx.Options(
	fx.Provide(utils.NewUtil),
	fx.Populate(&Utilities),
	fx.Populate(&CacheTable),
)
