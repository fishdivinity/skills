# Common Go Anti-Patterns

## 1. Error Handling Anti-Patterns

### Ignoring Errors

```go
// BAD: Ignored error
file.Close()
rows.Close()

// GOOD: Handle error
if err := file.Close(); err != nil {
    log.Printf("warning: failed to close file: %v", err)
}

// GOOD: Use defer with named returns
func ProcessFile(path string) (err error) {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer func() {
        if closeErr := file.Close(); closeErr != nil {
            err = fmt.Errorf("close file: %w (original: %v)", closeErr, err)
        }
    }()
    // ...
}
```

### Losing Error Context

```go
// BAD: Lost context
if err != nil {
    return err
}

// GOOD: Add context
if err != nil {
    return fmt.Errorf("failed to process user %s: %w", userID, err)
}
```

### Using panic for Business Logic

```go
// BAD: Panic for business error
func GetUser(id string) *User {
    user := db.Find(id)
    if user == nil {
        panic("user not found")
    }
    return user
}

// GOOD: Return error
func GetUser(id string) (*User, error) {
    user := db.Find(id)
    if user == nil {
        return nil, ErrUserNotFound
    }
    return user, nil
}
```

---

## 2. Concurrency Anti-Patterns

### Goroutine Leak

```go
// BAD: Goroutine leak
func Process(ch chan int) {
    go func() {
        for v := range ch {
            process(v)
        }
    }()
}

// GOOD: Context cancellation
func Process(ctx context.Context, ch chan int) {
    go func() {
        for {
            select {
            case v, ok := <-ch:
                if !ok {
                    return
                }
                process(v)
            case <-ctx.Done():
                return
            }
        }
    }()
}
```

### Race Condition

```go
// BAD: Shared state without synchronization
var counter int

func Increment() {
    counter++  // Race condition!
}

// GOOD: Use sync/atomic
var counter int64

func Increment() {
    atomic.AddInt64(&counter, 1)
}

// GOOD: Use mutex
var (
    counter int
    mu      sync.Mutex
)

func Increment() {
    mu.Lock()
    counter++
    mu.Unlock()
}
```

### Blocking in Hot Path

```go
// BAD: Blocking channel in hot path
func Process(items []Item) {
    ch := make(chan Result)
    for _, item := range items {
        go func(i Item) {
            ch <- process(i)  // Blocks if no receiver
        }(item)
    }
}

// GOOD: Buffered channel
func Process(items []Item) {
    ch := make(chan Result, len(items))
    for _, item := range items {
        go func(i Item) {
            ch <- process(i)
        }(item)
    }
}
```

---

## 3. Memory Anti-Patterns

### Unbounded Memory

```go
// BAD: Unbounded slice
var results []Result
for {
    r := <-ch
    results = append(results, r)  // Grows forever
}

// GOOD: Bounded with eviction
const maxResults = 1000
results := make([]Result, 0, maxResults)
for {
    r := <-ch
    if len(results) >= maxResults {
        results = results[1:]  // Evict oldest
    }
    results = append(results, r)
}
```

### Large Allocations in Loop

```go
// BAD: Allocate in loop
func Process(data [][]byte) [][]byte {
    var results [][]byte
    for _, d := range data {
        result := make([]byte, len(d))  // Many allocations
        copy(result, d)
        results = append(results, result)
    }
    return results
}

// GOOD: Pre-allocate
func Process(data [][]byte) [][]byte {
    results := make([][]byte, len(data))
    for i, d := range data {
        results[i] = make([]byte, len(d))
        copy(results[i], d)
    }
    return results
}
```

### String Concatenation in Loop

```go
// BAD: String concatenation
func BuildString(parts []string) string {
    var result string
    for _, p := range parts {
        result += p  // O(n^2) complexity
    }
    return result
}

// GOOD: Use strings.Builder
func BuildString(parts []string) string {
    var b strings.Builder
    b.Grow(len(parts) * 20)  // Estimate size
    for _, p := range parts {
        b.WriteString(p)
    }
    return b.String()
}
```

---

## 4. Interface Anti-Patterns

### Large Interface

```go
// BAD: Large interface
type Repository interface {
    GetUser() (*User, error)
    SaveUser(*User) error
    GetOrder() (*Order, error)
    SaveOrder(*Order) error
    GetProduct() (*Product, error)
    // ... 20 more methods
}

// GOOD: Small, focused interfaces
type UserRepository interface {
    Get(ctx context.Context, id string) (*User, error)
    Save(ctx context.Context, user *User) error
}

type OrderRepository interface {
    Get(ctx context.Context, id string) (*Order, error)
    Save(ctx context.Context, order *Order) error
}
```

### Returning Interface

```go
// BAD: Return interface
type UserService interface {
    GetUser(id string) (*User, error)
}

func NewUserService() UserService {
    return &userService{}  // Limits flexibility
}

// GOOD: Return concrete type
func NewUserService() *UserService {
    return &UserService{}
}
```

### Pointer to Interface

```go
// BAD: Pointer to interface
func Process(user *UserInterface) { }

// GOOD: Just interface
func Process(user UserInterface) { }
```

---

## 5. Context Anti-Patterns

### Not Passing Context

```go
// BAD: No context
func GetUser(id string) (*User, error) {
    return db.Query("SELECT * FROM users WHERE id = ?", id)
}

// GOOD: Accept context
func GetUser(ctx context.Context, id string) (*User, error) {
    return db.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", id)
}
```

### Storing Context in Struct

```go
// BAD: Context in struct
type Service struct {
    ctx context.Context
}

func (s *Service) Process() {
    // ctx might be stale
}

// GOOD: Pass context to methods
type Service struct {
    db DB
}

func (s *Service) Process(ctx context.Context) error {
    return s.db.QueryContext(ctx, "...")
}
```

---

## 6. Channel Anti-Patterns

### Unbuffered Channel Without Receiver

```go
// BAD: Deadlock potential
func Send(ch chan int) {
    ch <- 42  // Blocks if no receiver
}

// GOOD: Buffered or select with default
func Send(ch chan int) {
    select {
    case ch <- 42:
    default:
        // Handle full channel
    }
}
```

### Closing Channel from Sender

```go
// BAD: Close from sender (can cause panic)
func producer(ch chan int) {
    for i := 0; i < 10; i++ {
        ch <- i
    }
    close(ch)  // What if receiver closes too?
}

// GOOD: Sender never closes, use done channel
func producer(ctx context.Context, ch chan int) {
    for i := 0; i < 10; i++ {
        select {
        case ch <- i:
        case <-ctx.Done():
            return
        }
    }
}
```

---

## 7. Defer Anti-Patterns

### Defer in Loop

```go
// BAD: Defer in loop
func ProcessFiles(paths []string) error {
    for _, path := range paths {
        f, err := os.Open(path)
        if err != nil {
            return err
        }
        defer f.Close()  // All files open until function returns!
    }
    return nil
}

// GOOD: Extract to function
func ProcessFiles(paths []string) error {
    for _, path := range paths {
        if err := processFile(path); err != nil {
            return err
        }
    }
    return nil
}

func processFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()
    // Process file
    return nil
}
```

---

## 8. Slice Anti-Patterns

### Slice Append Shares Memory

```go
// BAD: Potential memory leak
func GetItems(data []byte) []byte {
    return data[10:20]  // Shares underlying array
}

// GOOD: Copy if needed
func GetItems(data []byte) []byte {
    result := make([]byte, 10)
    copy(result, data[10:20])
    return result
}
```

### Nil vs Empty Slice

```go
// BAD: Inconsistent nil handling
func GetItems() []Item {
    var items []Item
    if noItems {
        return nil  // nil slice
    }
    return items  // empty slice
}

// GOOD: Be consistent
func GetItems() []Item {
    return []Item{}  // Always return empty slice
}
```

---

## Anti-Pattern Detection Checklist

| Category | Pattern | Detection Method |
|----------|---------|------------------|
| Error | Ignored error | `grep -E "_ = .*\(" *.go` |
| Error | panic in business | `grep -n "panic(" *.go` |
| Concurrency | Goroutine without context | Check goroutine has ctx |
| Concurrency | Shared state | Look for global vars |
| Memory | Unbounded growth | Check slice/map growth |
| Interface | Large interface | Count methods |
| Context | Missing context | Check function signatures |
| Channel | Close from sender | Check close() calls |
