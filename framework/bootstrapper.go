package framework

import (
	"github.com/aphyx-framework/framework/app"
	"github.com/aphyx-framework/framework/framework/caching"
	"github.com/aphyx-framework/framework/framework/configuration"
	"github.com/aphyx-framework/framework/framework/database"
	"github.com/aphyx-framework/framework/framework/fiberServer"
	"github.com/aphyx-framework/framework/framework/logging"
	"github.com/aphyx-framework/framework/framework/router"
	"github.com/aphyx-framework/framework/framework/startupPrinter"
	"github.com/aphyx-framework/framework/framework/utils"
	"go.uber.org/fx"
)

func BoostrapKernel(enableNopLogger bool, cli bool) {
	nop := fx.Options()
	server := fx.Options()

	if enableNopLogger == false {
		nop = fx.Options(fx.NopLogger)
	}

	if cli == false {
		server = fx.Options(
			fx.Provide(fiberServer.NewFiberHttpServer),
			fx.Invoke(router.RegisterAllRoutes),
			fx.Provide(caching.LoadCacheTable),
			fx.Invoke(startupPrinter.PrintStartupInfo),
			fx.Invoke(fiberServer.EnableFiberServer),
		)
	}

	fx.New(
		nop,
		// Load essential dependencies
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

		// If the application is not a CLI application, load the server
		server,
	).Run()
}

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
