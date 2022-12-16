package framework

import (
	"github.com/rama-adi/RyFT-Framework/app"
	"github.com/rama-adi/RyFT-Framework/framework/caching"
	"github.com/rama-adi/RyFT-Framework/framework/configuration"
	"github.com/rama-adi/RyFT-Framework/framework/database"
	"github.com/rama-adi/RyFT-Framework/framework/fiberServer"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"github.com/rama-adi/RyFT-Framework/framework/router"
	"github.com/rama-adi/RyFT-Framework/framework/startupPrinter"
	"github.com/rama-adi/RyFT-Framework/framework/utils"
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

		// Populate the app package with the frameworks essential dependencies
		// To avoid cyclic dependencies if we were to use the framework package
		fx.Populate(&app.DB),
		fx.Populate(&app.Config),
		fx.Populate(&app.Logger),
		fx.Populate(&app.Utilities),
		fx.Populate(&app.CacheTable),

		// Load user defined dependencies
		app.Dependencies,

		fx.Provide(utils.InitializeFrameworkUtils),
		fx.Provide(fiberServer.NewFiberHttpServer),
		fx.Invoke(router.RegisterAllRoutes),
		fx.Provide(caching.LoadCacheTable),
		fx.Invoke(startupPrinter.PrintStartupInfo),
		fx.Invoke(fiberServer.EnableFiberServer),
	).Run()
}
