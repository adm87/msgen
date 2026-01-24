package msgen

import "github.com/adm87/msgen/models"

// Update regenerates the microservice code based on the provided specification.
func Update(msgenConfig *models.MSGenConfig, specFile string, outDir string) error {
	// TODO - add any pre-generation steps here
	return Generate(msgenConfig, specFile, outDir)
}
