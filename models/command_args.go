package models

type StringArg struct {
	Name        string
	Short       string
	Value       string
	Description string
}

type MSGenArgs struct {
	SpecFile   StringArg
	OutDir     StringArg
	WorkingDir StringArg
}

func DefaultMSGenArgs() MSGenArgs {
	return MSGenArgs{
		SpecFile:   StringArg{Name: "spec", Short: "s", Value: "", Description: "Path to the specification file relative to the working directory"},
		OutDir:     StringArg{Name: "output", Short: "o", Value: ".", Description: "Output directory"},
		WorkingDir: StringArg{Name: "workdir", Short: "w", Value: ".", Description: "msgen's working directory"},
	}
}

type CreateArgs struct {
	ModuleUrl StringArg
}

func DefaultCreateArgs() CreateArgs {
	return CreateArgs{
		ModuleUrl: StringArg{Name: "module", Short: "m", Value: "", Description: "Go module URL for the new microservice project"},
	}
}
