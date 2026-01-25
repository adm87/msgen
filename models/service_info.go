package models

// Spec holds analyzed information about the service specification.
type Spec struct {
	Name    string       // Service name; name of the interface, e.g., "ExampleService"
	Package string       // Package name of the spec file
	Path    string       // Path to the spec file
	Imports []SpecImport // Slice of imports used in the spec file
	Methods []SpecMethod // Slice of methods defined in the service specification
}

// SpecImport represents an import statement in the service specification.
type SpecImport struct {
	Path  string // Import path
	Alias string // Optional alias for the import
}

// SpecMethod represents a method defined in the service specification.
type SpecMethod struct {
	Name       string      // Method name
	Parameters []SpecField // Slice of parameters for the method
	Returns    []SpecField // Slice of return types for the method
}

// SpecField represents a parameter or return type in a method signature.
type SpecField struct {
	Name string // Parameter name
	Type string // Parameter type
}
