package models

import (
	"RyftFramework/database"
	"RyftFramework/utils"
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

func (_ User) Login(email string, password string) (*User, error) {
	var user User

	database.DB.Where("email = ?", email).First(&user)

	if utils.CheckPasswordHash(password, user.Password) {
		return &user, nil
	}

	return nil, errors.New("invalid email or password")
}
