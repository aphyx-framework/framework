package database

import "gorm.io/gorm"

type SeederDefinition struct {
	Amount int
	Run    func(*gorm.DB) error
}
