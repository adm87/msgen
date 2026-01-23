package msgen

import (
	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/msgen/analyzer"
)

func Generate(config models.MSGenConfig, specFile string, outDir string) error {
	serviceInfo, err := analyzer.AnalyzeSpec(specFile)

	if err != nil {
		return err
	}

	_ = serviceInfo

	return nil
}
