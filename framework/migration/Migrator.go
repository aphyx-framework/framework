package migration

import (
	"RyftFramework/app/models"
	"RyftFramework/framework/bootstrapper/logging"
	"RyftFramework/framework/configuration"
	"RyftFramework/framework/di"
	"fmt"
	"github.com/BurntSushi/toml"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var config configuration.Configuration
var DB *gorm.DB

func RunMigrator(fresh bool, seed bool) {
	logger := di.FrameworkDependency.Get(di.Logger).(logging.ApplicationLogger)
	logger.InfoLogger.Println("Migration Started")
	bootstrap()

	if fresh {
		logger.InfoLogger.Println("Doing a fresh migration..")
		dropAllTables()
	}

	doMigrations()

	if seed {
		logger.InfoLogger.Println("Seeding the migration..")
		runSeeder()
	}

}

// dropAllTables ---
//
// This function is responsible for dropping all the tables defined in RegisterModel
// If you pass -fresh flag, this function will run
func dropAllTables() {
	logger := di.FrameworkDependency.Get(di.Logger).(logging.ApplicationLogger)
	for _, model := range models.RegisteredModels() {
		logger.InfoLogger.Println("Dropping table for model: ", model.Name)
		err := DB.Migrator().DropTable(model.Model)

		if err != nil {
			logger.ErrorLogger.Println("Failed to drop table for model: ", model.Name)
			logger.ErrorLogger.Println(err)
		} else {
			logger.InfoLogger.Println("✓ Dropped table for model: ", model.Name)
		}
	}
}

func doMigrations() {
	logger := di.FrameworkDependency.Get(di.Logger).(logging.ApplicationLogger)
	for _, model := range models.RegisteredModels() {
		logger.InfoLogger.Println("Migrating the model: ", model.Name)
		err := DB.Migrator().CreateTable(model.Model)

		if err != nil {
			logger.ErrorLogger.Println("Failed to create table for model: ", model.Name)
			logger.ErrorLogger.Println(err)
		} else {
			logger.InfoLogger.Println("✓ Created table for model: ", model.Name)
		}
	}
}

// bootstrap ---
//
// Bootstrap the bare minimum for doing migration
func bootstrap() {
	_, err := toml.DecodeFile("./configuration/config.toml", &config)

	if err != nil {
		log.Fatalln("Failed to load framework config file!", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Name,
	)

	database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB = database
}
