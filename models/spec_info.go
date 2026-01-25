package models

const (
	SummaryAttr     = "Summary"
	DescriptionAttr = "Description"
	VersionAttr     = "Version"
	AcceptsAttr     = "Accepts"
	ReturnsAttr     = "Returns"
	SuccessAttr     = "Success"
	FailureAttr     = "Failure"
	TagsAttr        = "Tags"
	ParamAttr       = "Param"
	RouterAttr      = "Router"
)

// Spec holds analyzed information about the service specification.
type Spec struct {
	Name       string          // Service name; name of the interface, e.g., "ExampleService"
	Package    string          // Package name of the spec file
	Path       string          // Path to the spec file
	Attributes *SpecAttributes // Additional attributes for the spec
	Imports    []SpecImport    // Slice of imports used in the spec file
	Methods    []SpecMethod    // Slice of methods defined in the service specification
}

// SpecImport represents an import statement in the service specification.
type SpecImport struct {
	Path  string // Import path
	Alias string // Optional alias for the import
}

// SpecMethod represents a method defined in the service specification.
type SpecMethod struct {
	Name       string                // Method name
	Attributes *SpecMethodAttributes // Attributes associated with the method
	Parameters []SpecField           // Slice of parameters for the method
	Returns    []SpecField           // Slice of return types for the method
}

// SpecField represents a parameter or return type in a method signature.
type SpecField struct {
	Name string // Parameter name
	Type string // Parameter type
}

// SpecAttributes holds additional attributes for the service specification.
type SpecAttributes struct {
	Summary     string // Summary of the method
	Description string // Description of the method
}

// SpecMethodAttributes holds additional attributes for a method in the service specification.
type SpecMethodAttributes struct {
	Summary     string                        // Summary of the method
	Description string                        // Description of the method
	Version     string                        // Version of the method
	Accepts     string                        // Accepted content type
	Returns     string                        // Returned content type
	Success     SpecMethodResponseAttribute   // Attributes of the success response
	Router      SpecMethodRouterAttribute     // Router attributes
	Tags        []string                      // Tags associated with the method
	Failures    []SpecMethodResponseAttribute // Attributes of the failure responses
	Parameters  []SpecMethodParamAttribute    // Attributes of the method parameters
}

// SpecMethodParamAttribute represents attributes of a method parameter.
type SpecMethodParamAttribute struct {
	Name     string // Name of the parameter
	Type     string // Type of the parameter
	In       string // Location of the parameter (path, query, header, body)
	Required bool   // Whether the parameter is required
}

// SpecMethodResponseAttribute represents attributes of a method response.
type SpecMethodResponseAttribute struct {
	Code     string // HTTP status code for success or failure
	Type     string // Type of the response (object, array, etc.)
	DataType string // Data type of the response
}

// SpecMethodRouterAttribute represents routing attributes of a method.
type SpecMethodRouterAttribute struct {
	Path   string // Route path
	Method string // HTTP method (GET, POST, etc.)
}

// SpecRoutingNode represents a node in the routing tree.
type SpecRoutingNode struct {
	Methods  []SpecMethod                // Methods associated with this node
	Children map[string]*SpecRoutingNode // Child nodes in the routing tree
}
