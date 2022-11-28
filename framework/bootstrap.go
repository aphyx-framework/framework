package framework

import (
	"RyftFramework/utils"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/gofiber/fiber/v2"
	"log"
)

// BootstrapFramework ---
//
// This bootstrapper is responsible for initializing the framework
// and setting up the required dependencies.
func BootstrapFramework() {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		AppName:               ApplicationConfig.Application.Name,
	})
	printAsciiArt()
	loadConfigFile()
	loadRouter(app)
	utils.LoadLogger()
	checkSecurityConfig()
	checkAuthenticationConfig()
	connectDatabase()
	printEnabledFeature()

	utils.InfoLogger.Print("Application started on port " + ApplicationConfig.Application.Port)
	err := app.Listen(":" + ApplicationConfig.Application.Port)

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
	if ApplicationConfig.Database.Enabled {
		println(color.GreenBackground + color.Black + " [✓] Database " + color.Reset)
	} else {
		println(color.RedBackground + color.Black + " [X] Database " + color.Reset)
	}

	if ApplicationConfig.Authentication.Enabled {
		println(color.GreenBackground + color.Black + " [✓] Authentication " + color.Reset)
	} else {
		println(color.RedBackground + color.Black + " [X] Authentication " + color.Reset)
	}

	if ApplicationConfig.Caching.Enabled {
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
	if ApplicationConfig.Security.Key == "" {
		utils.ErrorLogger.Fatalln("Security key is not set")
	}

	if ApplicationConfig.Security.DebugMode == true && ApplicationConfig.Security.Production == true {
		utils.WarningLogger.Print("Debug mode is enabled in production mode")
	}
}

// checkAuthenticationConfig ---
//
// This function is responsible for checking the authentication config.
// For authentication to work, a valid URL and key must be set
// And database must be enabled
func checkAuthenticationConfig() {
	if ApplicationConfig.Authentication.Enabled == true && ApplicationConfig.Authentication.AuthenticationUrl == "" {
		utils.ErrorLogger.Fatalln("Authentication URL is not set")
	}

	if ApplicationConfig.Authentication.Enabled == true && ApplicationConfig.Database.Enabled == false {
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
