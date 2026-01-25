package msgen

import (
	"encoding/json"
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
	SpecPath    string
}

// Generate generates the microservice code based on the provided specification.
func Generate(msgenConfig *models.MSGenConfig, specFile string, workingDir string) error {
	if err := validateModUrl(msgenConfig, workingDir); err != nil {
		return err
	}

	if err := removeGeneratedDirectories(workingDir); err != nil {
		return err
	}

	serviceInfo, err := analyzer.AnalyzeSpec(filepath.Join(workingDir, specFile))

	if err != nil {
		return err
	}

	ctx := &Context{
		MSGenConfig: msgenConfig,
		ServiceInfo: serviceInfo,
		SpecPath:    filepath.Dir(specFile),
	}

	if err := generateFiles(ctx, templating.MicroserviceTemplateTree, workingDir); err != nil {
		return err
	}

	if err := writeMSGenConfig(ctx, workingDir); err != nil {
		return err
	}

	if err := utils.GoFmt(workingDir); err != nil {
		return err
	}

	if err := utils.ModTidy(workingDir); err != nil {
		return err
	}

	return nil
}

func validateModUrl(msgenConfig *models.MSGenConfig, path string) error {
	modUrl, err := utils.GetModuleUrl(path)

	if err != nil {
		return err
	}

	if msgenConfig.ModuleUrl != modUrl {
		log.Warn("Module URL mismatch in config and mod file. Using mod file:", "configUrl", msgenConfig.ModuleUrl, "modUrl", modUrl)

		msgenConfig.ModuleUrl = modUrl
	}

	return nil
}

func removeGeneratedDirectories(path string) error {
	return filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
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

			continue
		}

		if child.GenerateOnce {
			exists, err := utils.FileExists(childPath)

			if err != nil {
				return err
			}

			if exists {
				continue
			}
		}

		log.Info("Generating file:", "path", childPath)

		if err := templating.RenderTemplateToFile(child.Template, childPath, ctx); err != nil {
			return err
		}
	}

	return nil
}

func writeMSGenConfig(ctx *Context, path string) error {
	configPath := filepath.Join(path, models.MSConfigFileName)

	log.Info("Writing msgen config file:", "path", configPath)

	data, err := json.MarshalIndent(ctx.MSGenConfig, "", "  ")
	if err != nil {
		return err
	}

	return utils.WriteFile(configPath, string(data))
}
