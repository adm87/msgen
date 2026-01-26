package templating

// Node represents a file or directory in the template tree.
type Node struct {
	Name           string      // File or Directory name
	Template       string      // Template path to use to generate this file (empty for directories)
	Children       []*Node     // Child nodes (for directories)
	GenerateOnce   bool        // If true, the file will only be generated once
	ShouldGenerate func() bool // Function to determine if this node should be generated
}

// MicroserviceTemplateTree defines the directory and file structure for a microservice project
var MicroserviceTemplateTree = &Node{
	Name: "",
	Children: []*Node{
		{
			Name:         "main.go",
			Template:     "templates/microservice/main.tpl",
			GenerateOnce: true,
		},
		{
			Name: "server",
			Children: []*Node{
				{
					Name:         "controller.go",
					Template:     "templates/microservice/server/controller.tpl",
					GenerateOnce: true,
				},
				{
					Name: "generated",
					Children: []*Node{
						{
							Name: "controller",
							Children: []*Node{
								{
									Name:     "controller.go",
									Template: "templates/microservice/server/generated/controller/controller.tpl",
								},
							},
						},
						{
							Name: "ctx",
							Children: []*Node{
								{
									Name:     "context.go",
									Template: "templates/microservice/server/generated/ctx/context.tpl",
								},
							},
						},
						{
							Name: "web",
							Children: []*Node{
								{
									Name:     "responses.go",
									Template: "templates/microservice/server/generated/web/responses.tpl",
								},
								{
									Name:     "server_error.go",
									Template: "templates/microservice/server/generated/web/server_error.tpl",
								},
							},
						},
						{
							Name:     "handlers.go",
							Template: "templates/microservice/server/generated/handlers.tpl",
						},
						{
							Name:     "hooks.go",
							Template: "templates/microservice/server/generated/hooks.tpl",
						},
						{
							Name:     "server.go",
							Template: "templates/microservice/server/generated/server.tpl",
						},
					},
				},
			},
		},
	},
}
