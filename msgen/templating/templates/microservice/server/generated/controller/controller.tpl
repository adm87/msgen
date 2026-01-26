{{- template "disclaimer.go.noedit" }}
package controller

import (
    "{{ .MSGenConfig.ModuleUrl }}/server/generated/ctx"
    {{ template "import.statements" .Spec.Imports }}
)

type {{ .Spec.Name }}Controller interface {
{{- range .Spec.Methods }}
    {{ .Name }}(c *ctx.Context, {{ template "method.parameters" .Parameters }}) {{ template "method.returns" .Returns }}
{{- end }}
}