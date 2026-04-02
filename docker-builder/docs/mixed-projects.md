# Mixed Frontend/Backend Docker Guidelines

## Project Detection

### Detection Logic
- Look for both `go.mod` (backend) and `package.json` (frontend)
- Check directory structure for separate frontend/backend folders

## Multi-stage Build Strategy

### Complete Build Flow
```dockerfile
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

# Set entrypoint
ENTRYPOINT ["./{{EXECUTABLE_NAME}}"]
```

## Directory Structure

### Common Structures
- **Separate Folders**: `/frontend` and `/backend`
- **Monorepo**: Single repository with multiple packages
- **Integrated**: Frontend and backend in same directory

### Handling Different Structures
- Adjust COPY commands based on actual directory structure
- Use appropriate build contexts for each component

## Environment Variables

### Shared Variables
- `API_URL`: Backend API URL for frontend
- `PORT`: Server port
- `ENVIRONMENT`: Production/Development

### Component-specific Variables
- **Backend**: Database connection strings, secret keys
- **Frontend**: API endpoints, client-side configuration

## Network Configuration

### Internal Communication
- Ensure frontend can communicate with backend
- Configure appropriate network settings

### External Access
- Expose necessary ports
- Set up proper routing

## Build Optimization

### Parallel Builds
- Build frontend and backend simultaneously when possible
- Use build arguments for version control

### Caching
- Optimize layer caching for faster builds
- Separate dependency installation from source code changes

## Deployment Considerations

### Single Container vs. Multiple Containers
- **Single Container**: Simpler deployment, shared resources
- **Multiple Containers**: Better isolation, independent scaling

### Orchestration
- Consider Docker Compose for local development
- Use Kubernetes for production deployment