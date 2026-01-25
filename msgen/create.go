package msgen

import (
	"os/exec"

	"github.com/adm87/msgen/log"
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
		if err := initGoModule(moduleUrl, outDir); err != nil {
			return err
		}

		if err := goGetChi(outDir); err != nil {
			return err
		}
	}

	return Generate(msgenConfig, specFile, outDir)
}

func initGoModule(moduleUrl, outDir string) error {
	log.Info("Initializing go module:", "moduleUrl", moduleUrl)

	cmd := exec.Command("go", "mod", "init", moduleUrl)
	cmd.Dir = outDir

	return cmd.Run()
}

func goGetChi(outDir string) error {
	log.Info("Downloading chi module")

	cmd := exec.Command("go", "get", "github.com/go-chi/chi/v5")
	cmd.Dir = outDir

	return cmd.Run()
}
