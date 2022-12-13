package cli

import (
	"RyftFramework/framework/cli/migration"
	"RyftFramework/framework/configuration"
	"RyftFramework/framework/database"
	"RyftFramework/framework/logging"
	"flag"
	"go.uber.org/fx"
	"gorm.io/gorm"
	"os"
)

func RunCliApplication() {
	fx.New(
		fx.NopLogger,
		fx.Provide(configuration.NewConfiguration),
		fx.Provide(logging.NewLogger),
		fx.Provide(database.NewDbConnection),

		fx.Invoke(runCliCommand),
	).Run()
}

func runCliCommand(logger logging.ApplicationLogger, config configuration.Configuration, db *gorm.DB) {

	switch os.Args[1] {
	case "migrate":
		migratorFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
		fresh := migratorFlag.Bool("fresh", false, "Drop all table defined in RegisterModel")
		seed := migratorFlag.Bool("seed", false, "Seed the migration with data defined in the seeder")
		err := migratorFlag.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
		migration.RunMigrator(*fresh, *seed, logger, db)
	default:
		panic("Unknown command")
	}

}
