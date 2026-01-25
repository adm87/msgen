package msgen

import (
	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/utils"
)

// Create initializes a new Go module (if needed) and generates the microservice code.
func Create(msgenConfig *models.MSGenConfig, moduleUrl, specFile, outDir string) error {
	isModule, err := utils.IsGoModule(outDir)

	if err != nil {
		return err
	}

	if !isModule {
		msgenConfig.ModuleUrl = moduleUrl

		if err := utils.InitGoModule(moduleUrl, outDir); err != nil {
			return err
		}

		if err := utils.GoGet(outDir, "github.com/go-chi/chi/v5"); err != nil {
			return err
		}
	}

	return Generate(msgenConfig, specFile, outDir)
}
