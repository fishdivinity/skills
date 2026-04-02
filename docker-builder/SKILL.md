---
name: docker-builder
description: Creates optimized Dockerfiles for various project types. Invoke when user needs to build Docker images, ensuring correct base images, dependencies, and configuration for both frontend and backend projects.
---

# Docker Builder Skill

This skill creates optimized Dockerfiles for various project types, addressing common issues and ensuring best practices.

## When to Invoke

- User needs to create a Dockerfile for any project
- User wants to build a Docker image
- User needs help with Docker configuration
- User mentions Docker in the context of building applications

---

## Project Types Supported

- **Go Backend**
- **Node.js Frontend**
- **Mixed Frontend/Backend**
- **Other Language Projects**

---

## Core Workflow

### Step 1: Project Analysis
- Detect project type
- Identify required dependencies
- Check for special requirements (SQLite, CGO, etc.)

### Step 2: Version Detection
- **Go**: Check `go version` or ask user
- **Node.js**: Check `package.json` or ask user
- **Other**: Ask user for appropriate version

### Step 3: Base Image Selection
- Ask user for base image preferences
- Determine appropriate Alpine version
- Select optimized base images

### Step 4: Build Configuration
- **Executable Name**: Use project name or ask user
- **Build Arguments**: Configure based on project needs
- **Environment Variables**: Set up required variables

### Step 5: Generate Dockerfile
- Create multi-stage build if appropriate
- Include necessary dependencies
- Configure proper entrypoint

---

## Detection Logic

### Go Projects
- Search for `go.mod` files
- Check for SQLite usage (`_ "github.com/mattn/go-sqlite3"`)
- Look for version command implementation

### Node.js Projects
- Search for `package.json` files
- Check for build scripts
- Identify Node.js version requirements

### Mixed Projects
- Detect both frontend and backend components
- Create appropriate multi-stage builds

---

## User Interaction

### Required Questions
- **Base Image Version**: Alpine version for backend
- **Node.js Version**: For frontend projects
- **Executable Name**: If not auto-detected
- **Special Requirements**: Any project-specific needs

### Optional Questions
- **Build Arguments**: Custom build arguments
- **Environment Variables**: Additional environment variables
- **Network Access**: Whether build requires network access

---

## Best Practices

- Use multi-stage builds for smaller final images
- Set proper build arguments based on project needs
- Include necessary dependencies in build stages
- Use minimal runtime dependencies
- Configure proper entrypoint for applications
- Handle environment variables appropriately

---

## File Index

```
docker-builder/
├── SKILL.md           # This file - main entry point
├── docs/              # Documentation
│   ├── go-projects.md     # Go project specific guidelines
│   ├── node-projects.md   # Node.js project specific guidelines
│   └── mixed-projects.md  # Mixed frontend/backend guidelines
└── templates/         # Dockerfile templates
    ├── go-template.dockerfile     # Go project template
    ├── node-template.dockerfile   # Node.js project template
    └── mixed-template.dockerfile  # Mixed project template
```