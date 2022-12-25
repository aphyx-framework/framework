package startupPrinter

import (
	"github.com/TwiN/go-color"
	"github.com/aphyx-framework/framework/framework/caching"
	"github.com/aphyx-framework/framework/framework/configuration"
)

func printAllCacheTable(table map[string]caching.CacheTable, config configuration.Configuration) {
	if config.Caching.Enabled {
		println("Cache table loaded:")
		for key := range table {
			println(color.GreenBackground + color.Black + " [📦] " + color.Reset + " " + key)
		}
	}
}
