package routing

import (
	"RyftFramework/controllers/ExampleController"
	"github.com/gofiber/fiber/v2"
)

// ApiRoutes ---
//
// This function is used to register all the routing for the API.
// Define all of your API routing here.
func ApiRoutes(app fiber.Router) {
	app.Get("/hello", ExampleController.HelloHandler)
	app.Get("/echo/:message", ExampleController.EchoHandler)
}
