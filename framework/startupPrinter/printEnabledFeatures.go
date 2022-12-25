package startupPrinter

import (
	"github.com/TwiN/go-color"
	"github.com/aphyx-framework/framework/framework/configuration"
)

func printEnabledFeatures(config configuration.Configuration) {
	println("Enabled features: ")
	if config.Database.Enabled {
		println(color.GreenBackground + color.Black + " [✓] " + color.Reset + color.Green + " Database" + color.Reset)
	} else {
		println(color.RedBackground + color.Black + " [X] " + color.Reset + color.Red + " Database" + color.Reset)
	}

	if config.Authentication.Enabled {
		println(color.GreenBackground + color.Black + " [✓] " + color.Reset + color.Green + " Authentication" + color.Reset)
	} else {
		println(color.RedBackground + color.Black + " [X] " + color.Reset + color.Red + " Authentication" + color.Reset)
	}

	if config.Caching.Enabled {
		println(color.GreenBackground + color.Black + " [✓] " + color.Reset + color.Green + " Caching" + color.Reset)
	} else {
		println(color.RedBackground + color.Black + " [X] " + color.Reset + color.Red + " Caching" + color.Reset)
	}

	println("")
}
