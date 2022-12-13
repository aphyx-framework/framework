package routing

import (
	"github.com/gofiber/fiber/v2"
	ExampleController2 "github.com/rama-adi/RyFT-Framework/app/controllers/ExampleController"
)

// ApiRoutes ---
//
// This function is used to register all the routing for the API.
// Define all of your API routing here.
func ApiRoutes(app fiber.Router) {
	app.Get("/hello", ExampleController2.HelloHandler)
	app.Get("/echo/:message", ExampleController2.EchoHandler)
}
