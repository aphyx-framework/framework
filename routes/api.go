package routes

import (
	"RyftFramework/controllers/ExampleController"
	"github.com/gofiber/fiber/v2"
)

// ApiRoutes ---
//
// This function is used to register all the routes for the API.
// Define all of your API routes here.
func ApiRoutes(app fiber.Router) {
	app.Get("/hello", ExampleController.HelloHandler)
	app.Get("/echo/:message", ExampleController.EchoHandler)
}
