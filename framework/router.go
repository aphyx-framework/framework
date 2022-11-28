package framework

import (
	"RyftFramework/routes"
	"RyftFramework/utils"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

// loadRouter ---
//
// This function is responsible for loading all the routes.
// It will load the API routes and the Authentication routes
// It also supports middleware
func loadRouter(app *fiber.App) {
	loadAuthRoute(app)
	loadApiRoutes(app)
}

// loadApiRoutes --
//
// This function is responsible for loading all the API routes.
// Ryft is primarily an API framework so all the routes are loaded here
func loadApiRoutes(app *fiber.App) {
	apiRoutes := app.Group("/api")
	routes.ApiRoutes(apiRoutes)
}

// loadAuthRoute --
//
// This function is responsible for loading all the Authentication routes.
// If auth is enabled, then the auth routes will be loaded
// If not, then it will return a 404
func loadAuthRoute(app *fiber.App) {
	auth := app.Group(ApplicationConfig.Authentication.AuthenticationUrl)
	auth.Use(func(c *fiber.Ctx) error {
		if ApplicationConfig.Authentication.Enabled == false {
			if ApplicationConfig.Security.DebugMode == true {
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

	if ApplicationConfig.Authentication.Enabled == false {
		// Catch all route when authentication is disabled
		auth.All("*", func(c *fiber.Ctx) error {
			return c.SendString("")
		})
	}

	routes.AuthRoutes(auth)
}
