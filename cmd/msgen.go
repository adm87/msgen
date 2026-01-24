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
			wrkDir, err := filepath.Abs(cmdArgs.WorkingDir.Value)
			if err != nil {
				return err
			}
			cmdArgs.WorkingDir.Value = wrkDir

			outDir, err := filepath.Abs(cmdArgs.OutDir.Value)
			if err != nil {
				return err
			}
			cmdArgs.OutDir.Value = outDir

			configFile := filepath.Join(wrkDir, models.MSConfigFileName)
			if utils.FileExists(configFile) {
				loadedConfig, err := loadMSConfig(configFile)

				if err != nil {
					return err
				}

				config = loadedConfig
			}

			cmdArgs.SpecFile.Value = filepath.Join(wrkDir, cmdArgs.SpecFile.Value)
			if !utils.FileExists(cmdArgs.SpecFile.Value) {
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

	addPersistentStringFlags(c, &cmdArgs.WorkingDir)

	createCmd := create.Command(&config, &cmdArgs)
	addStringFlag(createCmd, &cmdArgs.SpecFile)
	addStringFlag(createCmd, &cmdArgs.OutDir)
	createCmd.MarkFlagRequired(cmdArgs.SpecFile.Name)

	updateCmd := update.Command(&config, &cmdArgs)
	addStringFlag(updateCmd, &cmdArgs.SpecFile)
	addStringFlag(updateCmd, &cmdArgs.OutDir)
	updateCmd.MarkFlagRequired(cmdArgs.SpecFile.Name)

	c.AddCommand(createCmd, updateCmd)

	return c, nil
}

func addPersistentStringFlags(cmd *cobra.Command, arg *models.StringArg) {
	cmd.PersistentFlags().StringVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
}

func addStringFlag(cmd *cobra.Command, arg *models.StringArg) {
	cmd.Flags().StringVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
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
