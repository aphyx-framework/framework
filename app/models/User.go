package models

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rama-adi/RyFT-Framework/app"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	PersonalAccessToken []PersonalAccessToken `gorm:"foreignKey:UserID"`
	Name                string                `gorm:"not null"`
	Email               string                `gorm:"unique"`
	Password            string                `json:"-"`
}

func (_ User) Login(email string, password string) (*User, error) {
	var user User

	app.DB.Where("email = ?", email).First(&user)

	if app.Utilities.Crypto.VerifyPassword(password, user.Password) {
		return &user, nil
	}

	return nil, errors.New("invalid email or password")
}

func (u User) Register() (*PersonalAccessTokenResponse, error) {

	if err := app.DB.Create(&u).Error; err != nil {
		return nil, err
	}

	token, err := PersonalAccessToken{}.CreateTokenForUser(u, "Login token", false)

	if err != nil {
		return nil, err
	}

	return &token, nil

}

func (_ User) FromAccessToken(token string) (*User, error) {

	var personalAccessToken PersonalAccessToken

	enc, err := app.Utilities.Crypto.EncryptWithAppKey(token)

	if err != nil {
		return nil, err
	}

	err = app.DB.Where("token = ?", enc).Preload("User").First(&personalAccessToken).Error

	if err != nil {
		return nil, err
	}

	if personalAccessToken.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	return &personalAccessToken.User, nil
}

func (_ User) LoggedInUser(c *fiber.Ctx) *User {
	return c.Locals("user").(*User)
}

func (_ User) LoggedInAccessToken(c *fiber.Ctx) string {
	return c.Locals("accessToken").(string)
}
