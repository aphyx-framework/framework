package routing

import (
	"RyftFramework/configuration"
	"RyftFramework/middlewares"
	"RyftFramework/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// LoadRouter ---
//
// This function is responsible for loading all the routing.
// It will load the API routing and the Authentication routing
// It also supports middleware
func LoadRouter(app *fiber.App) {
	loadAuthRoute(app)
	loadApiRoutes(app)
}

// loadApiRoutes --
//
// This function is responsible for loading all the API routing.
// Ryft is primarily an API framework so all the routing are loaded here
func loadApiRoutes(app *fiber.App) {
	apiRoutes := app.Group("/api")
	ApiRoutes(apiRoutes)
}

// loadAuthRoute --
//
// This function is responsible for loading all the Authentication routing.
// If auth is enabled, then the auth routing will be loaded
// If not, then it will return a 404
func loadAuthRoute(app *fiber.App) {
	auth := app.Group(configuration.ApplicationConfig.Authentication.AuthenticationUrl)
	auth.Use(func(c *fiber.Ctx) error {
		if configuration.ApplicationConfig.Authentication.Enabled == false {
			if configuration.ApplicationConfig.Security.DebugMode == true {
				utils.ErrorLogger.Print("Trying to access authentication route while authentication is disabled")
				return c.Status(http.StatusInternalServerError).JSON(utils.HttpResponse{
					Success: false,
					Message: "Authentication is not enabled. Check your config.toml file!",
					Data:    nil,
				})
			}
			return c.Status(http.StatusNotFound).JSON(utils.HttpResponse{
				Success: false,
				Message: "Not found",
				Data:    nil,
			})
		}
		return c.Next()
	})

	if configuration.ApplicationConfig.Authentication.Enabled == false {
		// Catch all route when authentication is disabled
		auth.All("*", func(c *fiber.Ctx) error {
			return c.SendString("")
		})
	}

	loginAuth := app.Group(configuration.ApplicationConfig.Authentication.AuthenticationUrl + "/user")
	loginAuth.Use(func(c *fiber.Ctx) error {
		if configuration.ApplicationConfig.Authentication.Enabled == false {
			if configuration.ApplicationConfig.Security.DebugMode == true {
				utils.ErrorLogger.Print("Trying to access authentication route while authentication is disabled")
				return c.Status(http.StatusInternalServerError).JSON(utils.HttpResponse{
					Success: false,
					Message: "Authentication is not enabled. Check your config.toml file!",
					Data:    nil,
				})
			}
			return c.Status(http.StatusNotFound).JSON(utils.HttpResponse{
				Success: false,
				Message: "Not found",
				Data:    nil,
			})
		}
		return c.Next()
	})

	if configuration.ApplicationConfig.Authentication.Enabled == false {
		// Catch all route when authentication is disabled
		loginAuth.All("*", func(c *fiber.Ctx) error {
			return c.SendString("")
		})
	}

	loginAuth.Use(middlewares.WithAuthenticationData)

	AuthThatNeedsLogin(loginAuth)
	AuthRoutes(auth)
}
