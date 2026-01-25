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
		msgenConfig.ModuleUrl = moduleUrl

		if err := initGoModule(moduleUrl, outDir); err != nil {
			return err
		}

		if err := goGetChi(outDir); err != nil {
			return err
		}
	} else {
		modUrl, err := getModuleUrl(outDir)

		if err != nil {
			return err
		}

		if msgenConfig.ModuleUrl != modUrl {
			log.Warn("Module URL mismatch in config and mod file. Using mod file:", "configUrl", msgenConfig.ModuleUrl, "modUrl", modUrl)

			msgenConfig.ModuleUrl = modUrl
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

func getModuleUrl(outDir string) (string, error) {
	cmd := exec.Command("go", "list", "-m")

	cmd.Dir = outDir

	output, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}
