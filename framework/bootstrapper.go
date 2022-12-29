package framework

import (
	"github.com/aphyx-framework/framework/app"
	"github.com/aphyx-framework/framework/framework/caching"
	"github.com/aphyx-framework/framework/framework/cli"
	"github.com/aphyx-framework/framework/framework/commands"
	"github.com/aphyx-framework/framework/framework/configuration"
	"github.com/aphyx-framework/framework/framework/database"
	"github.com/aphyx-framework/framework/framework/fiberServer"
	"github.com/aphyx-framework/framework/framework/logging"
	"github.com/aphyx-framework/framework/framework/router"
	"github.com/aphyx-framework/framework/framework/startupPrinter"
	"github.com/aphyx-framework/framework/framework/utils"
	"go.uber.org/fx"
)

func BoostrapKernel(enableFxLogger bool, cliMode bool) {

	// Initialize some empty Options for Fx
	fxLogger := fx.Options()
	server := fx.Options()

	// Fx by default logs everything to stdout, this is a workaround to disable it
	// You can enable it by setting the EnableNopLogger to true
	if enableFxLogger == false {
		fxLogger = fx.Options(fx.NopLogger)
	}

	// If the CLI flag is set to false, start the web application
	if cliMode == false {
		server = fx.Options(
			fx.Provide(fiberServer.NewFiberHttpServer),
			fx.Invoke(router.RegisterAllRoutes),
			fx.Invoke(startupPrinter.PrintStartupInfo),
			fx.Invoke(fiberServer.EnableFiberServer),
		)
	}

	// If the CLI flag is set to true, we will not start the server
	// But instead, we will invoke the CLI application
	if cliMode {
		server = fx.Options(
			fx.Provide(cli.MakeRegistry), // Make a new registry for the CLI commands
			commands.FrameworkCommands,   // Commands from the framework
			fx.Invoke(cli.RunCommand),
		)
	}

	// This is the main container, it will be used to inject all the dependencies
	// into the application
	fx.New(

		// add the Fx logger if the EnableFxLogger is set to true
		fxLogger,

		// Load essential dependencies
		fx.Provide(configuration.NewConfiguration),
		fx.Provide(logging.NewLogger),
		fx.Provide(database.NewDbConnection),
		fx.Provide(utils.InitializeFrameworkUtils),
		fx.Provide(caching.LoadCacheTable),

		// Populate the app package with the frameworks essential dependencies
		// To avoid cyclic dependencies if we were to use the framework package
		fx.Populate(&app.DB),
		fx.Populate(&app.Config),
		fx.Populate(&app.Logger),
		fx.Populate(&app.Utilities),
		fx.Populate(&app.CacheTable),

		// Load user defined dependencies
		app.Dependencies,

		// Load the server. The server var is either the commands or the http server
		server,
	).Run()
}
