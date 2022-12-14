package maker

import (
	"github.com/TwiN/go-color"
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

func writeToFile(directory string, fileName string, content string, logger logging.ApplicationLogger) {
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

func replaceAllPlaceholders(stub string, placeholderReplacer []PlaceholderReplacer) string {
	for _, placeholder := range placeholderReplacer {
		stub = strings.ReplaceAll(stub, placeholder.Placeholder, placeholder.Replacement)
	}

	return stub
}

func logSuccess(message string, logger logging.ApplicationLogger) {
	logger.InfoLogger.Println(color.GreenBackground + color.Black +
		" [✓] " + color.Reset + " " + message)
}
