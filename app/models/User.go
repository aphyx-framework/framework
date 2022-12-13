package models

import (
	"RyftFramework/app"
	"RyftFramework/app/utils"
	"RyftFramework/framework/bootstrapper/database"
	"RyftFramework/framework/configuration"
	"errors"
	"github.com/gofiber/fiber/v2"
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

type UserWithToken struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (_ User) Login(email string, password string) (*User, error) {
	var user User

	database.DB.Where("email = ?", email).First(&user)

	if utils.CheckPasswordHash(password, user.Password) {
		return &user, nil
	}

	return nil, errors.New("invalid email or password")
}

func (u User) Register() (*UserWithToken, error) {

	if err := database.DB.Create(&u).Error; err != nil {
		return nil, err
	}

	token, err := PersonalAccessToken{}.CreateTokenForUser(u, "Login token", false)

	if err != nil {
		return nil, err
	}

	return &UserWithToken{
		Name:      u.Name,
		Email:     u.Email,
		Token:     token.Token,
		ExpiresAt: token.ExpiresAt,
	}, nil

}

func (_ User) FromAccessToken(token string) (*User, error) {
	var config = app.Container.Get("config").(configuration.Configuration)
	var personalAccessToken PersonalAccessToken

	enc, err := utils.EncryptString(token, config.Security.Key)

	if err != nil {
		return nil, err
	}

	err = database.DB.Where("token = ?", enc).Preload("User").First(&personalAccessToken).Error

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
