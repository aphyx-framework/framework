package maker

import (
	"github.com/rama-adi/RyFT-Framework/framework/logging"
)

func ModelMakerInfo(logger logging.ApplicationLogger) {
	logger.InfoLogger.Println("Make: model")
	logger.InfoLogger.Println("Make a model inside of your model folder")
	logger.InfoLogger.Println("Usage: go run enterypoint.go make model <modelName>")
	logger.InfoLogger.Println("Example: go run enterypoint.go make model User")
}

func ModelMaker(modelName string, moduleName string, logger logging.ApplicationLogger) {
	logger.InfoLogger.Println("Creating model", modelName)
	modelStub := loadStubFile("model", logger)
	modelStub = replaceAllPlaceholders(modelStub, []PlaceholderReplacer{
		{
			Placeholder: "__MODEL_NAME__",
			Replacement: modelName,
		},
		{
			Placeholder: "__APP_BASE_PKG__",
			Replacement: moduleName,
		},
	})

	writeToFile("app/models", modelName+".go", modelStub, logger)
	logSuccess("Model created successfully", logger)
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
