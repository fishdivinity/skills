# Common Issues in Go Code

## Critical Issues (Must Fix)

### Security
- **SQL Injection**: Unparameterized SQL queries
- **Hardcoded Secrets**: API keys, passwords in source code
- **Missing Authentication**: Unprotected endpoints
- **Cross-Site Scripting (XSS)**: Unsanitized user input
- **Insecure Cryptography**: Weak algorithms or implementation

### Correctness
- **Race Conditions**: Unprotected concurrent access
- **Nil Pointer Dereference**: Accessing nil values
- **Resource Leaks**: Unclosed files, connections
- **Error Handling**: Ignored or improperly handled errors
- **Logic Errors**: Incorrect business logic implementation

### Performance
- **N+1 Queries**: Multiple database queries instead of batch operations
- **Unbounded Memory**: Memory leaks or excessive allocation
- **Blocking in Hot Paths**: Synchronous operations in critical code
- **Inefficient Algorithms**: Suboptimal time complexity
- **Excessive Logging**: High-volume logging in production

## Common Anti-Patterns

### Error Handling

#### Bad Example: Ignored Error
```go
// Bad: Error is ignored
_ = file.Close()
```

#### Good Example: Proper Error Handling
```go
// Good: Error is properly handled
if err := file.Close(); err != nil {
    return fmt.Errorf("close file: %w", err)
}
```

#### Bad Example: Panic in Business Logic
```go
// Bad: Panic in business logic
func ProcessOrder(orderID string) {
    order, err := db.GetOrder(orderID)
    if err != nil {
        panic("failed to get order") // Avoid panic in business logic
    }
    // Process order
}
```

#### Good Example: Return Error Instead of Panic
```go
// Good: Return error instead of panic
func ProcessOrder(orderID string) error {
    order, err := db.GetOrder(orderID)
    if err != nil {
        return fmt.Errorf("get order: %w", err)
    }
    // Process order
    return nil
}
```

### Code Structure

#### Bad Example: Global Mutable State
```go
// Bad: Global mutable state
var db *sql.DB

func InitDB() {
    var err error
    db, err = sql.Open("postgres", "connection-string")
    if err != nil {
        log.Fatal(err)
    }
}
```

#### Good Example: Dependency Injection
```go
// Good: Dependency injection
type Service struct {
    db *sql.DB
}

func NewService(db *sql.DB) *Service {
    return &Service{db: db}
}
```

### Security

#### Bad Example: Unparameterized SQL
```go
// Bad: Unparameterized SQL
userID := "123"
db.Query("SELECT * FROM users WHERE id = " + userID)
```

#### Good Example: Parameterized SQL
```go
// Good: Parameterized SQL
userID := "123"
db.Query("SELECT * FROM users WHERE id = $1", userID)
```

### Concurrency

#### Bad Example: Context Not Passed
```go
// Bad: Context not passed
func ProcessData(data []byte) error {
    // No context parameter
    // Process data
    return nil
}
```

#### Good Example: Context Passed
```go
// Good: Context passed
func ProcessData(ctx context.Context, data []byte) error {
    // Use context for cancellation and timeouts
    // Process data
    return nil
}
```

### Logging

#### Bad Example: Using fmt.Print instead of logger
```go
// Bad: Using fmt.Print instead of logger
fmt.Printf("Error: %v\n", err)
```

#### Good Example: Using proper logger
```go
// Good: Using proper logger
log.Errorf("Failed to process request: %v", err)
```

### Performance

#### Bad Example: Unpreallocated Slice
```go
// Bad: Unpreallocated slice
var results []string
for i := 0; i < 1000; i++ {
    results = append(results, fmt.Sprintf("item %d", i))
}
```

#### Good Example: Preallocated Slice
```go
// Good: Preallocated slice
results := make([]string, 0, 1000)
for i := 0; i < 1000; i++ {
    results = append(results, fmt.Sprintf("item %d", i))
}
```

### Memory Management

#### Bad Example: Unclosed File
```go
// Bad: Unclosed file
func ReadFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    // File is not closed if ReadAll fails
    return ioutil.ReadAll(file)
}
```

#### Good Example: Deferred File Close
```go
// Good: Deferred file close
func ReadFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close() // File will be closed regardless of error
    return ioutil.ReadAll(file)
}

## Detection Strategies

### Static Analysis
- Use `golangci-lint` for comprehensive checks
- Run `go vet` for common issues
- Use `gosec` for security-specific checks

### Code Review Checklist
- [ ] Check for SQL injection vulnerabilities
- [ ] Verify error handling is comprehensive
- [ ] Review concurrency patterns
- [ ] Validate input sanitization
- [ ] Check for resource leaks
- [ ] Review logging practices
- [ ] Verify proper use of context
- [ ] Check for hardcoded secrets