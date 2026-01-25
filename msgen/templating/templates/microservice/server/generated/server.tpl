{{- template "disclaimer.go.noedit" }}
package generated

{{- $spec := printf "%s.%s" .Spec.Package .Spec.Name }}

import (
	"net/http"

    "{{ .MSGenConfig.ModuleUrl }}{{ .Spec.Path }}"
	"github.com/go-chi/chi/v5"
)

// =================================================================
// Server implementation
// =================================================================

// Server represents the microservice server.
type Server struct {
    router chi.Router
    controller {{ $spec }}
}

// NewServer creates a new Server instance with the provided controller.
func NewServer(controller {{ $spec }}) *Server {
    return &Server{
        controller: controller,
    }
}

// Start setups the server routing, middleware, and starts the controller.
func (s *Server) Start() error {
    s.router = chi.NewRouter()

    if err := configureMiddleware(s); err != nil {
        return err
    }

    if err := configureRoutes(s); err != nil {
        return err
    }

    if err := startController(s); err != nil {
        return err
    }
    
    return http.ListenAndServe(":8080", s.router)
}

// Stop shuts down the server and stops the controller.
func (s *Server) Stop() error {
    return stopController(s)
}

func startController(s *Server) error {
    if controller, ok := s.controller.(OnStart); ok {
        if err := controller.Start(); err != nil {
            return err
        }
    }

    return nil
}

func stopController(s *Server) error {
    if controller, ok := s.controller.(OnStop); ok {
        if err := controller.Stop(); err != nil {
            return err
        }
    }

    return nil
}

func configureMiddleware(s *Server) error {
    s.router = chi.NewRouter()

    if controller, ok := s.controller.(OnConfigureMiddleware); ok {
        controller.ConfigureMiddleware(s.router)
    }

    return nil
}

func configureRoutes(s *Server) error {
    r := s.router
    {{- $routes := buildRoutingTree .Spec.Methods -}}
    {{- template "configure.routes" $routes }}
    return nil
}