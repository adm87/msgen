package create

import (
	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/msgen"
	"github.com/adm87/msgen/utils"
	"github.com/spf13/cobra"
)

func Command(msgenConfig *models.MSGenConfig, msgenArgs *models.MSGenArgs) *cobra.Command {
	cmdArgs := models.DefaultCreateArgs()

	c := &cobra.Command{
		Use:   "create",
		Short: "Create a new microservice project",
		Long:  `This command creates a new microservice project based on the provided service specification interface.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := utils.ValidateModuleUrl(cmdArgs.ModuleUrl.Value); err != nil {
				return err
			}
			return msgen.Create(msgenConfig, cmdArgs.ModuleUrl.Value, msgenArgs.SpecFile.Value, msgenArgs.WorkingDir.Value)
		},
	}

	utils.AddStringFlag(c, &cmdArgs.ModuleUrl)
	c.MarkFlagRequired(cmdArgs.ModuleUrl.Name)

	return c
}
