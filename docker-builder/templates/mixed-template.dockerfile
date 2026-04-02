# Frontend build stage
FROM node:{{NODE_VERSION}}-alpine{{ALPINE_VERSION}} AS frontend-builder
WORKDIR /app/frontend

# Install frontend dependencies
COPY frontend/package*.json ./
RUN npm install

# Copy frontend source
COPY frontend/ .

# Build frontend
RUN npm run build

# Backend build stage
FROM golang:{{GO_VERSION}}-alpine{{ALPINE_VERSION}} AS backend-builder
WORKDIR /app/backend

# Install backend dependencies
RUN apk add --no-cache git gcc musl-dev

# Copy backend source
COPY backend/ .

# Build backend
RUN go build -o {{EXECUTABLE_NAME}} .

# Final stage
FROM alpine:{{ALPINE_VERSION}}
WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy backend executable
COPY --from=backend-builder /app/backend/{{EXECUTABLE_NAME}} /app/

# Copy frontend build
COPY --from=frontend-builder /app/frontend/build /app/frontend

# Set environment variables
{{ENVIRONMENT_VARIABLES}}

# Set entrypoint
ENTRYPOINT ["./{{EXECUTABLE_NAME}}"]