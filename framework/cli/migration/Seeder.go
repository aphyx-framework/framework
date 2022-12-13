package migration

import (
	"github.com/rama-adi/RyFT-Framework/app/models"
	"github.com/rama-adi/RyFT-Framework/app/utils"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"gorm.io/gorm"
)

func runSeeder(logger logging.ApplicationLogger, db *gorm.DB) {

	for _, model := range models.RegisteredModels() {
		if model.Seeder != nil {
			logger.InfoLogger.Println("Seeding table for model: ", model.Name)
			doSeeding(*model.Seeder, logger, db)
			logger.InfoLogger.Println("✓ Seeded table for model: ", model.Name)
		}
	}
}

func doSeeding(seed utils.SeederDefinition, logger logging.ApplicationLogger, db *gorm.DB) {
	for i := 0; i < seed.Amount; i++ {
		err := seed.Run(db)
		if err != nil {
			logger.ErrorLogger.Fatalln(err)
		}
	}
}
