{{- template "disclaimer.go.edit" }}
package server

import (
    {{ template "import.statements" .Spec.Imports }}
)

type {{ .Spec.Name }}Controller struct {
    // Add fields as necessary
}

func New{{ .Spec.Name }}Controller() *{{ .Spec.Name }}Controller {
    return &{{ .Spec.Name }}Controller{
        // Initialize fields as necessary
    }
}

func (c *{{ .Spec.Name }}Controller) Startup() error {
    // Implement startup logic here
    return nil
}

func (c *{{ .Spec.Name }}Controller) Shutdown() error {
    // Implement shutdown logic here
    return nil
}

{{- range .Spec.Methods }}

func (c *{{ $.Spec.Name }}Controller) {{ .Name }}({{ template "method.parameters" .Parameters }}) {{ template "method.returns" .Returns }} {
    panic("not implemented")
}
{{- end }}