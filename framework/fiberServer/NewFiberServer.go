package fiberServer

import (
	"RyftFramework/framework/configuration"
	"RyftFramework/framework/logging"
	"context"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func NewFiberHttpServer(config configuration.Configuration) *fiber.App {

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		AppName:               config.Application.Name,
		EnablePrintRoutes:     true,
	})

	return app
}

func EnableFiberServer(lifecycle fx.Lifecycle, app *fiber.App, config configuration.Configuration, logger logging.ApplicationLogger) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				logger.InfoLogger.Print("Application started on port " + config.Application.Port)
				if err := app.Listen(":" + config.Application.Port); err != nil {
					panic(err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := app.Shutdown(); err != nil {
				panic(err)
			}
			return nil
		},
	})
}
