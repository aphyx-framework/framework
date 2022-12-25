package startupPrinter

import (
	"github.com/aphyx-framework/framework/framework/caching"
	"github.com/aphyx-framework/framework/framework/configuration"
	"github.com/aphyx-framework/framework/framework/logging"
)

func PrintStartupInfo(
	logger logging.ApplicationLogger,
	configuration configuration.Configuration,
	cacheTable map[string]caching.CacheTable,
) {
	printAsciiArt()
	checkSecurityConfig(logger, configuration)
	checkAuthenticationConfig(configuration, logger)
	printEnabledFeatures(configuration)
	printAllCacheTable(cacheTable, configuration)
}
