package di

import (
	"RyftFramework/bootstrapper/database"
	"RyftFramework/bootstrapper/logging"
	"RyftFramework/configuration"
	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

var (
	Dependency  di.Container
	Config      = "config"
	Database    = "database"
	Logger      = "logger"
	FiberServer = "fiberServer"
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
				config := Dependency.Get(Config).(configuration.Configuration)
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

				app := fiber.New(fiber.Config{
					DisableStartupMessage: true,
					AppName:               config.Application.Name,
				})

				logger.InfoLogger.Print("Application started on port " + config.Application.Port)
				err := app.Listen(":" + config.Application.Port)
				return app, err
			},
		},
	}...)

	if err != nil {
		panic(err)
	}

	Dependency = builder.Build()
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
				config := Dependency.Get(Config).(configuration.Configuration)
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

	Dependency = builder.Build()
}
