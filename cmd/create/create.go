package create

import (
	"github.com/adm87/msgen/models"
	"github.com/spf13/cobra"
)

func Command(msgenConfig *models.MSGenConfig, msgenArgs *models.MSGenArgs) *cobra.Command {
	c := &cobra.Command{
		Use:   "create",
		Short: "Create a new microservice project",
		Long:  `This command creates a new microservice project based on the provided service specification interface.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	c.MarkFlagRequired(msgenArgs.SpecFile.Name)
	return c
}
