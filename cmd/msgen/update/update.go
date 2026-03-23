package update

import (
	"fmt"

	"github.com/adm87/msgen/app"
	"github.com/adm87/msgen/cli"
	"github.com/adm87/msgen/models"
	"github.com/spf13/cobra"
)

func Command(msgArgs *models.MSGenArgs, cfg *models.MSGenCfg) *cobra.Command {
	args := &models.UpdateArgs{
		MSGenArgs: msgArgs,
	}

	c := &cobra.Command{
		Use: "update",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !cli.IsModule(msgArgs.OutputPath) {
				return fmt.Errorf("no Go module found at output path: %s", msgArgs.OutputPath)
			}
			return app.Run(cmd.Context(), []app.Task{})
		},
	}

	c.Flags().StringVarP(&args.SpecFile, "spec", "s", args.SpecFile, "path to OpenAPI spec file for update")
	c.MarkFlagRequired("spec")

	return c
}
