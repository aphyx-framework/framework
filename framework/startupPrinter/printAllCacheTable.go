package startupPrinter

import (
	"github.com/TwiN/go-color"
	"github.com/rama-adi/RyFT-Framework/framework/caching"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
)

func printAllCacheTable(table map[string]caching.CacheTable, config configuration.Configuration) {
	if config.Caching.Enabled {
		println("Cache table loaded:")
		for key, _ := range table {
			println(color.GreenBackground + color.Black + " [📦] " + color.Reset + " " + key)
		}
	}
}
