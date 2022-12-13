package routing

import (
	ExampleController2 "RyftFramework/app/controllers/ExampleController"
	"github.com/gofiber/fiber/v2"
)

// ApiRoutes ---
//
// This function is used to register all the routing for the API.
// Define all of your API routing here.
func ApiRoutes(app fiber.Router) {
	app.Get("/hello", ExampleController2.HelloHandler)
	app.Get("/echo/:message", ExampleController2.EchoHandler)
}
