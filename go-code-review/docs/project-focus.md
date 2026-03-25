# Project-Specific Review Focus

## 1. Detecting Project Type

### API Service

**Indicators:**
- `main.go` starts HTTP/gRPC server
- `handlers/` or `controllers/` directory
- Route definitions
- Middleware chain

**Focus Areas:**
- Input validation
- Error response format
- Rate limiting
- Authentication/Authorization
- Request logging

**Key Files to Review:**
```
cmd/server/main.go
internal/handlers/
internal/middleware/
internal/routes/
```

---

### CLI Tool

**Indicators:**
- Uses `flag`, `cobra`, or `urfave/cli`
- `main.go` with flag parsing
- Multiple subcommands

**Focus Areas:**
- Flag validation
- Error messages (user-friendly)
- Exit codes
- Help text
- Configuration file support

**Key Files to Review:**
```
cmd/cli/main.go
internal/commands/
internal/flags/
```

---

### Library/Package

**Indicators:**
- No `cmd/` directory (or minimal)
- Public API in root or `pkg/`
- Extensive documentation

**Focus Areas:**
- API stability
- Documentation quality
- Backward compatibility
- Error types
- Interface design

**Key Files to Review:**
```
pkg/
doc.go
README.md
```

---

### Microservice

**Indicators:**
- Service discovery integration
- Health check endpoints
- Metrics export
- Circuit breaker patterns

**Focus Areas:**
- Service registration
- Graceful shutdown
- Health/readiness probes
- Distributed tracing
- Circuit breakers

**Key Files to Review:**
```
cmd/service/main.go
internal/health/
internal/metrics/
internal/client/
```

---

### Data Pipeline / Worker

**Indicators:**
- Message queue consumers
- Background job processing
- Batch processing logic

**Focus Areas:**
- Idempotency
- Error recovery
- Dead letter handling
- Memory management
- Progress tracking

**Key Files to Review:**
```
cmd/worker/main.go
internal/processor/
internal/queue/
```

---

## 2. Review Focus by Project Type

### API Service Checklist

| Area | Priority | Items |
|------|----------|-------|
| **Security** | Critical | Auth middleware, input validation, rate limiting |
| **Error Handling** | High | Consistent error format, proper HTTP codes |
| **Logging** | High | Request ID, structured logging |
| **Performance** | Medium | Connection pooling, query optimization |
| **Observability** | Medium | Metrics, tracing |

### CLI Tool Checklist

| Area | Priority | Items |
|------|----------|-------|
| **UX** | Critical | Clear error messages, helpful output |
| **Flags** | High | Validation, defaults, help text |
| **Exit Codes** | High | Consistent exit codes |
| **Testing** | Medium | Command testing, flag testing |
| **Documentation** | Medium | Usage examples, man pages |

### Library Checklist

| Area | Priority | Items |
|------|----------|-------|
| **API Design** | Critical | Clear interfaces, sensible defaults |
| **Documentation** | Critical | Package docs, example code |
| **Compatibility** | Critical | No breaking changes |
| **Error Types** | High | Custom errors, error wrapping |
| **Testing** | High | Comprehensive test coverage |

### Microservice Checklist

| Area | Priority | Items |
|------|----------|-------|
| **Resilience** | Critical | Circuit breakers, retries, timeouts |
| **Observability** | Critical | Health checks, metrics, tracing |
| **Configuration** | High | External config, secrets management |
| **Graceful Shutdown** | High | Context cancellation, cleanup |
| **Service Discovery** | High | Registration, health checks |

### Data Pipeline Checklist

| Area | Priority | Items |
|------|----------|-------|
| **Reliability** | Critical | Idempotency, error recovery |
| **Performance** | Critical | Memory management, batch sizing |
| **Monitoring** | High | Progress tracking, alerting |
| **Data Integrity** | High | Validation, dead letter handling |
| **Scalability** | Medium | Parallel processing, backpressure |

---

## 3. Detecting Inconsistencies

### Middleware Adoption

**Scenario:** New auth middleware added, but some handlers still use hardcoded auth.

**Detection Method:**

```
1. Find middleware definition
2. Find routes using middleware
3. Find routes NOT using middleware
4. Check handlers for hardcoded auth
```

**Example:**

```go
// Middleware defined
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if !validateToken(token) {
            http.Error(w, "unauthorized", 401)
            return
        }
        next.ServeHTTP(w, r)
    })
}

// Routes using middleware
api := r.PathPrefix("/api").Subrouter()
api.Use(authMiddleware)
api.HandleFunc("/users", getUsers)    // Uses middleware

// Routes NOT using middleware
r.HandleFunc("/api/orders", getOrders)  // Missing middleware!

// Handler with hardcoded auth (inconsistent)
func getOrders(w http.ResponseWriter, r *http.Request) {
    token := r.Header.Get("Authorization")
    if token != "secret-token" {  // Hardcoded!
        http.Error(w, "unauthorized", 401)
        return
    }
    // ...
}
```

**Flag:** Mixed authentication patterns detected

---

### Database Access Pattern

**Scenario:** New repository pattern introduced, but some code still uses direct DB access.

**Detection Method:**

```
1. Find repository interfaces
2. Find usages of repository
3. Find direct DB calls (db.Query, db.Exec)
4. Compare locations
```

**Example:**

```go
// Repository pattern (new)
type UserRepository interface {
    GetByID(ctx context.Context, id string) (*User, error)
}

// Using repository (consistent)
func (s *Service) GetUser(ctx context.Context, id string) (*User, error) {
    return s.userRepo.GetByID(ctx, id)
}

// Direct DB access (inconsistent)
func (s *Service) GetOrder(ctx context.Context, id string) (*Order, error) {
    var order Order
    err := s.db.QueryRowContext(ctx, "SELECT * FROM orders WHERE id = ?", id).Scan(&order)
    // Direct DB access, should use OrderRepository
}
```

**Flag:** Mixed data access patterns detected

---

### Error Handling Pattern

**Scenario:** New error types introduced, but some code still uses plain errors.

**Detection Method:**

```
1. Find custom error types
2. Find usages of custom errors
3. Find plain errors.New() or fmt.Errorf()
4. Check if plain errors should be custom
```

**Example:**

```go
// Custom errors (new)
type NotFoundError struct {
    Resource string
    ID       string
}

func (e *NotFoundError) Error() string {
    return fmt.Sprintf("%s with ID %s not found", e.Resource, e.ID)
}

// Using custom error (consistent)
func (r *UserRepo) GetByID(ctx context.Context, id string) (*User, error) {
    // ...
    return nil, &NotFoundError{Resource: "user", ID: id}
}

// Plain error (inconsistent)
func (r *OrderRepo) GetByID(ctx context.Context, id string) (*Order, error) {
    // ...
    return nil, errors.New("order not found")  // Should use NotFoundError
}
```

**Flag:** Inconsistent error types detected

---

## 4. Context Size Management

### For Large Projects

**Strategy:** Read incrementally, summarize non-critical files

```
Round 1: Entry points (main.go, handlers)
Round 2: Core business logic (services)
Round 3: Data layer (repositories, models)
Round 4: Supporting code (utilities, middleware)
```

### Context Budget

| Total Files | Strategy |
|-------------|----------|
| 1-5 | Read all files |
| 6-15 | Read primary files, list secondary |
| 16-30 | Read changed files + direct dependencies |
| 31+ | Incremental review required |

### Incremental Review Process

```
1. Get git diff to identify changed files
2. Build dependency graph of changed files
3. Read changed files + immediate dependencies
4. Review in logical groups:
   - Security-critical code first
   - Business logic second
   - Tests and docs last
```

---

## 5. Memory-Based Workflow

See [docs/memory-workflow.md](memory-workflow.md) for complete details.

### Quick Reference

| Condition | Use Memory |
|-----------|------------|
| Files changed > 30 | Yes |
| Lines changed > 2000 | Yes |
| Cross-module changes | Yes |

### Memory File Location

```
<project-root>/.review/memory.md
```

---

## 6. Project Analysis Commands

```bash
# Identify project type
cat main.go | grep -E "(http.ListenAndServe|flag\.|cobra)"

# Find all HTTP handlers
grep -rn "http.HandlerFunc\|HandleFunc\|Handle(" --include="*.go"

# Find database access
grep -rn "db\.Query\|db\.Exec\|\.QueryRow" --include="*.go"

# Find middleware
grep -rn "func.*http\.Handler.*http\.Handler" --include="*.go"

# Find error definitions
grep -rn "type.*Error struct\|var Err.*= errors.New" --include="*.go"

# Find authentication code
grep -rn "Authorization\|Bearer\|JWT\|token" --include="*.go"
```
