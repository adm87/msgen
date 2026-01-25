{{- template "disclaimer.go.edit" }}
package server

import (
    {{ template "import.statements" .ServiceInfo.Imports }}
)

type {{ .ServiceInfo.Name }}Controller struct {
    // Add fields as necessary
}

func New{{ .ServiceInfo.Name }}Controller() *{{ .ServiceInfo.Name }}Controller {
    return &{{ .ServiceInfo.Name }}Controller{
        // Initialize fields as necessary
    }
}

func (c *{{ .ServiceInfo.Name }}Controller) Startup() error {
    // Implement startup logic here
    return nil
}

func (c *{{ .ServiceInfo.Name }}Controller) Shutdown() error {
    // Implement shutdown logic here
    return nil
}

{{- range .ServiceInfo.Methods }}

func (c *{{ $.ServiceInfo.Name }}Controller) {{ .Name }}({{ template "method.parameters" .Parameters }}) {{ template "method.returns" .Returns }} {
    panic("not implemented")
}
{{- end }}