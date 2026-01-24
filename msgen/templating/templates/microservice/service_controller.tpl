{{- template "disclaimer.go" }}
package server

import (
    {{- include "import.statements" .Imports -}}
)

// {{ .Name }}Controller handles requests for the {{ .Name }} service.
type {{ .Name }}Controller struct {
}

func New{{ .Name }}Controller() *{{ .Name }}Controller {
    return &{{ .Name }}Controller{}
}

{{- range .Methods }}

func (c *{{ $.Name }}Controller) {{ .Name }}({{- include "method.signature.parameters" .Parameters -}}) {{- include "method.signature.returns" .Returns -}} {
    panic("not implemented")
}
{{- end }}