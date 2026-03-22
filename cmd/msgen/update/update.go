package update

import (
	"github.com/adm87/msgen/models"
	"github.com/spf13/cobra"
)

func Command(msgArgs *models.MSGenArgs, cfg *models.MSGenCfg) *cobra.Command {
	c := &cobra.Command{
		Use: "update",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}
	return c
}
