package msgen

import (
	"os"
	"path/filepath"

	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/msgen/analyzer"
	"github.com/adm87/msgen/msgen/templating"
)

func Generate(config models.MSGenConfig, specFile string, outDir string) error {
	serviceInfo, err := analyzer.AnalyzeSpec(specFile)

	if err != nil {
		return err
	}

	content, err := templating.RenderTemplate("test", "templates/microservice/service_controller.tpl", serviceInfo)

	if err != nil {
		return err
	}

	println(content)

	if err := os.MkdirAll(outDir, 0755); err != nil {
		return err
	}

	outputFile := filepath.Join(outDir, "service_controller.go")
	if err := os.WriteFile(outputFile, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}
