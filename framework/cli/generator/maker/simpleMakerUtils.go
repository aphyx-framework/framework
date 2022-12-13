package maker

import (
	"github.com/rama-adi/RyFT-Framework/framework/logging"
	"os"
	"strings"
)

type PlaceholderReplacer struct {
	Placeholder string
	Replacement string
}

func loadStubFile(stubFile string, logger logging.ApplicationLogger) string {
	fileByte, err := os.ReadFile("./framework/cli/generator/stubs/" + stubFile + ".stub")
	if err != nil {
		logger.ErrorLogger.Panicln("Failed to read stub file", err)
	}
	return string(fileByte)
}

func checkOrMakeDirectory(directory string, logger logging.ApplicationLogger) {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		mkdirError := os.Mkdir(directory, 0755)
		if mkdirError != nil {
			logger.ErrorLogger.Panicf("Failed to create directory %s", directory, mkdirError)
		}
	}
}

func replaceAllPlaceholders(stub string, placeholderReplacer []PlaceholderReplacer) string {
	for _, placeholder := range placeholderReplacer {
		stub = strings.ReplaceAll(stub, placeholder.Placeholder, placeholder.Replacement)
	}

	return stub
}
