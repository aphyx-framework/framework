package models

import (
	utils2 "RyftFramework/app/utils"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func RegisteredModels() []Model {
	return []Model{
		{
			Name:  "User",
			Model: User{},
			Seeder: &utils2.SeederDefinition{
				Amount: 10,
				Run: func(db *gorm.DB) error {
					password, err := utils2.HashPassword("password")

					if err != nil {
						return err
					}

					user := User{
						Name:     gofakeit.Name(),
						Email:    gofakeit.Email(),
						Password: password,
					}

					return db.Create(&user).Error
				},
			},
		},
		{
			Name:   "Personal Access Token",
			Model:  PersonalAccessToken{},
			Seeder: nil,
		},
	}
}

type Model struct {
	Name   string
	Model  interface{}
	Seeder *utils2.SeederDefinition
}
