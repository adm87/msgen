{{- template "disclaimer.go.noedit" }}
package generated

{{- $spec := printf "%s.%s" .Spec.Package .Spec.Name }}

import (
    "{{ .MSGenConfig.ModuleUrl }}{{ .Spec.Path }}"
)

// ServerController defines a controller to handle server operations.
//
// It must implement the methods defined in the service specification.
type ServerController interface {
    {{ $spec }}

    Startup() error
    Shutdown() error
}

// Server represents the microservice server.
type Server struct {
    controller ServerController
}

func NewServer(controller ServerController) *Server {
    return &Server{
        controller: controller,
    }
}

func (s *Server) Start() error {
    return nil
}