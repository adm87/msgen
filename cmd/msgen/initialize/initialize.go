package initialize

import (
	"fmt"

	"github.com/adm87/msgen/app"
	"github.com/adm87/msgen/cli"
	"github.com/adm87/msgen/models"
	"github.com/adm87/msgen/tasks"
	"github.com/spf13/cobra"
)

func Command(msgArgs *models.MSGenArgs, cfg *models.MSGenCfg) *cobra.Command {
	args := &models.InitializeArgs{
		MSGenArgs: msgArgs,
	}

	c := &cobra.Command{
		Use: "init",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if cli.IsModule(msgArgs.OutputPath) {
				return fmt.Errorf("Go module already exists at output path: %s", msgArgs.OutputPath)
			}
			return app.Run(cmd.Context(), []app.Task{
				tasks.InitModule(msgArgs.OutputPath, args.Module),
				tasks.ModTidy(msgArgs.OutputPath),
				tasks.GenerateSpec(msgArgs.OutputPath),
			})
		},
	}

	c.Flags().StringVarP(&args.Module, "module", "m", args.Module, "name of the Go module to initialize (e.g. github.com/username/repo)")
	c.MarkFlagRequired("module")

	return c
}
