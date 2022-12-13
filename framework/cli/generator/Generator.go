package generator

import (
	"github.com/rama-adi/RyFT-Framework/framework/cli/generator/maker"
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"golang.org/x/mod/modfile"
	"os"
)

func getModuleName(logger logging.ApplicationLogger) string {
	goModBytes, err := os.ReadFile("./go.mod")
	if err != nil {
		logger.ErrorLogger.Fatalln("Failed to read go.mod file", err)
	}

	modName := modfile.ModulePath(goModBytes)

	return modName
}

func Generator(generatorType string, logger logging.ApplicationLogger) {
	moduleName := getModuleName(logger)

	switch generatorType {
	case "controller":

		if len(os.Args) < 4 {
			logger.ErrorLogger.Println("Insufficient arguments. Please read the documentation below")
			maker.ControllerMakerInfo(logger)
		} else {
			handlerName := os.Args[4]
			maker.ControllerMaker(os.Args[3], moduleName, handlerName, logger)
		}
	}
}
