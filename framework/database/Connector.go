package database

import (
	"fmt"
	"github.com/aphyx-framework/framework/framework/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDbConnection(config configuration.Configuration) (*gorm.DB, error) {
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
