package AuthController

import (
	"RyftFramework/database"
	"RyftFramework/models"
	"RyftFramework/utils"
	"encoding/json"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/mail"
)

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Repeat   string `json:"repeat"`
}

func RegisterHandler(c *fiber.Ctx) error {
	var user UserRegister
	_ = json.Unmarshal(c.Body(), &user)

	err := user.performValidation()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    err,
		})
	}

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: utils.HashPassword(user.Password),
	}

	register, err := newUser.Register()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.HttpResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(utils.HttpResponse{
		Success: true,
		Message: "User registered successfully",
		Data:    register,
	})
}

func (u UserRegister) performValidation() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&u.Email, validation.Required, validation.Length(1, 255), validation.By(func(value interface{}) error {
			_, err := mail.ParseAddress(value.(string))

			if err != nil {
				return errors.New("invalid email address")
			}

			var user models.User
			if err := database.DB.Where("email = ?", value.(string)).First(&user).Error; err == gorm.ErrRecordNotFound {
				return nil
			} else {
				return errors.New("email already in use")
			}
		})),
		validation.Field(&u.Password, validation.Required, validation.Length(1, 255)),
		validation.Field(&u.Repeat, validation.Required, validation.Length(1, 255), validation.By(func(value interface{}) error {
			if value.(string) != u.Password {
				return errors.New("passwords do not match")
			}
			return nil
		})),
	)
}
