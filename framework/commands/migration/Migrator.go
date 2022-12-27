package migration

import (
	"github.com/TwiN/go-color"
	"github.com/aphyx-framework/framework/app/models"
	"github.com/aphyx-framework/framework/framework/logging"
	"gorm.io/gorm"
)

func RunMigrator(fresh bool, seed bool, logger logging.ApplicationLogger, db *gorm.DB) {
	logger.InfoLogger.Println(color.GreenBackground + color.Black + "                  - Migration Started -                  " + color.Reset)

	if fresh {
		logger.InfoLogger.Println(color.YellowBackground + color.Black +
			" [V] " + color.Reset + " Doing a fresh migration..")
		dropAllTables(logger, db)
	}

	doMigrations(logger, db)

	if seed {
		logger.InfoLogger.Println(color.YellowBackground + color.Black +
			" [V] " + color.Reset + " Seeding the migration..")
		runSeeder(logger, db)
	}

	logger.InfoLogger.Println(color.GreenBackground + color.Black + "                   - Migration  Done -                   " + color.Reset)

}

// dropAllTables ---
//
// This function is responsible for dropping all the tables defined in RegisterModel
// If you pass -fresh flag, this function will run
func dropAllTables(logger logging.ApplicationLogger, db *gorm.DB) {
	for _, model := range models.RegisteredModels() {
		logger.InfoLogger.Println(color.CyanBackground+color.Black+
			" [O] "+color.Reset+" Dropping table for model: ", model.Name)
		err := db.Migrator().DropTable(model.Model)

		if err != nil {
			logger.ErrorLogger.Println("Failed to drop table for model: ", model.Name)
			logger.ErrorLogger.Println(err)
		} else {
			logger.InfoLogger.Println(color.GreenBackground+color.Black+
				" [✓] "+color.Reset+" Dropped table for model: ", model.Name)
		}
	}
}

func doMigrations(logger logging.ApplicationLogger, db *gorm.DB) {
	for _, model := range models.RegisteredModels() {
		logger.InfoLogger.Println(color.CyanBackground+color.Black+
			" [O] "+color.Reset+" Migrating the model: ", model.Name)
		err := db.Migrator().CreateTable(model.Model)

		if err != nil {
			logger.ErrorLogger.Println("Failed to create table for model: ", model.Name)
			logger.ErrorLogger.Println(err)
		} else {
			logger.InfoLogger.Println(color.GreenBackground+color.Black+
				" [✓] "+color.Reset+" Created table for model: ", model.Name)
		}
	}
}
