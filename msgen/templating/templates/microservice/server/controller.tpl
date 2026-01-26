{{- template "disclaimer.go.edit" }}
package server

import (

    "{{ .MSGenConfig.ModuleUrl }}/server/generated/ctx"
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

{{- range .Spec.Methods }}

func (c *{{ $.Spec.Name }}Controller) {{ .Name }}(ctx *ctx.Context, {{ template "method.parameters" .Parameters }}) {{ template "method.returns" .Returns }} {
    panic("not implemented")
}
{{- end }}