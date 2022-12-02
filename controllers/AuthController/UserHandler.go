package AuthController

import (
	"RyftFramework/models"
	"RyftFramework/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func UserHandler(c *fiber.Ctx) error {

	authorizationHeader := c.Get("Authorization", "")

	if authorizationHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.HttpResponse{
			Success: false,
			Message: "Missing Authorization header",
			Data:    nil,
		})
	}

	rep := strings.Replace(authorizationHeader, "Bearer ", "", 1)

	user, err := models.User{}.FromAccessToken(rep)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	})

}
