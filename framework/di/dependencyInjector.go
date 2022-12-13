package di

import (
	"RyftFramework/app"
	"RyftFramework/framework/bootstrapper/database"
	"RyftFramework/framework/bootstrapper/logging"
	"RyftFramework/framework/bootstrapper/router"
	"RyftFramework/framework/configuration"
	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

var (
	FrameworkDependency di.Container
	Config              = "config"
	Database            = "database"
	Logger              = "logger"
	FiberServer         = "fiberServer"
	Router              = "router"
)

func BuildAppFull() {
	builder, _ := di.NewBuilder()
	err := builder.Add([]di.Def{
		{
			Name: Config,
			Build: func(ctn di.Container) (interface{}, error) {
				var config configuration.Configuration
				_, err := toml.DecodeFile("./config.toml", &config)

				return config, err
			},
		},
		{
			Name: Database,
			Build: func(ctn di.Container) (interface{}, error) {
				config := FrameworkDependency.Get(Config).(configuration.Configuration)
				return database.ConnectDatabase(config)
			},
		},
		{
			Name: Logger,
			Build: func(ctn di.Container) (interface{}, error) {
				return logging.LoadLogger()
			},
		},
		{
			Name: FiberServer,
			Build: func(ctn di.Container) (interface{}, error) {
				config := ctn.Get(Config).(configuration.Configuration)
				logger := ctn.Get(Logger).(logging.ApplicationLogger)

				webServer := fiber.New(fiber.Config{
					DisableStartupMessage: true,
					AppName:               config.Application.Name,
					EnablePrintRoutes:     true,
				})

				logger.InfoLogger.Print("Application started on port " + config.Application.Port)
				err := webServer.Listen(":" + config.Application.Port)
				return webServer, err
			},
		},
		{
			Name: Router,
			Build: func(ctn di.Container) (any, error) {
				webServer := ctn.Get(FiberServer).(*fiber.App)
				logger := ctn.Get(Logger).(logging.ApplicationLogger)
				config := ctn.Get(Config).(configuration.Configuration)
				router.LoadAuthRoute(webServer, logger, config)
				router.LoadApiRoutes(webServer)
				return nil, nil
			},
		},
	}...)

	if err != nil {
		panic(err)
	}

	dependencyContainer := builder.Build()

	FrameworkDependency = dependencyContainer
	app.Container = dependencyContainer
}

func BuildForMigrator() {
	builder, _ := di.NewBuilder()
	err := builder.Add([]di.Def{
		{
			Name: "config",
			Build: func(ctn di.Container) (interface{}, error) {
				var config configuration.Configuration
				_, err := toml.DecodeFile("./config.toml", &config)

				return config, err
			},
		},
		{
			Name: "database",
			Build: func(ctn di.Container) (interface{}, error) {
				config := FrameworkDependency.Get(Config).(configuration.Configuration)
				return database.ConnectDatabase(config)
			},
		},
		{
			Name: "logger",
			Build: func(ctn di.Container) (interface{}, error) {
				return logging.LoadLogger()
			},
		},
	}...)

	if err != nil {
		panic(err)
	}

	FrameworkDependency = builder.Build()
}
