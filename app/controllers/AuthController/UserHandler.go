package AuthController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rama-adi/RyFT-Framework/app/models"
	"github.com/rama-adi/RyFT-Framework/app/utils"
)

func UserHandler(c *fiber.Ctx) error {

	user := models.User{}.LoggedInUser(c)

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	})
}
