# MSGen - Microservice Generator (Work in Progress)

MSGen is a code generation tool for building Go microservices from interface specifications. Define your service API as a Go interface with annotations, and MSGen generates a complete HTTP microservice with routing, handlers, parameter validation, and error handling.

## Features

- 🚀 **Interface-Driven Development** - Define your API as annotated Go interfaces
- 🔄 **Full HTTP Server Generation** - Complete chi-router based server with handlers
- ✅ **Parameter Validation** - Automatic path, query, and body parameter extraction and validation
- 📝 **Structured Logging** - Built-in structured logging with slog
- 🎯 **Type-Safe** - Leverages Go's type system and AST parsing
- 🔧 **Lifecycle Hooks** - Optional hooks for startup, shutdown, and configuration
- ♻️ **Regeneration Safe** - User code protected from regeneration

## Installation

```bash
go install github.com/adm87/msgen@latest
```

Or clone and build from source:

```bash
git clone https://github.com/adm87/msgen.git
cd msgen
go build -o msgen
```

## Quick Start

### 1. Create a Service Specification

Create a directory for your project and define your service interface:

```bash
mkdir my-service
cd my-service
mkdir spec
```

Create `spec/spec.go`:

```go
package spec

import "my-service/models"

// @Summary User Service
// @Description Service for managing user profiles
type UserService interface {
    // @Summary Get user by ID
    // @Description Retrieves user profile by ID
    // @Version 1
    // @Param userID string path true
    // @Success 200 {object} models.User
    // @Failure 404 {object} models.ErrorResponse
    // @Router /users/{userID} [get]
    GetUser(userID string) (models.User, error)
    
    // @Summary Create new user
    // @Description Creates a new user profile
    // @Version 1
    // @Param user models.CreateUserRequest body true
    // @Success 201 {object} models.User
    // @Failure 400 {object} models.ErrorResponse
    // @Router /users [post]
    CreateUser(user models.CreateUserRequest) (models.User, error)
}
```

### 2. Generate the Microservice

```bash
msgen create --module github.com/yourname/my-service --spec spec/spec.go
```

This generates:
- `main.go` - Application entry point
- `server/controller.go` - Implementation stub (edit this)
- `server/generated/` - Generated server code (don't edit)
  - `server.go` - HTTP server
  - `handlers.go` - HTTP handlers
  - `hooks.go` - Lifecycle hook interfaces
  - `controller/controller.go` - Controller interface
  - `ctx/context.go` - Request context
  - `web/` - HTTP utilities
- `msgen.json` - Configuration file

### 3. Implement Your Controller

Edit `server/controller.go` to implement your business logic:

```go
package server

import (
    "my-service/models"
    "my-service/server/generated/ctx"
    "my-service/server/generated/web"
    "net/http"
)

type UserServiceController struct {
    // Add your dependencies
}

func NewUserServiceController() *UserServiceController {
    return &UserServiceController{}
}

func (c *UserServiceController) GetUser(ctx *ctx.Context, userID string) (models.User, error) {
    // Your implementation here
    user := models.User{
        ID:   userID,
        Name: "John Doe",
    }
    return user, nil
}

func (c *UserServiceController) CreateUser(ctx *ctx.Context, req models.CreateUserRequest) (models.User, error) {
    // Validate and create user
    if req.Name == "" {
        return models.User{}, web.NewServerError(http.StatusBadRequest, "name is required")
    }
    
    user := models.User{
        ID:   "new-id",
        Name: req.Name,
    }
    return user, nil
}
```

### 4. Run Your Service

```bash
go run main.go
```

Your service is now running on `http://localhost:8080`!

## Annotation Reference

### Service-Level Annotations

```go
// @Summary Brief service description
// @Description Detailed service description
type YourService interface { ... }
```

### Method-Level Annotations

#### Required Annotations

- `@Router <path> [method]` - HTTP route and method (GET, POST, PUT, DELETE, PATCH)
- `@Success <code> {type} <model>` - Success response definition

#### Optional Annotations

- `@Summary <text>` - Brief method description
- `@Description <text>` - Detailed method description  
- `@Version <number>` - API version
- `@Param <name> <type> <in> <required>` - Parameter definition
  - `<in>`: `path`, `query`, or `body`
  - `<required>`: `true` or `false`
- `@Failure <code> {type} <model>` - Error response definition

#### Parameter Types

- **Path**: `@Param userID string path true`
- **Query**: `@Param country string query false`
- **Body**: `@Param user models.User body true`

### Example with All Annotations

```go
// @Summary Update user profile
// @Description Updates user information with validation
// @Version 1
// @Param userID string path true
// @Param user models.UpdateUserRequest body true
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /users/{userID} [put]
UpdateUser(userID string, user models.UpdateUserRequest) (models.User, error)
```

## Lifecycle Hooks

Implement optional hooks in your controller for server lifecycle events:

```go
// Called when server starts (after routes are configured, before accepting requests)
func (c *YourController) Start(logger *slog.Logger) error {
    // Initialize resources, connect to databases, etc.
    logger.Info("starting controller")
    return nil
}

// Called when server stops
func (c *YourController) Stop(logger *slog.Logger) error {
    // Cleanup resources
    logger.Info("stopping controller")
    return nil
}

// Customize the logger (called during server initialization)
func (c *YourController) ConfigureLogger(logger *slog.Logger) {
    // Customize logging configuration
}

// Add custom middleware to the main router
func (c *YourController) ConfigureRouter(r chi.Router, logger *slog.Logger) {
    r.Use(middleware.RequestID)
    r.Use(middleware.Logger)
    logger.Info("configured router middleware")
}

// Customize routers for specific path scopes
func (c *YourController) ConfigureScopedRouter(r chi.Router, path string, logger *slog.Logger) {
    // Add path-specific middleware
    if path == "/admin" {
        r.Use(middleware.BasicAuth("admin-realm", map[string]string{"admin": "password"}))
    }
    logger.Info("configured scoped router", "path", path)
}
```

## Generated Code Structure

```
your-project/
├── main.go                          # Entry point (generated once, can edit)
├── msgen.json                       # MSGen configuration
├── Dockerfile                       # Docker build config (generated once, can edit)
├── .dockerignore                    # Docker ignore file (generated once, can edit)
├── spec/
│   └── spec.go                      # Your service specification
├── models/                          # Your data models
│   ├── user.go
│   └── error.go
└── server/
    ├── controller.go                # Your implementation (generated once, can edit)
    └── generated/                   # Auto-generated code (DO NOT EDIT)
        ├── server.go                # HTTP server
        ├── handlers.go              # HTTP handlers
        ├── hooks.go                 # Lifecycle interfaces
        ├── controller/
        │   └── controller.go        # Controller interface
        ├── ctx/
        │   └── context.go           # Request context
        └── web/
            ├── responses.go         # Response helpers
            └── server_error.go      # Error types
```

## Commands

### Create

Initialize a new microservice project:

```bash
msgen create --module <module-url> --spec <spec-file> [--workdir <directory>]
```

**Flags:**
- `--module, -m` (required) - Go module URL (e.g., `github.com/user/project`)
- `--spec, -s` (required) - Path to specification file (e.g., `spec/spec.go`)
- `--workdir, -w` (optional) - Working directory (default: `.`)

**Example:**
```bash
msgen create -m github.com/myorg/userservice -s spec/spec.go
```

### Update

Regenerate code after specification changes:

```bash
msgen update --spec <spec-file> [--workdir <directory>]
```

**Flags:**
- `--spec, -s` (required) - Path to specification file
- `--workdir, -w` (optional) - Working directory (default: `.`)

**Example:**
```bash
# Add a new method to your spec interface
# Then regenerate:
msgen update -s spec/spec.go
```

**Note:** The `update` command regenerates all files in `generated/` directories but preserves your hand-written code in `controller.go` and other user-editable files.

## Working with Generated Code

### Files You Should Edit

- `main.go` - Entry point customization
- `server/controller.go` - Business logic implementation
- `models/*.go` - Data models
- `Dockerfile` - Docker build configuration
- `.dockerignore` - Files to exclude from Docker build context

### Files You Should NOT Edit

Any file in a `generated/` directory:
- `server/generated/**/*.go`

These files are regenerated when you run `msgen update`.

## Docker Support

MSGen automatically generates a multi-stage Dockerfile optimized for Go applications. This file is generated only once during project creation, allowing you to customize it for your specific needs.

### Generated Dockerfile

The default Dockerfile uses a multi-stage build to create a minimal production image:

```dockerfile
ARG GO_IMAGE_VERSION

# Stage 1: Build the Go application
FROM golang:${GO_IMAGE_VERSION} AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o YourService .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/YourService .

ENTRYPOINT ["/app/YourService"] 
CMD ["--port", "8080"]
```

### Building and Running with Docker

```bash
# Build the Docker image
docker build --build-arg GO_IMAGE_VERSION=1.25.6 -t my-service:latest .

# Run the container
docker run -p 8080:8080 my-service:latest

# Run with custom port
docker run -p 3000:3000 my-service:latest --port 3000
```

### Customizing the Dockerfile

Since the Dockerfile is only generated once, you're free to modify it as needed:

- Add additional build arguments or environment variables
- Install system dependencies in the Alpine image
- Change the base images
- Add health checks
- Configure multi-platform builds
- Add additional build steps or tools

The generated `.dockerignore` file excludes common development files from the build context to keep your images lean. You can also customize this file to suit your needs.

### Error Handling

Return structured errors using the generated `web.ServerError`:

```go
import (
    "my-service/server/generated/web"
    "net/http"
)

func (c *Controller) GetUser(ctx *ctx.Context, id string) (User, error) {
    if id == "" {
        return User{}, web.NewServerError(http.StatusBadRequest, "ID is required")
    }
    
    user, err := c.db.FindUser(id)
    if err != nil {
        return User{}, web.NewServerError(http.StatusNotFound, "User not found")
    }
    
    return user, nil
}
```

### Using the Context

Each handler receives a context with logger and HTTP request:

```go
func (c *Controller) GetUser(ctx *ctx.Context, id string) (User, error) {
    // Access logger
    ctx.Logger().Info("fetching user", "id", id)
    
    // Access raw HTTP request if needed
    req := ctx.Request()
    userAgent := req.Header.Get("User-Agent")
    
    // ...
}
```

## Configuration File

MSGen creates a `msgen.json` configuration file:

```json
{
  "codegen_version": "0.0.0-unreleased",
  "module_url": "github.com/yourname/your-service"
}
```

This tracks the code generator version (set to the msgen version at generation time) and module URL for consistency during updates.

## Examples

See the [example/](./example) directory for a complete working example with:
- Service specification with multiple endpoints
- Model definitions
- Controller implementation
- Full CRUD operations

To run the example:

```bash
cd example
go run main.go
```

Test the endpoints:

```bash
# Get user
curl http://localhost:8080/v1/users/123

# Create user
curl -X POST http://localhost:8080/v1/users \
  -H "Content-Type: application/json" \
  -d '{"username":"john","email":"john@example.com","country":"US"}'

# Filter users
curl http://localhost:8080/v1/users/filter?country=US
```

## Requirements

- Go 1.25.6 or higher
- Dependencies (automatically installed):
  - `github.com/spf13/cobra` - CLI framework
  - `github.com/go-chi/chi/v5` - HTTP router

## How It Works

1. **AST Parsing** - MSGen uses Go's `go/ast` and `go/parser` packages to analyze your interface specification
2. **Template Generation** - Embedded Go templates generate the server code based on your spec
3. **Code Generation** - Generated code includes:
   - HTTP handlers with parameter extraction and validation
   - Router configuration from `@Router` annotations
   - Type-safe controller interface
   - Error handling and response serialization
4. **Formatting** - Generated code is formatted with `go fmt`

## Limitations & Future Work

Current limitations (WIP):
- Only supports REST/JSON APIs (no GraphQL, gRPC)
- Limited to basic Go types and model references
- No OpenAPI/Swagger generation yet
- No client SDK generation

Planned features:
- OpenAPI/Swagger documentation generation
- Custom middleware support
- Validation tag support
- Client SDK generation
- Testing utilities generation

## Contributing

This project is a work in progress. Contributions, issues, and feature requests are welcome!

## License

See [LICENSE](LICENSE) file for details.

## Author

adm87 - https://github.com/adm87

## Disclaimer

This README was AI generated.
