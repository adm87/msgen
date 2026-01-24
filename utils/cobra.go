package utils

import (
	"github.com/adm87/msgen/models"
	"github.com/spf13/cobra"
)

func AddPersistentStringFlags(cmd *cobra.Command, arg *models.StringArg) {
	cmd.PersistentFlags().StringVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
}

func AddStringFlag(cmd *cobra.Command, arg *models.StringArg) {
	cmd.Flags().StringVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
}
