package data

import "github.com/adm87/msgen/data/args"

type MSGenArgs struct {
	WorkingDir     args.StringArg
	SpecFile       args.StringArg
	NonInteractive args.BoolArg
	Create         args.BoolArg
	Update         args.BoolArg
}

func NewMSGenArgs() *MSGenArgs {
	return &MSGenArgs{
		WorkingDir: args.StringArg{
			Arg: args.Arg{Name: "working-dir", Description: "The working directory to use for the generated code", Short: "w"},
		},
		SpecFile: args.StringArg{
			Arg: args.Arg{Name: "spec-file", Description: "The OpenAPI spec file to use for generating the code", Short: "s"},
		},
		NonInteractive: args.BoolArg{
			Arg: args.Arg{Name: "non-interactive", Description: "Whether to run in non-interactive mode or not", Short: "n"},
		},
		Create: args.BoolArg{
			Arg: args.Arg{Name: "create", Description: "Whether to create new files or update existing ones", Short: "c"},
		},
		Update: args.BoolArg{
			Arg: args.Arg{Name: "update", Description: "Whether to update existing files or create new ones", Short: "u"},
		},
	}
}
