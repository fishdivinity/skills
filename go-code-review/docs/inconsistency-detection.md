# Inconsistency Detection Guide

## Overview

Detect partial adoption of new patterns across the codebase. Common scenarios:
- Middleware adoption (some routes use it, others don't)
- Repository pattern (mixed direct DB access)
- Error handling (mixed error types)

---

## Detection Strategy

```
1. Identify the new pattern (e.g., middleware usage)
2. Search for:
   - Direct usage of the new pattern
   - Old pattern still in use (hardcoded values, direct calls)
3. Compare call sites:
   - Which code paths use new pattern?
   - Which code paths still use old pattern?
4. Flag inconsistencies for review
```

---

## Common Inconsistency Patterns

### 1. Middleware Inconsistency

**Scenario:** New auth middleware added, but some handlers still use hardcoded auth.

```go
// NEW: Using auth middleware
router.Use(authMiddleware)
router.GET("/api/users", getUsers)

// OLD: Still hardcoded in some handlers
func getOrders(w http.ResponseWriter, r *http.Request) {
    token := r.Header.Get("Authorization")
    if token != "hardcoded-token" {  // INCONSISTENT!
        http.Error(w, "unauthorized", 401)
    }
}
```

**Detection:**
```bash
# Find middleware usage
grep -rn "router.Use\|\.Use(" --include="*.go"

# Find hardcoded auth
grep -rn "Authorization.*hardcoded\|token.*=.*\"" --include="*.go"
```

---

### 2. Database Access Pattern

**Scenario:** New repository pattern introduced, but some code still uses direct DB access.

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

**Detection:**
```bash
# Find repository interfaces
grep -rn "type.*Repository interface" --include="*.go"

# Find direct DB calls
grep -rn "db\.Query\|db\.Exec\|\.QueryRow" --include="*.go"
```

---

### 3. Error Handling Pattern

**Scenario:** New error types introduced, but some code still uses plain errors.

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

**Detection:**
```bash
# Find custom error types
grep -rn "type.*Error struct\|type.*Error interface" --include="*.go"

# Find plain errors
grep -rn "errors.New\|fmt.Errorf" --include="*.go"
```

---

### 4. Configuration Pattern

**Scenario:** New config library introduced, but some code still uses os.Getenv directly.

```go
// Config struct (new)
type Config struct {
    DatabaseURL string `env:"DATABASE_URL,required"`
}

// Using config (consistent)
func NewService(cfg *Config) *Service { ... }

// Direct env access (inconsistent)
func NewOrderService() *OrderService {
    dbURL := os.Getenv("DATABASE_URL")  // Should use config
}
```

**Detection:**
```bash
# Find config usage
grep -rn "config\.\|cfg\." --include="*.go"

# Find direct env access
grep -rn "os\.Getenv\|os\.LookupEnv" --include="*.go"
```

---

### 5. Logging Pattern

**Scenario:** New structured logging introduced, but some code still uses log.Printf.

```go
// Structured logging (new)
logger.Info("user created", "user_id", userID, "email", email)

// Printf logging (inconsistent)
log.Printf("user created: %s, %s", userID, email)  // Should use structured logging
```

**Detection:**
```bash
# Find structured logging
grep -rn "logger\.\|slog\.\|zap\." --include="*.go"

# Find printf logging
grep -rn "log\.Printf\|fmt\.Printf" --include="*.go"
```

---

## Detection Checklist

When reviewing for inconsistencies:

- [ ] New abstraction fully adopted?
- [ ] No mixed patterns in same module?
- [ ] Migration path documented?
- [ ] Tests updated for both patterns?
- [ ] No TODO comments indicating incomplete migration?

---

## Reporting Inconsistencies

When found, report in this format:

```markdown
### Inconsistency: [Pattern Name]

**New Pattern:** [description]
**Old Pattern:** [description]

**Files using new pattern:**
- file1.go
- file2.go

**Files using old pattern:**
- file3.go (line 45)
- file4.go (line 120)

**Recommendation:** [action to take]
```
