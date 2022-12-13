package routing

import (
	"github.com/gofiber/fiber/v2"
	AuthController2 "github.com/rama-adi/RyFT-Framework/app/controllers/AuthController"
)

// AuthRoutes ---
//
// This function is used to register all the routing related to authentication.
// If authentication is not enabled, this will not be called
func AuthRoutes(route fiber.Router) {
	route.Post("/login", AuthController2.LoginHandler)
	route.Post("/register", AuthController2.RegisterHandler)
}

// AuthThatNeedsLogin ---
// This route requires authentication
func AuthThatNeedsLogin(route fiber.Router) {
	route.Get("/", AuthController2.UserHandler)
	route.Delete("/logout", AuthController2.LogoutHandler)
}
