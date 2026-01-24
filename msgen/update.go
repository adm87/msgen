package msgen

import "github.com/adm87/msgen/models"

func Update(msgenConfig *models.MSGenConfig, specFile string, outDir string) error {
	// TODO - add any pre-generation steps here
	return Generate(msgenConfig, specFile, outDir)
}
