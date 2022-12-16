package startupPrinter

import (
	"github.com/rama-adi/RyFT-Framework/framework/caching"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
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
