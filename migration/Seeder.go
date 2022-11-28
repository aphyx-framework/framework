package migration

import (
	"RyftFramework/database"
	"RyftFramework/models"
	"RyftFramework/utils"
)

func runSeeder() {
	for _, model := range models.RegisteredModels() {
		if model.Seeder != nil {
			utils.InfoLogger.Println("Seeding table for model: ", model.Name)
			doSeeding(*model.Seeder)
			utils.InfoLogger.Println("✓ Seeded table for model: ", model.Name)
		}
	}
}

func doSeeding(seed utils.SeederDefinition) {
	for i := 0; i < seed.Amount; i++ {
		err := seed.Run(database.DB)
		if err != nil {
			utils.ErrorLogger.Fatalln(err)
		}
	}
}
