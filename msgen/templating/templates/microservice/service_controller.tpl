{{- template "disclaimer.go" }}
package server

import (
    {{ template "import.statements" .Imports }}
)

// {{ .Name }}Controller handles requests for the {{ .Name }} service.
type {{ .Name }}Controller struct {
}

func New{{ .Name }}Controller() *{{ .Name }}Controller {
    return &{{ .Name }}Controller{}
}

{{- range .Methods }}

func (c *{{ $.Name }}Controller) {{ .Name }}({{ template "method.signature.parameters" .Parameters }}) {{ template "method.signature.returns" .Returns }} {
    panic("not implemented")
}
{{- end }}