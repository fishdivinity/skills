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

# Set environment variables
{{ENVIRONMENT_VARIABLES}}

# Set entrypoint
ENTRYPOINT ["./{{EXECUTABLE_NAME}}"]