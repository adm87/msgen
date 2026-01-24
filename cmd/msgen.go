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
	cmdArgs := models.DefaultMSGenArgs()

	c := &cobra.Command{
		Use:     "msgen",
		Short:   "Microservice code generator",
		Long:    `MSGen is a tool to generate boilerplate code for microservices.`,
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			var err error

			cmdArgs.WorkingDir.Value, err = filepath.Abs(cmdArgs.WorkingDir.Value)

			if err != nil {
				return err
			}

			cmdArgs.OutDir.Value, err = filepath.Abs(cmdArgs.OutDir.Value)

			if err != nil {
				return err
			}

			configFile := filepath.Join(cmdArgs.WorkingDir.Value, models.MSConfigFileName)
			fileExists, err := utils.FileExists(configFile)

			if err != nil {
				return err
			}

			if fileExists {
				loadedConfig, err := loadMSConfig(configFile)
				if err != nil {
					return err
				}
				config = loadedConfig
			}

			cmdArgs.SpecFile.Value = filepath.Join(cmdArgs.WorkingDir.Value, cmdArgs.SpecFile.Value)
			fileExists, err = utils.FileExists(cmdArgs.SpecFile.Value)

			if !fileExists {
				return errors.Join(ErrMissingSpecFile, err)
			}

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	utils.AddPersistentStringFlags(c, &cmdArgs.WorkingDir)

	createCmd := create.Command(&config, &cmdArgs)
	utils.AddStringFlag(createCmd, &cmdArgs.SpecFile)
	utils.AddStringFlag(createCmd, &cmdArgs.OutDir)
	createCmd.MarkFlagRequired(cmdArgs.SpecFile.Name)

	updateCmd := update.Command(&config, &cmdArgs)
	utils.AddStringFlag(updateCmd, &cmdArgs.SpecFile)
	utils.AddStringFlag(updateCmd, &cmdArgs.OutDir)
	updateCmd.MarkFlagRequired(cmdArgs.SpecFile.Name)

	c.AddCommand(createCmd, updateCmd)

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
