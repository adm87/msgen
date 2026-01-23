package cmd

import (
	"github.com/spf13/cobra"
)

func MSGen(version string) *cobra.Command {
	c := &cobra.Command{
		Use:     "msgen",
		Short:   "Microservice code generator",
		Long:    `MSGen is a tool to generate boilerplate code for microservices.`,
		Version: version,
	}
	return c
}
