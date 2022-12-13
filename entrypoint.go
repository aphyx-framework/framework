package main

import (
	"RyftFramework/framework/bootstrapper/logging"
	"RyftFramework/framework/configuration"
	container "RyftFramework/framework/di"
	"RyftFramework/framework/migration"
	"flag"
	"fmt"
	"github.com/TwiN/go-color"
	"os"
)

func main() {
	migratorFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
	fresh := migratorFlag.Bool("fresh", false, "Drop all table defined in RegisterModel")
	seed := migratorFlag.Bool("seed", false, "Seed the migration with data defined in the seeder")

	if len(os.Args) < 2 {
		// If no argument is passed, start the server
		BootstrapFramework()
	} else {
		switch os.Args[1] {
		case "migrate":
			err := migratorFlag.Parse(os.Args[2:])
			if err != nil {
				panic(err)
			}
			container.BuildForMigrator()
			migration.RunMigrator(*fresh, *seed)
		default:
			panic("Unknown command")
		}
	}

}

// BootstrapFramework ---
//
// This bootstrapper is responsible for initializing the framework
// and setting up the required dependencies.
func BootstrapFramework() {
	container.BuildAppFull()
	printAsciiArt()
	checkSecurityConfig()
	checkAuthenticationConfig()
	printEnabledFeature()
}

// printEnabledFeature --
//
// This function is responsible for printing the enabled features.
// If feature is enabled, it will show a green check mark
// If feature is disabled, it will show a red cross mark
func printEnabledFeature() {
	config := container.FrameworkDependency.Get(container.Config).(configuration.Configuration)

	println("Enabled features: ")
	if config.Database.Enabled {
		println(color.GreenBackground + color.Black + " [✓] Database " + color.Reset)
	} else {
		println(color.RedBackground + color.Black + " [X] Database " + color.Reset)
	}

	if config.Authentication.Enabled {
		println(color.GreenBackground + color.Black + " [✓] Authentication " + color.Reset)
	} else {
		println(color.RedBackground + color.Black + " [X] Authentication " + color.Reset)
	}

	if config.Caching.Enabled {
		println(color.GreenBackground + color.Black + " [✓] Caching " + color.Reset)
	} else {
		println(color.RedBackground + color.Black + " [X] Caching " + color.Reset)
	}
}

// checkSecurityConfig ---
//
// This function is responsible for checking the security config.
// It is a basic check to make sure that the secret key is set
// And you didn't enable debug mode in production
func checkSecurityConfig() {
	config := container.FrameworkDependency.Get(container.Config).(configuration.Configuration)
	logger := container.FrameworkDependency.Get(container.Logger).(logging.ApplicationLogger)

	if config.Security.Key == "" {
		logger.ErrorLogger.Fatalln("Security key is not set")
	}

	if config.Security.DebugMode == true && config.Security.Production == true {
		logger.WarningLogger.Print("Debug mode is enabled in production mode")
	}
}

// checkAuthenticationConfig ---
//
// This function is responsible for checking the authentication config.
// For authentication to work, a valid URL and key must be set
// And migration must be enabled
func checkAuthenticationConfig() {
	config := container.FrameworkDependency.Get(container.Config).(configuration.Configuration)
	logger := container.FrameworkDependency.Get(container.Logger).(logging.ApplicationLogger)

	if config.Authentication.Enabled == true && config.Authentication.AuthenticationUrl == "" {
		logger.ErrorLogger.Fatalln("Authentication URL is not set")
	}

	if config.Authentication.Enabled == true && config.Database.Enabled == false {
		logger.ErrorLogger.Fatalln("Database must be enabled to use authentication")
	}
}

// printAsciiArt ---
//
// This function is responsible for printing the ASCII art.
// It prints ascii art of the framework name
func printAsciiArt() {
	fmt.Printf(`
  _____        __ _   
 |  __ \      / _| |  
 | |__) |   _| |_| |_ 
 |  _  / | | |  _| __|
 | | \ \ |_| | | | |_ 
 |_|  \_\__, |_|  \__| %s - MVC Go Framework powered by Fiber %s
         __/ |        
        |___/         
`, color.CyanBackground+color.Black, color.Reset)
}
