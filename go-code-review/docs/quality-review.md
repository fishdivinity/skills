# Code Quality Review Checklist

## 1. Error Handling

### Error Checking

- [ ] All errors are checked (no ignored errors)
- [ ] Errors are wrapped with context using `%w`
- [ ] Error types are checked with `errors.As`
- [ ] Custom errors implement `Unwrap()` for chaining

```go
// GOOD: Proper error handling
if err != nil {
    return fmt.Errorf("failed to process request %s: %w", reqID, err)
}

// GOOD: Type assertion with errors.As
var validationErr *ValidationError
if errors.As(err, &validationErr) {
    return handleValidationError(validationErr)
}

// BAD: Ignored error
file.Close()  // Error not checked

// BAD: Lost context
if err != nil {
    return err  // No context added
}
```

### Panic Usage

- [ ] No panic in business logic
- [ ] Panic only for truly unrecoverable states
- [ ] Recover used appropriately in HTTP handlers

```go
// GOOD: Recover in middleware
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("panic recovered: %v", err)
                http.Error(w, "internal error", 500)
            }
        }()
        next.ServeHTTP(w, r)
    })
}

// BAD: Panic in business logic
func CalculateTotal(items []Item) float64 {
    if len(items) == 0 {
        panic("no items")  // Should return error instead
    }
    // ...
}
```

---

## 2. Resource Management

### Defer Usage

- [ ] Resources are closed with defer
- [ ] No defer in loops (memory leak)
- [ ] Defer is in the same function that creates resource

```go
// GOOD: Defer immediately after resource creation
f, err := os.Open(path)
if err != nil {
    return err
}
defer f.Close()

// BAD: Defer in loop
for _, path := range paths {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()  // Files not closed until function returns!
}

// GOOD: Close in loop body
for _, path := range paths {
    func() {
        f, err := os.Open(path)
        if err != nil {
            return
        }
        defer f.Close()
        // process file
    }()
}
```

### Goroutine Management

- [ ] Goroutines have a way to stop (context or channel)
- [ ] Goroutine leaks are prevented
- [ ] WaitGroup or errgroup used for coordination

```go
// GOOD: Context for cancellation
func ProcessItems(ctx context.Context, items <-chan Item) {
    for {
        select {
        case item, ok := <-items:
            if !ok {
                return
            }
            process(item)
        case <-ctx.Done():
            return
        }
    }
}

// BAD: Goroutine leak
func ProcessItems(items <-chan Item) {
    for item := range items {
        go process(item)  // No way to stop these goroutines
    }
}
```

---

## 3. Concurrency Safety

### Race Conditions

- [ ] Shared data is protected by mutex or channels
- [ ] No data races (run with `-race` flag)
- [ ] Maps are not accessed concurrently without synchronization

```go
// GOOD: Mutex protection
type SafeCounter struct {
    mu    sync.RWMutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

// BAD: Concurrent map access
var cache = make(map[string]string)

func Get(key string) string {
    return cache[key]  // Race condition!
}
```

### Deadlock Prevention

- [ ] Locks acquired in consistent order
- [ ] Locks always released (with defer)
- [ ] Timeouts on blocking operations

```go
// GOOD: Consistent lock order
func Transfer(a, b *Account, amount int) {
    // Always lock lower ID first
    first, second := a, b
    if a.ID > b.ID {
        first, second = b, a
    }
    first.mu.Lock()
    defer first.mu.Unlock()
    second.mu.Lock()
    defer second.mu.Unlock()
    // ...
}
```

---

## 4. Memory Management

### Memory Leaks

- [ ] No unbounded slice growth
- [ ] Large objects are released when done
- [ ] Caches have size limits and eviction

```go
// GOOD: Bounded cache
type Cache struct {
    items map[string]*Item
    mu    sync.RWMutex
    max   int
}

func (c *Cache) Add(key string, item *Item) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    if len(c.items) >= c.max {
        // Evict oldest
        for k := range c.items {
            delete(c.items, k)
            break
        }
    }
    c.items[key] = item
}

// BAD: Unbounded growth
var cache = make(map[string]*Item)
```

### Slice Pre-allocation

- [ ] Slices pre-allocated when size is known
- [ ] `append` doesn't cause excessive reallocations

```go
// GOOD: Pre-allocate
result := make([]Item, 0, len(items))
for _, item := range items {
    result = append(result, process(item))
}

// OK but less efficient
var result []Item
for _, item := range items {
    result = append(result, process(item))
}
```

---

## 5. Code Readability

### Naming

- [ ] Names are descriptive and consistent
- [ ] No cryptic abbreviations
- [ ] Interface names are nouns or adjectives (Reader, Writer, Stringer)

### Function Length

- [ ] Functions are focused on single responsibility
- [ ] Functions are < 50 lines (guideline)
- [ ] Complex logic is extracted to helper functions

### Comments

- [ ] Public functions have documentation comments
- [ ] Complex logic is explained
- [ ] No commented-out code

```go
// GOOD: Documentation comment
// ProcessOrder validates and processes an order.
// It returns an error if the order is invalid or processing fails.
func ProcessOrder(ctx context.Context, order *Order) error {
    // ...
}

// GOOD: Explaining complex logic
// We use a two-phase commit because we need to ensure
// both the inventory and payment systems are updated atomically.
func commitTransaction(tx *Transaction) error {
    // ...
}
```

---

## 6. Go Idioms

### Interface Design

- [ ] Interfaces are small and focused
- [ ] Interfaces defined where they are used
- [ ] Accept interfaces, return structs

```go
// GOOD: Small interface
type Reader interface {
    Read(p []byte) (n int, err error)
}

// BAD: Large interface
type Service interface {
    CreateUser() error
    DeleteUser() error
    UpdateUser() error
    GetUser() error
    CreateOrder() error
    // ... many more methods
}
```

### Error Values

- [ ] Sentinel errors for comparison
- [ ] Error types for type assertion
- [ ] Custom errors provide useful information

```go
// GOOD: Sentinel error
var ErrNotFound = errors.New("not found")

// GOOD: Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on %s: %s", e.Field, e.Message)
}
```

---

## 7. Go 1.26+ Modernization Guide

### New Recommended Patterns

#### errors.AsType (Go 1.26+)

The new generic version of `errors.As` is type-safe, faster, and recommended over `errors.As`:

```go
// Go 1.26+: Recommended - type-safe and faster
var validationErr *ValidationError
if errors.AsType(err, &validationErr) {
    handleValidationError(validationErr)
}

// Legacy pattern (still valid but not recommended)
var validationErr *ValidationError
if errors.As(err, &validationErr) {
    handleValidationError(validationErr)
}
```

#### new() with Expression (Go 1.26+)

The built-in `new` function now accepts expressions for optional fields:

```go
// Go 1.26+: Recommended for optional fields
type Person struct {
    Name string `json:"name"`
    Age  *int   `json:"age"` // nil if unknown
}

func personJSON(name string, born time.Time) ([]byte, error) {
    return json.Marshal(Person{
        Name: name,
        Age:  new(yearsSince(born)), // Clean and idiomatic
    })
}

// Legacy pattern (still valid but verbose)
func personJSONLegacy(name string, born time.Time) ([]byte, error) {
    age := yearsSince(born)
    return json.Marshal(Person{
        Name: name,
        Age:  &age,
    })
}
```

#### Generic Self-Reference (Go 1.26+)

Type constraints can now refer to themselves:

```go
// Go 1.26+: Self-referencing constraints are now legal
type Adder[A Adder[A]] interface {
    Add(A) A
}

func algo[A Adder[A]](x, y A) A {
    return x.Add(y)
}
```

#### io.ReadAll Optimization (Go 1.26+)

`io.ReadAll` now allocates less memory and is \~2x faster:

- **No code change needed**: Existing calls benefit automatically
- **Review note**: Don't suggest alternatives for existing `io.ReadAll` usage

#### B.Loop in Benchmarks (Go 1.26+)

`B.Loop` no longer prevents inlining:

```go
// Go 1.26+: Recommended for benchmarks
func BenchmarkFunc(b *testing.B) {
    for b.Loop() {
        // ...
    }
}

// Legacy pattern (b.N prevents inlining)
func BenchmarkFuncLegacy(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // ...
    }
}
```

### Deprecated Patterns to Flag

#### crypto Subpackages

Multiple crypto functions now always use secure random sources:

```go
// Now ALWAYS uses secure random (parameter ignored)
GenerateKey DSA/Crypto/ECDSA/Ed25519/RSA(..., random) // random parameter ignored
```

**Flag when**: Code relies on the old behavior for testing (should use `testing/cryptotest.SetGlobalRandom` instead).

#### crypto/rsa PKCS#1 v1.5

Unsafe PKCS #1 v1.5 encryption is now deprecated:

```go
// DEPRECATED: Use EncryptOAEPWithOptions instead
EncryptPKCS1v15(...)
DecryptPKCS1v15(...)
DecryptPKCS1v15SessionKey(...)
```

#### tls.Config Settings (Go 1.27)

These GODEBUG settings will be removed in Go 1.27:

| Setting           | Behavior Change                            |
| ----------------- | ------------------------------------------ |
| `tlsunsafeekm`    | ExportKeyingMaterial requires TLS 1.3      |
| `tlsrsakex`       | RSA-only key exchanges disabled by default |
| `tls10server`     | Minimum TLS version becomes 1.2            |
| `tls3des`         | 3DES cipher suites removed                 |
| `x509keypairleaf` | Certificate.Leaf always populated          |

### go fix Recommendations

The rewritten `go fix` command automates API migrations:

```bash
# Run modernizers to update code automatically
go fix ./...

# Preview changes without applying
go fix -diff ./...
```

**When to suggest**: When code uses deprecated patterns that `go fix` can automatically update.

---

## 8. Context-Aware Review Guidelines

### Code Type Recognition

#### Generated Code

Generated code should NOT be reviewed for optimization or style issues:

- **Swagger/OpenAPI specs**: `swagger.json`, `openapi.yaml` - generated by `swag init`
- **Protobuf**: `*.pb.go` files - generated by protoc
- **Mock files**: `*_mock.go` files - generated by mockgen
- **YAML/JSON schemas**: Generated from definitions

**When to flag**: Only flag generated code if it causes actual build errors or security vulnerabilities.

```go
// DO NOT FLAG: Generated swagger documentation
// @Summary Get user by ID
// @Param id path int true "User ID"
// GENERATED BY: swag init -g handlers/user.go
```

#### Template Code

Code generated from templates should follow same rules as generated code:

- **Boilerplate reduction**: If patterns repeat across files, consider template improvements
- **Not personal failure**: Style issues in template output are template problems, not author problems

### Output Destination Awareness

#### Interactive vs Non-Interactive Context

**Flag inappropriate output behavior based on context:**

| Context Type      | Example               | Appropriate Output     |
| ----------------- | --------------------- | ---------------------- |
| CLI Tool          | User-facing commands  | stdout/stderr messages |
| Interactive Shell | Terminal application  | User prompts, progress |
| Daemon Service    | dockerfile ENTRYPOINT | Log files, not stdout  |
| Initialization    | Config generation     | stdout acceptable      |
| Debug Mode        | Development builds    | Verbose output OK      |

**When to flag sensitive output:**

```go
// ACCEPTABLE: Docker/daemon context - stdout is only option
func main() {
    password := generatePassword()
    fmt.Println("Admin password:", password)  // OK for init containers
}

// ACCEPTABLE: Interactive CLI context
func initUser() {
    password := promptPassword("Enter admin password:")
    fmt.Printf("Password set: %s\n", maskPassword(password))
}

// FLAG: Sensitive data in inappropriate contexts
if os.Getenv("ENV") == "production" {
    fmt.Println("DB_PASSWORD:", dbPassword)  // FLAG: Never in prod
}
```

**Guidelines:**

- **CLI tools**: Output to user via stdout is normal and expected
- **Services (Dockerfile/daemon)**: stdout is acceptable when no log file configured
- **Production environments**: Never output secrets, even to stdout
- **Configuration options**: Suggest adding flags to control output behavior

### Project-Specific Conventions

#### When to Ask

When encountering patterns that could be legitimate project conventions:

1. **Large service initializations**: Ask if startup output is expected
2. **Generated file modifications**: Ask if intentional customization was made
3. **Non-standard patterns**: Ask about team-specific practices

#### Example Questions

- "This swagger.json appears to be auto-generated. Should I skip optimization review?"
- "This service prints to stdout during startup. Is this intentional for Docker logging?"
- "This configuration is typically generated. Has it been manually modified?"

### P2 Issue Clarification

**P2 (Suggested Fix) should NOT include:**

1. Generated code style issues
2. Output patterns appropriate for the deployment context
3. Template-generated boilerplate
4. Code that follows documented project conventions

**P2 SHOULD include:**

1. Suboptimal algorithms or data structures
2. Missing error handling (when not by design)
3. Performance improvements with clear ROI
4. Code smell that reduces maintainability
5. Inconsistent naming or structure (when not by convention)

