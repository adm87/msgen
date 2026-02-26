package args

type Arg struct {
	Name        string
	Description string
	Short       string
}

type StringArg struct {
	Arg
	Value string
}

type BoolArg struct {
	Arg
	Value bool
}

type IntArg struct {
	Arg
	Value int
}

type StringSliceArg struct {
	Arg
	Value []string
}

type IntSliceArg struct {
	Arg
	Value []int
}

type BoolSliceArg struct {
	Arg
	Value []bool
}
