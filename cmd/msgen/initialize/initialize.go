package initialize

import (
	"github.com/adm87/msgen/models"
	"github.com/spf13/cobra"
)

func Command(msgArgs *models.MSGenArgs, cfg *models.MSGenCfg) *cobra.Command {
	c := &cobra.Command{
		Use: "init",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}
	return c
}
