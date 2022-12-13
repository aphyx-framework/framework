package migration

import (
	"github.com/rama-adi/RyFT-Framework/app/models"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"gorm.io/gorm"
)

func RunMigrator(fresh bool, seed bool, logger logging.ApplicationLogger, db *gorm.DB) {
	logger.InfoLogger.Println("Migration Started")

	if fresh {
		logger.InfoLogger.Println("Doing a fresh migration..")
		dropAllTables(logger, db)
	}

	doMigrations(logger, db)

	if seed {
		logger.InfoLogger.Println("Seeding the migration..")
		runSeeder(logger, db)
	}

}

// dropAllTables ---
//
// This function is responsible for dropping all the tables defined in RegisterModel
// If you pass -fresh flag, this function will run
func dropAllTables(logger logging.ApplicationLogger, db *gorm.DB) {
	for _, model := range models.RegisteredModels() {
		logger.InfoLogger.Println("Dropping table for model: ", model.Name)
		err := db.Migrator().DropTable(model.Model)

		if err != nil {
			logger.ErrorLogger.Println("Failed to drop table for model: ", model.Name)
			logger.ErrorLogger.Println(err)
		} else {
			logger.InfoLogger.Println("✓ Dropped table for model: ", model.Name)
		}
	}
}

func doMigrations(logger logging.ApplicationLogger, db *gorm.DB) {
	for _, model := range models.RegisteredModels() {
		logger.InfoLogger.Println("Migrating the model: ", model.Name)
		err := db.Migrator().CreateTable(model.Model)

		if err != nil {
			logger.ErrorLogger.Println("Failed to create table for model: ", model.Name)
			logger.ErrorLogger.Println(err)
		} else {
			logger.InfoLogger.Println("✓ Created table for model: ", model.Name)
		}
	}
}
