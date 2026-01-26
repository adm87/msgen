{{- template "disclaimer.go.noedit" }}
package generated

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

    "{{ .MSGenConfig.ModuleUrl }}/server/generated/controller"

	"github.com/go-chi/chi/v5"
)

// =================================================================
// Server implementation
// =================================================================

// Server represents the microservice server.
type Server struct {
    router chi.Router
    controller controller.{{ .Spec.Name }}Controller

    logger *slog.Logger
}

// NewServer creates a new Server instance with the provided controller.
func NewServer(controller controller.{{ .Spec.Name }}Controller) *Server {
    return &Server{
        controller: controller,
		logger: slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})).With("service", "{{ .Spec.Name }}"),
    }
}

// Start setups the server routing, middleware, and starts the controller.
func (s *Server) Start() error {
	s.router = chi.NewRouter()

	s.router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method %s not allowed\n", r.Method)
	})

	if controller, ok := s.controller.(OnConfigureLogger); ok {
		controller.ConfigureLogger(s.logger,)
	}

	if controller, ok := s.controller.(OnConfigureRouter); ok {
		controller.ConfigureRouter(s.router, s.logger)
	}

	configureRoutes(s)

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
		if err := controller.Start(s.logger); err != nil {
			return err
		}
	}

	return nil
}

func stopController(s *Server) error {
	if controller, ok := s.controller.(OnStop); ok {
		if err := controller.Stop(s.logger); err != nil {
			return err
		}
	}

	return nil
}

func configureRoutes(s *Server) {
    {{- $routes := buildRoutingTree .Spec.Methods -}}
    {{- template "configure.routes" dict "routes" $routes "router" "s.router" "server" "s" "path" "" -}}
}

func configureScopedRouter(r chi.Router, s *Server, path string) {
	if controller, ok := s.controller.(OnConfigureScopedRouter); ok {
		controller.ConfigureScopedRouter(r, path, s.logger)
	}
}
