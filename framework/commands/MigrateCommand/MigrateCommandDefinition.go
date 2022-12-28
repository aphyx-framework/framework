package MigrateCommand

import (
	"github.com/TwiN/go-color"
	"github.com/aphyx-framework/framework/framework/cli"
	"github.com/aphyx-framework/framework/framework/logging"
	"gorm.io/gorm"
)

func Definition(registry cli.Registry, logger logging.ApplicationLogger, db *gorm.DB) {
	cmd := cli.Command{
		Command:     "migrate",
		Title:       "Database Migration and Seeder Command",
		Description: "This is a wrapper for GORM to migrate and seed the database. It is recommended to use this command instead of GORM directly to ensure that the database is migrated and seeded correctly.",
		Args: []cli.CommandArgument{
			{
				Name: "fresh",
				Description: "This will drop all tables and re-run all migrations.\n" +
					"This is useful for when you have made changes to your migrations" +
					"\nand want to start fresh.",
				Required: false,
			},
			{
				Name: "seed",
				Description: "This will run the database seeder.\n" +
					"This will fill the database with dummy data that you've defined in the seeder.",
				Required: false,
			},
		},
		ExampleUsage: map[string]string{
			"migrate":              "This will run the migration without dropping the tables or running the seeder.",
			"migrate fresh:":       "This will drop all tables and re-run all migrations.",
			"migrate fresh: seed:": "This will drop all tables and re-run all migrations, then run the seeder.",
			"migrate seed:":        "This will run the seeder only.",
		},
		Handler: func(c cli.CommandArgumentValue) {
			logger.InfoLogger.Println(color.GreenBackground + color.Black + "                  - Migration Started -                  " + color.Reset)

			// If the user wants to drop all tables, we will do that first
			if c.ArgumentExist("fresh") {
				logger.InfoLogger.Println(color.YellowBackground + color.Black +
					" [V] " + color.Reset + " Doing a fresh migration..")
				dropAllTables(logger, db)
			}

			// If the user only wants to run the seeder, we don't need to run the migration
			if c.ArgumentOnly("seed") == false || c.NoArguments() {
				doMigrations(logger, db)
			}

			// Run the seeder
			if c.ArgumentExist("seed") {
				logger.InfoLogger.Println(color.YellowBackground + color.Black +
					" [V] " + color.Reset + " Seeding the migration..")
				runSeeder(logger, db)
			}

			logger.InfoLogger.Println(color.GreenBackground + color.Black + "                   - Migration  Done -                   " + color.Reset)
		},
	}
	registry.AddCommand(cmd)
}
