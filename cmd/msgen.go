package cmd

import (
	"github.com/adm87/msgen/app"
	"github.com/adm87/msgen/data"
	"github.com/adm87/msgen/data/args"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func MSGen() error {
	a := data.NewMSGenArgs()

	c := &cobra.Command{
		Use:   "msgen",
		Short: "A tool for generating code from OpenAPI specs",
		Long:  `MSGen is a tool for generating code from OpenAPI specs. It can be used to generate code for various programming languages and frameworks, and it can be used to create new files or update existing ones.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			m, err := app.NewModel(a)

			if err != nil {
				return err
			}

			if _, err := tea.NewProgram(m).Run(); err != nil {
				return err
			}

			return m.Error()
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	AddStringArg(c, &a.WorkingDir)
	AddStringArg(c, &a.SpecFile)
	AddBoolArg(c, &a.NonInteractive)
	AddBoolArg(c, &a.Create)
	AddBoolArg(c, &a.Update)

	if err := c.MarkFlagRequired(a.SpecFile.Name); err != nil {
		return err
	}

	c.MarkFlagsOneRequired(
		a.Create.Name,
		a.Update.Name,
	)
	c.MarkFlagsMutuallyExclusive(
		a.Create.Name,
		a.Update.Name,
	)

	return c.Execute()
}

func AddStringArg(c *cobra.Command, arg *args.StringArg) {
	c.Flags().StringVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
}

func AddBoolArg(c *cobra.Command, arg *args.BoolArg) {
	c.Flags().BoolVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
}

func AddIntArg(c *cobra.Command, arg *args.IntArg) {
	c.Flags().IntVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
}

func AddStringSliceArg(c *cobra.Command, arg *args.StringSliceArg) {
	c.Flags().StringSliceVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
}

func AddIntSliceArg(c *cobra.Command, arg *args.IntSliceArg) {
	c.Flags().IntSliceVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
}

func AddBoolSliceArg(c *cobra.Command, arg *args.BoolSliceArg) {
	c.Flags().BoolSliceVarP(&arg.Value, arg.Name, arg.Short, arg.Value, arg.Description)
}
