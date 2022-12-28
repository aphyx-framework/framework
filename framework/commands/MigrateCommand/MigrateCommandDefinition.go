package MigrateCommand

import (
	"github.com/aphyx-framework/framework/framework/cli"
)

func Definition(registry cli.Registry) {
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
			"migrate seed:":        "This will run the seeder.",
		},
		Handler: func(c cli.CommandArgumentValue) {

		},
	}
	registry.AddCommand(cmd)
}
