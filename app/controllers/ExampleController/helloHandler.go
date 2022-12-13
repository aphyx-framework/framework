package ExampleController

import (
	"RyftFramework/app/utils"
	"github.com/gofiber/fiber/v2"
)

// HelloHandler ---
//
// Very simple example of a handler
// Just returns a string
func HelloHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "Hello World!",
		Data:    nil,
	})
}
