package models

import (
	"github.com/aphyx-framework/framework/app"
	"github.com/aphyx-framework/framework/framework/database"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func RegisteredModels() []Model {
	return []Model{
		{
			Name:  "User",
			Model: User{},
			Seeder: &database.SeederDefinition{
				Amount: 10,
				Run: func(db *gorm.DB) error {
					password, err := app.Utilities.Crypto.HashPassword("password")

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
	Seeder *database.SeederDefinition
}
