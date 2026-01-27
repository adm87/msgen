ARG GO_IMAGE_VERSION

# Stage 1: Build the Go application
FROM golang:${GO_IMAGE_VERSION} AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ExampleService .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/ExampleService .

ENTRYPOINT ["/app/ExampleService"] 
CMD ["--port", "8080"]