package MakeModelCommand

import (
	"github.com/aphyx-framework/framework/framework/logging"
	"github.com/aphyx-framework/framework/framework/utils"
	"os"
)

func makeModel(
	modelName string,
	logger logging.ApplicationLogger,
	stringUtils utils.Strings,
	modfile utils.Modfile,
) {
	logger.InfoLogger.Println("Creating controller", modelName)

	modelStub := loadModelStub(logger)
	modelStub = stringUtils.Replace(modelStub, []utils.PlaceholderReplacer{
		{
			Find:      "__MODEL_NAME__",
			ReplaceTo: modelName,
		},
		{
			Find:      "__APP_BASE_PKG__",
			ReplaceTo: modfile.GetModuleName(),
		},
	})
	writeController("app/models", modelName+".go", modelStub, logger)
	logger.InfoLogger.Println("Model created successfully!")
	logger.InfoLogger.Println("Please register the model in app/models/RegisterModel.go -> registerModels()")
	println("{")
	println("\tName:   \"" + modelName + "\",")
	println("\tModel:  " + modelName + "{},")
	println("\tSeeder: nil,")
	println("},")
	logger.InfoLogger.Println("If you need to seed the model, please add a seeder definition. fill the seeder with this example:")
	println("Seeder: &utils.SeederDefinition{")
	println("\tAmount: 10,")
	println("\tRun: func(db *gorm.DB) error {")
	println("\t\tmodel := " + modelName + "{} // fill the fields here")
	println("\t\treturn db.Create(&model).Error")
	println("\t},")
}

func loadModelStub(logger logging.ApplicationLogger) string {
	fileByte, err := os.ReadFile("./framework/commands/MakeModelCommand/model.stub")
	if err != nil {
		logger.ErrorLogger.Panicln("Failed to read stub file", err)
	}
	return string(fileByte)
}

func writeController(directory string, fileName string, content string, logger logging.ApplicationLogger) {
	file, err := os.Create(directory + "/" + fileName)
	if err != nil {
		logger.ErrorLogger.Panicln("Failed to create model file "+fileName, err)
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
