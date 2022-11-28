package models

import (
	"RyftFramework/utils"
	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"
)

func RegisteredModels() []Model {
	return []Model{
		{
			Name:  "User",
			Model: User{},
			Seeder: &utils.SeederDefinition{
				Amount: 10,
				Run: func(db *gorm.DB) error {
					user := User{
						Name:     gofakeit.Name(),
						Email:    gofakeit.Email(),
						Password: gofakeit.Password(true, true, true, true, true, 32),
					}
					return db.Create(&user).Error
				},
			},
		},
	}
}

type Model struct {
	Name   string
	Model  interface{}
	Seeder *utils.SeederDefinition
}
