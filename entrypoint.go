package main

import (
	"RyftFramework/app"
	"RyftFramework/framework/bootstrapper"
	"RyftFramework/framework/cli"
	"RyftFramework/framework/configuration"
	"RyftFramework/framework/database"
	"RyftFramework/framework/fiberServer"
	"RyftFramework/framework/logging"
	"RyftFramework/framework/router"
	"go.uber.org/fx"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		// If no argument is passed, start the server
		runApplication()
	} else {
		cli.RunCliApplication()
	}
}

func runApplication() {
	fx.New(
		fx.NopLogger,
		fx.Provide(configuration.NewConfiguration),
		fx.Provide(logging.NewLogger),
		fx.Provide(database.NewDbConnection),

		// Populate the app package with the configuration, logger and database connection
		fx.Populate(&app.DB),
		fx.Populate(&app.Config),
		fx.Populate(&app.Logger),

		fx.Provide(fiberServer.NewFiberHttpServer),
		fx.Invoke(bootstrapper.AllBootstrapper),
		fx.Invoke(router.RegisterAllRoutes),
		fx.Invoke(fiberServer.EnableFiberServer),
	).Run()
}

//func main() {
//
//	migratorFlag := flag.NewFlagSet("migrate", flag.ExitOnError)
//	fresh := migratorFlag.Bool("fresh", false, "Drop all table defined in RegisterModel")
//	seed := migratorFlag.Bool("seed", false, "Seed the migration with data defined in the seeder")
//
//	if len(os.Args) < 2 {
//		// If no argument is passed, start the server
//		BootstrapFramework()
//	} else {
//		switch os.Args[1] {
//		case "migrate":
//			err := migratorFlag.Parse(os.Args[2:])
//			if err != nil {
//				panic(err)
//			}
//			container.BuildForMigrator()
//			migration.RunMigrator(*fresh, *seed)
//		default:
//			panic("Unknown command")
//		}
//	}
//
//}
