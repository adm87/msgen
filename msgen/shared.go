package msgen

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/adm87/msgen/log"
	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/msgen/analyzer"
	"github.com/adm87/msgen/msgen/templating"
	"github.com/adm87/msgen/utils"
)

// Context holds the context information needed during code generation.
type Context struct {
	MSGenConfig *models.MSGenConfig
	ServiceInfo *models.ServiceInfo
}

// Generate generates the microservice code based on the provided specification.
func Generate(msgenConfig *models.MSGenConfig, specFile string, outDir string) error {
	if err := removeGeneratedDirectories(outDir); err != nil {
		return err
	}

	serviceInfo, err := analyzer.AnalyzeSpec(specFile)

	if err != nil {
		return err
	}

	ctx := &Context{
		MSGenConfig: msgenConfig,
		ServiceInfo: serviceInfo,
	}

	if err := generateFiles(ctx, templating.MicroserviceTemplateTree, outDir); err != nil {
		return err
	}

	return nil
}

func removeGeneratedDirectories(outDir string) error {
	return filepath.Walk(outDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && info.Name() == "generated" {
			log.Info("Removing generated directory:", "path", path)

			if err := os.RemoveAll(path); err != nil {
				return err
			}

			return filepath.SkipDir
		}

		return nil
	})
}

func generateFiles(ctx *Context, node *templating.Node, currentPath string) error {
	for _, child := range node.Children {
		if shouldGenerate := child.ShouldGenerate; shouldGenerate != nil && !shouldGenerate() {
			log.Info("Skipping generation of:", "name", child.Name)

			continue
		}

		childPath := filepath.Join(currentPath, child.Name)

		if child.Template == "" {
			if err := utils.MakeDirectory(childPath); err != nil {
				return err
			}

			if err := generateFiles(ctx, child, childPath); err != nil {
				return err
			}

			return nil
		}

		log.Info("Generating file:", "path", childPath)

		if err := templating.RenderTemplateToFile(child.Template, childPath, ctx); err != nil {
			return err
		}
	}

	return nil
}
