package main

import (
	"github.com/rama-adi/RyFT-Framework/app"
	"github.com/rama-adi/RyFT-Framework/framework/bootstrapper"
	"github.com/rama-adi/RyFT-Framework/framework/cli"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"github.com/rama-adi/RyFT-Framework/framework/database"
	"github.com/rama-adi/RyFT-Framework/framework/fiberServer"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"github.com/rama-adi/RyFT-Framework/framework/router"
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
