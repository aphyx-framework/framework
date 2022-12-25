package AuthController

import (
	"github.com/aphyx-framework/framework/app"
	"github.com/aphyx-framework/framework/app/cache"
	models2 "github.com/aphyx-framework/framework/app/models"
	"github.com/aphyx-framework/framework/framework/utils"
	"github.com/gofiber/fiber/v2"
)

func LogoutHandler(c *fiber.Ctx) error {

	token := models2.User{}.LoggedInAccessToken(c)
	err := models2.PersonalAccessToken{}.RevokeToken(token)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	// Remove the token and user from the cache
	err = app.CacheTable[cache.AuthToken].BustCache(token)

	if err != nil {
		app.Logger.ErrorLogger.Println(err)
	}

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "Successfully logged out",
		Data:    nil,
	})
}
