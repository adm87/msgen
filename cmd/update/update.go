package update

import (
	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/msgen"
	"github.com/spf13/cobra"
)

func Command(msgenConfig *models.MSGenConfig, msgenArgs *models.MSGenArgs) *cobra.Command {
	return &cobra.Command{
		Use:   "update",
		Short: "Update an existing microservice project",
		Long:  `This command updates an existing microservice project based on changes in the service specification interface.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return msgen.Generate(msgenConfig, msgenArgs.SpecFile.Value, msgenArgs.WorkingDir.Value)
		},
	}
}
