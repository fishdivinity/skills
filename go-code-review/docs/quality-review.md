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

## Quality Review Priority Matrix

| Issue Type | Severity | Action |
|------------|----------|--------|
| Race Condition | Critical | Block merge |
| Resource Leak | Critical | Block merge |
| Ignored Error | High | Must fix |
| Panic in Business Logic | High | Must fix |
| Missing Documentation | Medium | Flag |
| Long Functions | Low | Suggest refactoring |
