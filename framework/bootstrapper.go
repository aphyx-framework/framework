package framework

import (
	"github.com/rama-adi/RyFT-Framework/app"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"github.com/rama-adi/RyFT-Framework/framework/database"
	"github.com/rama-adi/RyFT-Framework/framework/fiberServer"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"github.com/rama-adi/RyFT-Framework/framework/router"
	"github.com/rama-adi/RyFT-Framework/framework/startupPrinter"
	"go.uber.org/fx"
)

func RunWebApplication(enableNopLogger bool) {
	nop := fx.Options()

	if enableNopLogger == false {
		nop = fx.Options(fx.NopLogger)
	}

	fx.New(
		nop,
		fx.Provide(configuration.NewConfiguration),
		fx.Provide(logging.NewLogger),
		fx.Provide(database.NewDbConnection),

		// Populate the app package with the configuration, logger and database connection
		fx.Populate(&app.DB),
		fx.Populate(&app.Config),
		fx.Populate(&app.Logger),

		// Load user defined dependencies
		app.Dependencies,

		fx.Provide(fiberServer.NewFiberHttpServer),
		fx.Invoke(startupPrinter.AllBootstrapper),
		fx.Invoke(router.RegisterAllRoutes),
		fx.Invoke(fiberServer.EnableFiberServer),
	).Run()
}
