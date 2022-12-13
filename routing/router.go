package routing

import (
	"RyftFramework/bootstrapper/logging"
	"RyftFramework/configuration"
	"RyftFramework/middlewares"
	"RyftFramework/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// LoadApiRoutes --
//
// This function is responsible for loading all the API routing.
// Ryft is primarily an API framework so all the routing are loaded here
func LoadApiRoutes(app *fiber.App) {
	apiRoutes := app.Group("/api")
	ApiRoutes(apiRoutes)
}

// LoadAuthRoute --
//
// This function is responsible for loading all the Authentication routing.
// If auth is enabled, then the auth routing will be loaded
// If not, then it will return a 404
func LoadAuthRoute(app *fiber.App, logger logging.ApplicationLogger, config configuration.Configuration) {
	auth := app.Group(config.Authentication.AuthenticationUrl)
	auth.Use(func(c *fiber.Ctx) error {
		if config.Authentication.Enabled == false {
			if config.Security.DebugMode == true {
				logger.ErrorLogger.Print("Trying to access authentication route while authentication is disabled")
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

	if config.Authentication.Enabled == false {
		// Catch all route when authentication is disabled
		auth.All("*", func(c *fiber.Ctx) error {
			return c.SendString("")
		})
	}

	loginAuth := app.Group(config.Authentication.AuthenticationUrl + "/user")
	loginAuth.Use(func(c *fiber.Ctx) error {
		if config.Authentication.Enabled == false {
			if config.Security.DebugMode == true {
				logger.ErrorLogger.Print("Trying to access authentication route while authentication is disabled")
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

	if config.Authentication.Enabled == false {
		// Catch all route when authentication is disabled
		loginAuth.All("*", func(c *fiber.Ctx) error {
			return c.SendString("")
		})
	}

	loginAuth.Use(middlewares.WithAuthenticationData)

	AuthThatNeedsLogin(loginAuth)
	AuthRoutes(auth)
}
