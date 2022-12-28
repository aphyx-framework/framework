package MakeControllerCommand

import (
	"github.com/aphyx-framework/framework/framework/logging"
	"github.com/aphyx-framework/framework/framework/utils"
	"os"
)

func makeController(
	controllerName string,
	handlerName string,
	logger logging.ApplicationLogger,
	stringUtils utils.Strings,
	modfile utils.Modfile,
) {
	logger.InfoLogger.Println("Creating controller", controllerName)
	checkOrMakeControllerDirectory("app/controllers/"+controllerName, logger)

	logger.InfoLogger.Println("Creating handler", handlerName)
	handlerStub := loadHandlerStub(logger)
	handlerStub = stringUtils.Replace(handlerStub, []utils.PlaceholderReplacer{
		{
			Find:      "__CONTROLLER_PKG_NAME__",
			ReplaceTo: controllerName,
		},
		{
			Find:      "__APP_BASE_PKG__",
			ReplaceTo: modfile.GetModuleName(),
		},
		{
			Find:      "__HANDLER_PKG_NAME__",
			ReplaceTo: handlerName,
		},
	})
	writeController("app/controllers/"+controllerName, handlerName+".go", handlerStub, logger)
	logger.InfoLogger.Println("Controller created successfully!")
}

func checkOrMakeControllerDirectory(directory string, logger logging.ApplicationLogger) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		mkdirError := os.Mkdir(directory, 0755)
		if mkdirError != nil {
			logger.ErrorLogger.Panicf("Failed to create directory %s", directory, mkdirError)
		}
	}
}

func loadHandlerStub(logger logging.ApplicationLogger) string {
	fileByte, err := os.ReadFile("./framework/commands/MakeControllerCommand/handler.stub")
	if err != nil {
		logger.ErrorLogger.Panicln("Failed to read stub file", err)
	}
	return string(fileByte)
}

func writeController(directory string, fileName string, content string, logger logging.ApplicationLogger) {
	file, err := os.Create(directory + "/" + fileName)
	if err != nil {
		logger.ErrorLogger.Panicln("Failed to create file "+fileName, err)
	}

	defer func(file *os.File) {
		fileMakerError := file.Close()
		if fileMakerError != nil {
			logger.ErrorLogger.Panicln("Failed to close file "+fileName, fileMakerError)
		}
	}(file)

	_, err = file.WriteString(content)
	if err != nil {
		logger.ErrorLogger.Panicln("Failed to write to file "+fileName, err)
	}
}
