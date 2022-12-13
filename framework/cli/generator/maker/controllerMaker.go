package maker

import (
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"os"
)

func ControllerMakerInfo(logger logging.ApplicationLogger) {
	logger.InfoLogger.Println("Make: controller")
	logger.InfoLogger.Println("Make a controller with handler. If the controller folder exists, it will be skipped and only the handler will be created")
	logger.InfoLogger.Println("Usage: go run enterypoint.go make controller <controllerName> <handlerName>")
	logger.InfoLogger.Println("Example: go run enterypoint.go make controller UserController UserHandler")
	os.Exit(0)
}

func ControllerMaker(controllerName string, moduleName string, handlerName string, logger logging.ApplicationLogger) {
	logger.InfoLogger.Println("Creating controller", controllerName)
	checkOrMakeDirectory("app/controllers/"+controllerName, logger)

	logger.InfoLogger.Println("Creating handler", handlerName)
	handlerFile := loadStubFile("handler", logger)
	handlerFile = replaceAllPlaceholders(handlerFile, []PlaceholderReplacer{
		{
			Placeholder: "__CONTROLLER_PKG_NAME__",
			Replacement: controllerName,
		},
		{
			Placeholder: "__APP_BASE_PKG__",
			Replacement: moduleName,
		},
		{
			Placeholder: "__HANDLER_PKG_NAME__",
			Replacement: handlerName,
		},
	})
	writeToFile("app/controllers/"+controllerName, handlerName+".go", handlerFile, logger)
}
