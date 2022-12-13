package di

import (
	"RyftFramework/bootstrapper/database"
	"RyftFramework/bootstrapper/logging"
	"RyftFramework/configuration"
	"github.com/gofiber/fiber/v2"
	"github.com/sarulabs/di"
)

var Dependency di.Container

func BuildAppFull() {
	builder, _ := di.NewBuilder()
	err := builder.Add([]di.Def{
		{
			Name: "config",
			Build: func(ctn di.Container) (interface{}, error) {
				return configuration.LoadConfigFile()
			},
		},
		{
			Name: "database",
			Build: func(ctn di.Container) (interface{}, error) {
				return database.ConnectDatabase()
			},
		},
		{
			Name: "logger",
			Build: func(ctn di.Container) (interface{}, error) {
				return logging.LoadLogger()
			},
		},
		{
			Name: "fiberServer",
			Build: func(ctn di.Container) (interface{}, error) {
				config := ctn.Get("config").(configuration.Configuration)
				logger := ctn.Get("logger").(logging.ApplicationLogger)

				app := fiber.New(fiber.Config{
					DisableStartupMessage: true,
					AppName:               config.Application.Name,
				})
				logger.InfoLogger.Print("Application started on port " + config.Application.Port)
				err := app.Listen(":" + config.Application.Port)
				return app, err
			},
		},
		{
			Name: "router",
			Build: func(ctn di.Container) (interface{}, error) {
				config := ctn.Get("config").(configuration.Configuration)
				logger := ctn.Get("logger").(logging.ApplicationLogger)

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
