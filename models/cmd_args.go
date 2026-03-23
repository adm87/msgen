package models

// ------------------------------------------------------------
// MSGen global commandline arguments
// ------------------------------------------------------------

type MSGenArgs struct {
	MSGenYML   string
	OutputPath string
	AutoAccept bool
}

// ------------------------------------------------------------
// MSGen initialization commandline arguments
// ------------------------------------------------------------

type InitializeArgs struct {
	*MSGenArgs

	Module string
}

// ------------------------------------------------------------
// MSGen update commandline arguments
// ------------------------------------------------------------

type UpdateArgs struct {
	*MSGenArgs

	SpecFile string
}
