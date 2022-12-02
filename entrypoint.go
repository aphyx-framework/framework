package main

import (
	"RyftFramework/configuration"
	"RyftFramework/database"
	"RyftFramework/migration"
	"RyftFramework/routing"
	"RyftFramework/utils"
	"flag"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	migratorFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
	fresh := migratorFlag.Bool("fresh", false, "Drop all table defined in RegisterModel")
	seed := migratorFlag.Bool("seed", false, "Seed the migration with data defined in the seeder")

	if len(os.Args) < 2 {
		// If no argument is passed, start the server
		BootstrapFramework()
	}

	switch os.Args[1] {
	case "migrate":
		err := migratorFlag.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
		migration.RunMigrator(*fresh, *seed)
	default:
		panic("Unknown command")
	}

}

// BootstrapFramework ---
//
// This bootstrapper is responsible for initializing the framework
// and setting up the required dependencies.
func BootstrapFramework() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		AppName:               configuration.ApplicationConfig.Application.Name,
	})
	printAsciiArt()
	configuration.LoadConfigFile()
	routing.LoadRouter(app)
	utils.LoadLogger()
	checkSecurityConfig()
	checkAuthenticationConfig()
	database.ConnectDatabase()
	printEnabledFeature()

	utils.InfoLogger.Print("Application started on port " + configuration.ApplicationConfig.Application.Port)
	err := app.Listen(":" + configuration.ApplicationConfig.Application.Port)

	if err != nil {
		log.Fatalln(err)
	}

}

// printEnabledFeature --
//
// This function is responsible for printing the enabled features.
// If feature is enabled, it will show a green check mark
// If feature is disabled, it will show a red cross mark
func printEnabledFeature() {
	println("Enabled features: ")
	if configuration.ApplicationConfig.Database.Enabled {
		println(color.GreenBackground + color.Black + " [✓] Database " + color.Reset)
	} else {
		println(color.RedBackground + color.Black + " [X] Database " + color.Reset)
	}

	if configuration.ApplicationConfig.Authentication.Enabled {
		println(color.GreenBackground + color.Black + " [✓] Authentication " + color.Reset)
	} else {
		println(color.RedBackground + color.Black + " [X] Authentication " + color.Reset)
	}

	if configuration.ApplicationConfig.Caching.Enabled {
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
	if configuration.ApplicationConfig.Security.Key == "" {
		utils.ErrorLogger.Fatalln("Security key is not set")
	}

	if configuration.ApplicationConfig.Security.DebugMode == true && configuration.ApplicationConfig.Security.Production == true {
		utils.WarningLogger.Print("Debug mode is enabled in production mode")
	}
}

// checkAuthenticationConfig ---
//
// This function is responsible for checking the authentication config.
// For authentication to work, a valid URL and key must be set
// And migration must be enabled
func checkAuthenticationConfig() {
	if configuration.ApplicationConfig.Authentication.Enabled == true && configuration.ApplicationConfig.Authentication.AuthenticationUrl == "" {
		utils.ErrorLogger.Fatalln("Authentication URL is not set")
	}

	if configuration.ApplicationConfig.Authentication.Enabled == true && configuration.ApplicationConfig.Database.Enabled == false {
		utils.ErrorLogger.Fatalln("Database must be enabled to use authentication")
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
