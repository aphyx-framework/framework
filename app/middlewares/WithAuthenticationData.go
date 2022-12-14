package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rama-adi/RyFT-Framework/app"
	"github.com/rama-adi/RyFT-Framework/app/models"
	"github.com/rama-adi/RyFT-Framework/app/utils"
	"strings"
	"time"
)

func WithAuthenticationData(c *fiber.Ctx) error {

	authorizationHeader := c.Get("Authorization", "")
	var user *models.User

	if authorizationHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.HttpResponse{
			Success: false,
			Message: "Missing Authorization header",
			Data:    nil,
		})
	}

	rep := strings.Replace(authorizationHeader, "Bearer ", "", 1)

	// cache the user data for the specified token duration
	userCache, err := app.CacheTable.CacheOrMake(app.CacheTable.Auth, "user:bytoken:"+rep, func() (interface{}, error, time.Duration) {
		fromAccessToken, fromAccessTokenError := models.User{}.FromAccessToken(rep)
		return fromAccessToken, fromAccessTokenError, time.Until(models.PersonalAccessToken{}.Find(rep).ExpiresAt)
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	user = userCache.(*models.User)

	c.Locals("accessToken", rep)
	c.Locals("user", user)

	return c.Next()
}
