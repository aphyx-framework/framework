package database

import (
	"RyftFramework/framework/configuration"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase ---
//
// This function is used to connect to the migration. Ryft uses MySQL as the default migration.
// Database is handled by GORM.
// Find the docs here: https://gorm.io
func ConnectDatabase(config configuration.Configuration) (*gorm.DB, error) {

	if config.Database.Enabled {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.Database.Username,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
		)
		database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		return database, err
	}
	return nil, nil
}
