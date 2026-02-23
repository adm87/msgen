{{- template "disclaimer.go.noedit" }}
package generated

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

    {{ template "import.statements" .Spec.Imports }}
    {{- if gt (len .Spec.Imports) 0 }}
    "{{ .MSGenConfig.ModuleUrl }}/server/generated/ctx"
    {{- end }}
    "{{ .MSGenConfig.ModuleUrl }}/server/generated/web"

	"github.com/go-chi/chi/v5"
)

func getPathParam(r *http.Request, name string, required bool) (string, *web.ServerError) {
    value := chi.URLParam(r, name)
    if required && value == "" {
        return "", web.NewServerError(http.StatusBadRequest, "missing required path parameter: "+name)
    }
    return value, nil
}

func getQueryParam(r *http.Request, name string, required bool) (string, *web.ServerError) {
    value := r.URL.Query().Get(name)
    if required && value == "" {
        return "", web.NewServerError(http.StatusBadRequest, "missing required query parameter: "+name)
    }
    return value, nil
}

func getBodyParam[T any](r *http.Request, required bool) (T, *web.ServerError) {
	var param T

	if r.Body == nil {
		if required {
			return param, web.NewServerError(http.StatusBadRequest, "missing request body")
		}
		return param, nil
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&param); err != nil {
		if errors.Is(err, io.EOF) {
			if required {
				return param, web.NewServerError(http.StatusBadRequest, "missing request body")
			}
			return param, nil
		}

		return param, web.NewServerError(http.StatusBadRequest, "invalid request body: "+err.Error())
	}

	return param, nil
}

{{- range $method := .Spec.Methods }}

// {{ .Name }} handles the {{ .Attributes.Router.Method }} {{ .Attributes.Router.Path }} endpoint.
func {{ .Name }}(s *Server) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        logger := s.logger.With("handler", "{{ .Name }}", "method", r.Method, "path", "{{ .Attributes.Router.Path }}")
        context := ctx.NewContext(logger, r)

        {{ range $param := $method.Attributes.Parameters }}
        {{- if $param.FromPath }}
        {{ $param.Name }}, err := getPathParam(r, "{{ $param.Name }}", true)
        if err != nil {
            logger.Error("failed to get path parameter", "param", "{{ $param.Name }}", "error", err)
            web.WriteErrorResponse(w, err)
            return
        }
        {{- end }}
        {{- if $param.FromQuery }}
        {{ $param.Name }}, err := getQueryParam(r, "{{ $param.Name }}", false)
        if err != nil {
            logger.Error("failed to get query parameter", "param", "{{ $param.Name }}", "error", err)
            web.WriteErrorResponse(w, err)
            return
        }
        {{- end }}
        {{- if $param.FromBody }}
        {{ $param.Name }}, err := getBodyParam[{{ $param.Type }}](r, true)
        if err != nil {
            logger.Error("failed to get request body", "param", "{{ $param.Name }}", "error", err)
            web.WriteErrorResponse(w, err)
            return
        }
        {{- end }}
        {{ end }}

        {{- if gt (len .Returns) 1 }}
        result, handlerErr := s.controller.{{ .Name }}(context{{- range .Parameters }}, {{ .Name }}{{- end }})
        {{- else }}
        handlerErr := s.controller.{{ .Name }}(context{{- range .Parameters }}, {{ .Name }}{{- end }})
        {{- end }}

        if handlerErr != nil {
            logger.Error(handlerErr.Error())
            if serverErr, ok := handlerErr.(*web.ServerError); ok {                
                web.WriteErrorResponse(w, serverErr)
            } else {
                web.WriteErrorResponse(w, web.NewServerError(http.StatusInternalServerError, handlerErr.Error()))
            }
            return
        }

        {{- if gt (len .Returns) 1 }}
        web.WriteJSONResponse(w, {{ $method.Attributes.Success.Code }}, result)
        {{- else }}
        w.WriteHeader({{ $method.Attributes.Success.Code }})
        {{- end }}
    }
}
{{- end }}
