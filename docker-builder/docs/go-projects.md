# Go Project Docker Guidelines

## Go Version Detection

### Automatic Detection
1. Run `go version` to detect current Go version
2. If Go is not installed, ask user for desired Go version

### Version Format
- Use `golang:{{GO_VERSION}}-alpine{{ALPINE_VERSION}}` as base image
- Example: `golang:1.25.6-alpine3.22`

## SQLite/CGO Detection

### Detection Logic
- Search for `_ "github.com/mattn/go-sqlite3"` in imports
- Check for `CGO_ENABLED` references in code

### Handling
- If SQLite is detected, ensure `CGO_ENABLED=1` during build
- Include `gcc` and `musl-dev` in build dependencies

## Version Command Support

### Detection
- Search for version-related code in CLI implementation
- Look for `version` command definition

### Git Integration
- If version uses git information, include `.git` directory in build
- Use `git describe --tags` for versioning if needed

## Executable Name

### Automatic Detection
- Use project directory name as default executable name
- Ask user to confirm or provide custom name

## Build Configuration

### Multi-stage Build
```dockerfile
# Build stage
FROM golang:{{GO_VERSION}}-alpine{{ALPINE_VERSION}} AS builder
WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git gcc musl-dev

# Copy source code
COPY . .

# Build application
RUN go build -o {{EXECUTABLE_NAME}} .

# Final stage
FROM alpine:{{ALPINE_VERSION}}
WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy executable
COPY --from=builder /app/{{EXECUTABLE_NAME}} /app/

# Set entrypoint
ENTRYPOINT ["./{{EXECUTABLE_NAME}}"]
```

## Environment Variables

### Common Variables
- `GO_ENV`: Production/Development
- `PORT`: Server port
- `DATABASE_URL`: Database connection string

### Detection
- Search for `os.Getenv` usage in code
- Check for configuration files that use environment variables