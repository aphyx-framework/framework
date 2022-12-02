package AuthController

import (
	"RyftFramework/models"
	"RyftFramework/utils"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type UserLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Remember bool   `json:"remember"`
}

func LoginHandler(c *fiber.Ctx) error {

	var user UserLogin
	_ = json.Unmarshal(c.Body(), &user)

	validate := validator.New()
	err := validate.Struct(user)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse{
			Success: false,
			Message: "Validation error",
			Data:    utils.GetErrors(err),
		})
	}

	getUser, err := models.User{}.Login(user.Email, user.Password)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	token, err := models.PersonalAccessToken{}.CreateTokenForUser(*getUser, "Personal Access Token", user.Remember)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "Token generated successfully",
		Data:    token,
	})

}
