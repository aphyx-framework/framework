package AuthController

import (
	"RyftFramework/models"
	"RyftFramework/utils"
	"github.com/gofiber/fiber/v2"
)

func LogoutHandler(c *fiber.Ctx) error {

	token := models.User{}.LoggedInAccessToken(c)
	err := models.PersonalAccessToken{}.RevokeToken(token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "Successfully logged out",
		Data:    nil,
	})
}
