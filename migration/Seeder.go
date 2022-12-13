package migration

import (
	"RyftFramework/bootstrapper/logging"
	"RyftFramework/di"
	"RyftFramework/models"
	"RyftFramework/utils"
)

func runSeeder() {
	logger := di.Dependency.Get(di.Logger).(logging.ApplicationLogger)
	for _, model := range models.RegisteredModels() {
		if model.Seeder != nil {
			logger.InfoLogger.Println("Seeding table for model: ", model.Name)
			doSeeding(*model.Seeder)
			logger.InfoLogger.Println("✓ Seeded table for model: ", model.Name)
		}
	}
}

func doSeeding(seed utils.SeederDefinition) {
	logger := di.Dependency.Get(di.Logger).(logging.ApplicationLogger)
	for i := 0; i < seed.Amount; i++ {
		err := seed.Run(DB)
		if err != nil {
			logger.ErrorLogger.Fatalln(err)
		}
	}
}
