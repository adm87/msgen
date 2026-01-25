{{- template "disclaimer.go.noedit" }}
package generated

import "net/http"

{{- range $method := .Spec.Methods }}

// {{ .Name }} handles the {{ .Attributes.Router.Method }} {{ .Attributes.Router.Path }} endpoint.
func {{ .Name }}(w http.ResponseWriter, r *http.Request) {

}
{{- end }}