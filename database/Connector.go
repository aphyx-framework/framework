package database

import (
	"RyftFramework/configuration"
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
func ConnectDatabase() {

	if configuration.ApplicationConfig.Database.Enabled {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			configuration.ApplicationConfig.Database.Username,
			configuration.ApplicationConfig.Database.Password,
			configuration.ApplicationConfig.Database.Host,
			configuration.ApplicationConfig.Database.Port,
			configuration.ApplicationConfig.Database.Name,
		)
		database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		DB = database
	}
}
