# Node.js Project Docker Guidelines

## Node.js Version Detection

### Automatic Detection
1. Check `package.json` for `engines` field
2. If not specified, ask user for desired Node.js version

### Version Format
- Use `node:{{NODE_VERSION}}-alpine{{ALPINE_VERSION}}` as base image
- Example: `node:20-alpine3.22`

## Build Configuration

### Multi-stage Build
```dockerfile
# Build stage
FROM node:{{NODE_VERSION}}-alpine{{ALPINE_VERSION}} AS builder
WORKDIR /app

# Install dependencies
COPY package*.json ./
RUN npm install

# Copy source code
COPY . .

# Build application
RUN npm run build

# Final stage
FROM node:{{NODE_VERSION}}-alpine{{ALPINE_VERSION}}
WORKDIR /app

# Install runtime dependencies only
COPY package*.json ./
RUN npm install --only=production

# Copy built application
COPY --from=builder /app/build /app/build

# Set entrypoint
CMD ["npm", "start"]
```

## Network Access

### During Build
- Node.js projects typically require network access for npm installs
- Ensure build environment has network connectivity

### Runtime
- Configure appropriate network settings based on application needs

## Environment Variables

### Common Variables
- `NODE_ENV`: Production/Development
- `PORT`: Server port
- `API_URL`: Backend API URL

### Detection
- Search for `process.env` usage in code
- Check for `.env` files or environment configuration

## Dependencies

### Build Dependencies
- Include build tools if needed (e.g., `python3`, `make`, `g++`)
- Use `npm ci` for deterministic builds

### Runtime Dependencies
- Use `npm install --only=production` for smaller images
- Consider using `npm prune` to remove unnecessary files

## Frontend Specifics

### Static Builds
- For static frontend builds, consider using Nginx as final stage
- Example Nginx configuration:

```dockerfile
# Final stage with Nginx
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

### Development Mode
- For development, include live reload and hot module replacement
- Map local files for real-time changes