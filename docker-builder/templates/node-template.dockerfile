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

# Set environment variables
{{ENVIRONMENT_VARIABLES}}

# Set entrypoint
CMD ["npm", "start"]