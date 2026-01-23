package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/msgen"
	"github.com/adm87/msgen/utils"
	"github.com/spf13/cobra"
)

func MSGen(version string) *cobra.Command {
	config := models.DefaultMSGenConfig()

	var (
		workingDir string
		outputDir  string
	)

	c := &cobra.Command{
		Use:     "msgen",
		Short:   "Microservice code generator",
		Long:    `MSGen is a tool to generate boilerplate code for microservices.`,
		Version: version,
		RunE: func(cmd *cobra.Command, args []string) error {
			wrkDir, err := filepath.Abs(workingDir)

			if err != nil {
				return err
			}

			outDir, err := filepath.Abs(outputDir)

			if err != nil {
				return err
			}

			if utils.FileExists(wrkDir) {
				loadedConfig, err := loadMSConfig(filepath.Join(wrkDir, models.MSConfigFileName))

				if err != nil {
					return err
				}

				config = loadedConfig
			}

			return msgen.Generate(config, outDir)
		},
	}

	c.PersistentFlags().StringVarP(&workingDir, "workdir", "w", ".", "Working directory for the microservice")
	c.PersistentFlags().StringVarP(&outputDir, "output", "o", "./output", "Output directory for the generated code")

	return c
}

func loadMSConfig(path string) (models.MSGenConfig, error) {
	var config models.MSGenConfig

	file, err := os.Open(path)

	if err != nil {
		return config, err
	}

	defer file.Close()

	err = json.NewDecoder(file).Decode(&config)

	return config, err
}
