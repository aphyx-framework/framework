package models

import (
	"RyftFramework/database"
	"RyftFramework/utils"
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	PersonalAccessToken []PersonalAccessToken
	Name                string
	Email               string
	Password            string
}

func (_ User) Login(email string, password string) (*User, error) {
	var user User

	database.DB.Where("email = ?", email).First(&user)

	if utils.CheckPasswordHash(password, user.Password) {
		return &user, nil
	}

	return nil, errors.New("invalid email or password")
}

func (_ User) FromAccessToken(token string) (*User, error) {
	var personalAccessToken PersonalAccessToken

	enc, err := utils.EncryptString(token)

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
