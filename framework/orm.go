package framework

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

// ConnectToDatabase ---
//
// This function is used to connect to the database. Ryft uses MySQL as the default database.
// Database is handled by GORM.
// Find the docs here: https://gorm.io
func connectDatabase() {
	if ApplicationConfig.Database.Enabled {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			ApplicationConfig.Database.Username,
			ApplicationConfig.Database.Password,
			ApplicationConfig.Database.Host,
			ApplicationConfig.Database.Port,
			ApplicationConfig.Database.Name,
		)
		database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		Db = database
	}
}
