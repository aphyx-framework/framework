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
	route.Post("/register", AuthController.RegisterHandler)
}

// AuthThatNeedsLogin ---
// This route requires authentication
func AuthThatNeedsLogin(route fiber.Router) {
	route.Get("/", AuthController.UserHandler)
	route.Delete("/logout", AuthController.LogoutHandler)
}
