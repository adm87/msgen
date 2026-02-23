{{- template "disclaimer.go.noedit" }}
package controller

import (
    {{- if gt (len .Spec.Imports) 0 }}
    "{{ .MSGenConfig.ModuleUrl }}/server/generated/ctx"
    {{- end }}
    {{ template "import.statements" .Spec.Imports }}
)

type {{ .Spec.Name }}Controller interface {
{{- range .Spec.Methods }}
    {{ .Name }}(c *ctx.Context, {{ template "method.parameters" .Parameters }}) {{ template "method.returns" .Returns }}
{{- end }}
}