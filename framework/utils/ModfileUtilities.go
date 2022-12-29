package utils

import (
	"github.com/aphyx-framework/framework/framework/logging"
	"golang.org/x/mod/modfile"
	"os"
)

type Modfile struct {
	logger logging.ApplicationLogger
}

func (m Modfile) GetModuleName() string {
	goModBytes, err := os.ReadFile("./go.mod")
	if err != nil {
		m.logger.ErrorLogger.Fatalln("Failed to read go.mod file", err)
	}
	modName := modfile.ModulePath(goModBytes)

	return modName
}
