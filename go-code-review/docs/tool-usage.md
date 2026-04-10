# Tool Usage Guide

## Recommended Tools

### Static Analysis
- **golangci-lint**: Comprehensive linting tool
  - Installation: `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`
  - Usage: `golangci-lint run`
  - Configuration: `.golangci.yml` file in project root

- **go vet**: Built-in Go vetting tool
  - Usage: `go vet ./...`
  - Focus: Common mistakes and suspicious constructs

- **shadow**: Variable shadowing detection
  - Installation: `go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest`
  - Usage: `shadow ./...`

### Security Scanning
- **gosec**: Security scanner for Go code
  - Installation: `go install github.com/securego/gosec/cmd/gosec@latest`
  - Usage: `gosec ./...`
  - Focus: Common security issues

### Performance Analysis
- **pprof**: Go profiling tool
  - Usage:
    - CPU profiling: `go test -cpuprofile cpu.prof ./...`
    - Memory profiling: `go test -memprofile mem.prof ./...`
    - Analysis: `go tool pprof cpu.prof`

### Code Coverage
- **go test -cover**: Built-in coverage tool
  - Usage: `go test -cover ./...`
  - For HTML report: `go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out`

## Integration Tips

### Git Integration
- **git diff**: Identify changed files
  - Usage: `git diff HEAD~1 --name-only`
  - For specific files: `git diff HEAD~1 -- path/to/file.go`

- **git log**: Review commit history
  - Usage: `git log --oneline -n 10`
  - For file history: `git log --oneline -p -- path/to/file.go`

### IDE Features
- **Code Navigation**: Use IDE features to jump between files
- **Go to Definition**: Quickly understand function implementations
- **Find References**: Identify all usages of a variable or function
- **Code Completion**: Ensure consistent naming and reduce typos

### Automation Scripts
- **Pre-commit Hooks**: Run checks before commits
  - Example: `.git/hooks/pre-commit` script

- **Makefiles**: Automate common tasks
  - Example:
    ```makefile
    lint:
        golangci-lint run
    test:
        go test ./...
    cover:
        go test -cover ./...
    ```

## CI/CD Integration

### GitHub Actions
- **Example Workflow**:
  ```yaml
  name: Go Code Review
  
  on:
    push:
      branches: [ main, master ]
    pull_request:
      branches: [ main, master ]
  
  jobs:
    review:
      runs-on: ubuntu-latest
      steps:
        - uses: actions/checkout@v3
        
        - name: Set up Go
          uses: actions/setup-go@v3
          with:
            go-version: 1.26
        
        - name: Install dependencies
          run: go mod tidy
        
        - name: Lint
          run: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && golangci-lint run
        
        - name: Test
          run: go test ./...
        
        - name: Security Scan
          run: go install github.com/securego/gosec/cmd/gosec@latest && gosec ./...
  ```

### GitLab CI
- **Example Configuration**:
  ```yaml
  stages:
    - test
    - security
  
  lint:
    stage: test
    script:
      - go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
      - golangci-lint run
  
  test:
    stage: test
    script:
      - go test ./...
  
  security:
    stage: security
    script:
      - go install github.com/securego/gosec/cmd/gosec@latest
      - gosec ./...
  ```

### Jenkins
- **Pipeline Example**:
  ```groovy
  pipeline {
    agent {
      docker {
        image 'golang:1.26'
      }
    }
    stages {
      stage('Lint') {
        steps {
          sh 'go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest'
          sh 'golangci-lint run'
        }
      }
      stage('Test') {
        steps {
          sh 'go test ./...'
        }
      }
      stage('Security') {
        steps {
          sh 'go install github.com/securego/gosec/cmd/gosec@latest'
          sh 'gosec ./...'
        }
      }
    }
  }
  ```

## Custom Review Tool

### Installation
- **From Source**:
  ```bash
  cd review
  make install  # Build and install to tools directory
  ```

- **Build Only**:
  ```bash
  cd review
  make build  # Build to bin directory
  ```

### Usage
- **Analyze Project**:
  ```bash
  tools/review analyze --path <project-path>
  ```

- **Generate Report**:
  ```bash
  tools/review report --format markdown
  ```

### Features
- **Map Preallocation Detection**: Identify maps that should be preallocated
- **Duplicate Code Detection**: Find repeated code patterns
- **Scale Detection**: Analyze project size and complexity
- **Markdown Reports**: Generate structured review reports

### Configuration
- **Command Line Options**:
  - `--path`: Project path (required)
  - `--format`: Output format (markdown, json)
  - `--verbose`: Enable verbose output

## Troubleshooting

### Common Issues
- **Tool Installation Failures**:
  - Ensure Go is properly installed
  - Check GOPATH and GOBIN environment variables
  - Use `go install` with specific versions

- **False Positives**:
  - Configure linters to ignore specific patterns
  - Use `//nolint` comments for legitimate exceptions

- **Performance Issues**:
  - Run tools incrementally on changed files
  - Use caching where available
  - Limit analysis to relevant directories