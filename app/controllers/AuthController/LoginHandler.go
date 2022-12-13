package AuthController

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	models2 "github.com/rama-adi/RyFT-Framework/app/models"
	"github.com/rama-adi/RyFT-Framework/app/utils"
)

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

func LoginHandler(c *fiber.Ctx) error {

	var user UserLogin                  // Create a new instance of the UserLogin struct
	_ = json.Unmarshal(c.Body(), &user) // Unmarshal the request body into the struct

	err := user.performValidation() // Perform validation on the struct

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
	}

	token, err := user.performLogin() // Perform login

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "Token generated successfully",
		Data:    token,
	})

}

func (u UserLogin) performValidation() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, validation.Length(1, 255)),
		validation.Field(&u.Password, validation.Required, validation.Length(1, 255)),
		validation.Field(&u.Remember, validation.In(true, false)),
	)
}

func (u UserLogin) performLogin() (*models2.PersonalAccessTokenResponse, error) {
	getUser, err := models2.User{}.Login(u.Email, u.Password)

	if err != nil {
		return nil, err
	}

	token, err := models2.PersonalAccessToken{}.CreateTokenForUser(*getUser, "Personal Access Token", u.Remember)

	if err != nil {
		return nil, err
	}

	return &token, nil
}
