package create

import (
	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/msgen"
	"github.com/spf13/cobra"
)

func Command(msgenConfig *models.MSGenConfig, msgenArgs *models.MSGenArgs) *cobra.Command {
	c := &cobra.Command{
		Use:   "create",
		Short: "Initialize a new microservice project",
		Long:  `This command initializes a new microservice project with the necessary configuration files.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return msgen.Create(msgenConfig, msgenArgs.SpecFile, msgenArgs.OutDir)
		},
	}
	return c
}
