package migration

import (
	"RyftFramework/app/models"
	"RyftFramework/app/utils"
	"RyftFramework/framework/logging"
)

func runSeeder() {
	logger := di.FrameworkDependency.Get(di.Logger).(logging.ApplicationLogger)
	for _, model := range models.RegisteredModels() {
		if model.Seeder != nil {
			logger.InfoLogger.Println("Seeding table for model: ", model.Name)
			doSeeding(*model.Seeder)
			logger.InfoLogger.Println("✓ Seeded table for model: ", model.Name)
		}
	}
}

func doSeeding(seed utils.SeederDefinition) {
	logger := di.FrameworkDependency.Get(di.Logger).(logging.ApplicationLogger)
	for i := 0; i < seed.Amount; i++ {
		err := seed.Run(DB)
		if err != nil {
			logger.ErrorLogger.Fatalln(err)
		}
	}
}
