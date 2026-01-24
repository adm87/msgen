package cmd

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"

	"github.com/adm87/msgen/cmd/create"
	"github.com/adm87/msgen/cmd/update"
	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/utils"
	"github.com/spf13/cobra"
)

var (
	ErrMissingSpecFile = errors.New("service specification interface must be provided")
)

// MSGen creates the msgen command, validates flags and runs the generator
func MSGen(version string) (*cobra.Command, error) {
	config := models.DefaultMSGenConfig()
	cmdArgs := models.MSGenArgs{}

	c := &cobra.Command{
		Use:     "msgen",
		Short:   "Microservice code generator",
		Long:    `MSGen is a tool to generate boilerplate code for microservices.`,
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			wrkDir, err := filepath.Abs(cmdArgs.WorkingDir)
			if err != nil {
				return err
			}
			cmdArgs.WorkingDir = wrkDir

			outDir, err := filepath.Abs(cmdArgs.OutDir)
			if err != nil {
				return err
			}
			cmdArgs.OutDir = outDir

			configFile := filepath.Join(wrkDir, models.MSConfigFileName)
			if utils.FileExists(configFile) {
				loadedConfig, err := loadMSConfig(configFile)

				if err != nil {
					return err
				}

				config = loadedConfig
			}

			cmdArgs.SpecFile = filepath.Join(wrkDir, cmdArgs.SpecFile)
			if !utils.FileExists(cmdArgs.SpecFile) {
				return ErrMissingSpecFile
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	c.AddCommand(
		create.Command(&config, &cmdArgs),
		update.Command(&config, &cmdArgs),
	)

	c.PersistentFlags().StringVarP(&cmdArgs.WorkingDir, "workdir", "w", ".", "Working directory for the microservice")
	c.PersistentFlags().StringVarP(&cmdArgs.OutDir, "output", "o", "./output", "Output directory for the generated code")
	c.PersistentFlags().StringVarP(&cmdArgs.SpecFile, "spec", "s", "", "Path to the specification file relative to the working directory")

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
