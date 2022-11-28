package ExampleController

import (
	"RyftFramework/utils"
	"github.com/gofiber/fiber/v2"
)

// EchoHandler ---
//
// An example of a simple echo handler.
// This handler will return the message that was passed in the URL.
func EchoHandler(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "Echo: " + utils.DecodeUrlParam(c.Params("message")),
		Data:    nil,
	})
}
