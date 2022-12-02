package routing

import (
	"RyftFramework/controllers/AuthController"
	"github.com/gofiber/fiber/v2"
)

// AuthRoutes ---
//
// This function is used to register all the routing related to authentication.
// If authentication is not enabled, this will not be called
func AuthRoutes(route fiber.Router) {
	route.Post("/login", AuthController.LoginHandler)
	route.Get("/user", AuthController.UserHandler)
}
