package ExampleController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rama-adi/RyFT-Framework/app/utils"
)

// EchoHandler ---
//
// An example of a simple echo handler.
// This handler will return the message that was passed in the URL.
func EchoHandler(c *fiber.Ctx) error {
	message, err := utils.DecodeUrlParam(c.Params("message"))

	if err != nil {
		return c.Status(400).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "Echo: " + message,
		Data:    nil,
	})
}
