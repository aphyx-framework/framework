package migration

import (
	"RyftFramework/configuration"
	"RyftFramework/database"
	"RyftFramework/models"
	"RyftFramework/utils"
	"github.com/BurntSushi/toml"
	"log"
)

var config configuration.Configuration

func RunMigrator(fresh bool, seed bool) {
	utils.LoadLogger()
	utils.InfoLogger.Println("Migration Started")
	bootstrap()

	if fresh {
		utils.InfoLogger.Println("Doing a fresh migration..")
		dropAllTables()
	}

	doMigrations()

	if seed {
		utils.InfoLogger.Println("Seeding the migration..")
		runSeeder()
	}

}

// dropAllTables ---
//
// This function is responsible for dropping all the tables defined in RegisterModel
// If you pass -fresh flag, this function will run
func dropAllTables() {
	for _, model := range models.RegisteredModels() {
		utils.InfoLogger.Println("Dropping table for model: ", model.Name)
		err := database.DB.Migrator().DropTable(model.Model)

		if err != nil {
			utils.ErrorLogger.Println("Failed to drop table for model: ", model.Name)
			utils.ErrorLogger.Println(err)
		} else {
			utils.InfoLogger.Println("✓ Dropped table for model: ", model.Name)
		}
	}
}

func doMigrations() {
	for _, model := range models.RegisteredModels() {
		utils.InfoLogger.Println("Migrating the model: ", model.Name)
		err := database.DB.Migrator().CreateTable(model.Model)

		if err != nil {
			utils.ErrorLogger.Println("Failed to create table for model: ", model.Name)
			utils.ErrorLogger.Println(err)
		} else {
			utils.InfoLogger.Println("✓ Created table for model: ", model.Name)
		}
	}
}

// bootstrap ---
//
// Bootstrap the bare minimum for doing migration
func bootstrap() {
	_, err := toml.DecodeFile("./config.toml", &config)
	if err != nil {
		log.Fatalln("Failed to load framework config file!", err)
	}

	database.ConnectDatabase()
}
