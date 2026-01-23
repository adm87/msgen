package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/msgen"
	"github.com/adm87/msgen/utils"
	"github.com/spf13/cobra"
)

var (
	ErrMissingSpecFile = errors.New("service specification interface must be provided")
)

// MSGen creates the msgen command, validates flags and runs the generator
func MSGen(version string) (*cobra.Command, error) {
	config := models.DefaultMSGenConfig()

	var (
		workingDir string
		outputDir  string
		specFile   string
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

			configFile := filepath.Join(wrkDir, models.MSConfigFileName)

			if utils.FileExists(configFile) {
				loadedConfig, err := loadMSConfig(configFile)

				if err != nil {
					return err
				}

				config = loadedConfig
			}

			specFile = filepath.Join(wrkDir, specFile)

			if !utils.FileExists(specFile) {
				return ErrMissingSpecFile
			}

			return msgen.Generate(config, specFile, outDir)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	c.PersistentFlags().StringVarP(&workingDir, "workdir", "w", ".", "Working directory for the microservice")
	c.PersistentFlags().StringVarP(&outputDir, "output", "o", "./output", "Output directory for the generated code")
	c.PersistentFlags().StringVarP(&specFile, "spec", "s", "", "Path to the specification file relative to the working directory")

	if err := c.MarkPersistentFlagRequired("spec"); err != nil {
		return nil, err
	}

	return c, nil
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
