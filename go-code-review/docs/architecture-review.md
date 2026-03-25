# Architecture Review Checklist

## 1. Package Structure

### Standard Layout

- [ ] Follows standard Go project layout (if applicable)
- [ ] Packages are cohesive (single responsibility)
- [ ] Import cycles are avoided
- [ ] Internal packages used appropriately

```
Standard Project Layout:
├── cmd/           # Main applications
├── internal/      # Private application code
├── pkg/           # Public library code
├── api/           # API definitions
├── configs/       # Configuration files
├── scripts/       # Build/deploy scripts
└── go.mod
```

### Package Naming

- [ ] Package names are lowercase, single word
- [ ] Package names describe the purpose
- [ ] No stutter in names (e.g., `http.HTTPServer` is bad)

```go
// GOOD
package user
type Service struct { ... }

// BAD
package userService
type UserService struct { ... }
```

---

## 2. Dependency Management

### Dependency Injection

- [ ] Dependencies are injected (not created inside)
- [ ] Interfaces are used for external dependencies
- [ ] Constructor functions return interfaces

```go
// GOOD: Dependency injection
type UserService struct {
    db     UserRepository
    cache  Cache
    logger Logger
}

func NewUserService(db UserRepository, cache Cache, logger Logger) *UserService {
    return &UserService{db: db, cache: cache, logger: logger}
}

// BAD: Hardcoded dependencies
type UserService struct {
    db *sql.DB  // Concrete type, created inside
}
```

### Interface Design

- [ ] Interfaces are defined by consumers
- [ ] Interfaces are small and focused
- [ ] Accept interfaces, return structs

```go
// GOOD: Consumer-defined interface
type UserRepository interface {
    GetByID(ctx context.Context, id string) (*User, error)
    Save(ctx context.Context, user *User) error
}

// BAD: Large interface
type Repository interface {
    GetUser() (*User, error)
    GetOrder() (*Order, error)
    GetProduct() (*Product, error)
    // ... 20 more methods
}
```

---

## 3. Layered Architecture

### Handler/Service/Repository Pattern

- [ ] Handlers handle HTTP, no business logic
- [ ] Services contain business logic
- [ ] Repositories handle data access
- [ ] Clear separation of concerns

```go
// Handler - HTTP layer only
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        respondError(w, http.StatusBadRequest, err)
        return
    }
    
    user, err := h.service.CreateUser(r.Context(), req)
    if err != nil {
        respondError(w, http.StatusInternalServerError, err)
        return
    }
    
    respondJSON(w, http.StatusCreated, user)
}

// Service - Business logic
func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*User, error) {
    if err := s.validateUser(req); err != nil {
        return nil, err
    }
    
    user := &User{
        ID:       generateID(),
        Email:    req.Email,
        Name:     req.Name,
        Created:  time.Now(),
    }
    
    if err := s.repo.Save(ctx, user); err != nil {
        return nil, fmt.Errorf("save user: %w", err)
    }
    
    return user, nil
}

// Repository - Data access only
func (r *Repository) Save(ctx context.Context, user *User) error {
    query := `INSERT INTO users (id, email, name, created) VALUES (?, ?, ?, ?)`
    _, err := r.db.ExecContext(ctx, query, user.ID, user.Email, user.Name, user.Created)
    return err
}
```

---

## 4. Middleware Pattern

### Middleware Chain

- [ ] Middleware follows standard pattern
- [ ] Request context is used for passing data
- [ ] Middleware order is correct

```go
// Standard middleware pattern
type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
    for i := len(middlewares) - 1; i >= 0; i-- {
        h = middlewares[i](h)
    }
    return h
}

// Example middleware
func LoggingMiddleware(logger Logger) Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            next.ServeHTTP(w, r)
            logger.Info("request completed",
                "method", r.Method,
                "path", r.URL.Path,
                "duration", time.Since(start),
            )
        })
    }
}
```

### Middleware Inconsistency Detection

**Check for partial middleware adoption:**

```go
// INCONSISTENT: Some routes use middleware, others don't
func setupRoutes(r *mux.Router, auth AuthMiddleware) {
    // Uses auth middleware
    api := r.PathPrefix("/api").Subrouter()
    api.Use(auth.Middleware)
    api.HandleFunc("/users", getUsers).Methods("GET")
    
    // BUT: Direct handler without middleware
    r.HandleFunc("/api/orders", getOrders).Methods("GET")  // Missing auth!
}

// CONSISTENT: All protected routes use middleware
func setupRoutes(r *mux.Router, auth AuthMiddleware) {
    api := r.PathPrefix("/api").Subrouter()
    api.Use(auth.Middleware)
    api.HandleFunc("/users", getUsers).Methods("GET")
    api.HandleFunc("/orders", getOrders).Methods("GET")
    
    // Public routes explicitly separated
    public := r.PathPrefix("/public").Subrouter()
    public.HandleFunc("/health", healthCheck).Methods("GET")
}
```

---

## 5. Configuration Management

### Environment Configuration

- [ ] Configuration is externalized
- [ ] Environment variables are used for secrets
- [ ] Default values are sensible
- [ ] Configuration is validated at startup

```go
type Config struct {
    DatabaseURL  string `env:"DATABASE_URL,required"`
    RedisURL     string `env:"REDIS_URL" envDefault:"localhost:6379"`
    LogLevel     string `env:"LOG_LEVEL" envDefault:"info"`
    JWTSecret    string `env:"JWT_SECRET,required"`
    Port         int    `env:"PORT" envDefault:"8080"`
}

func LoadConfig() (*Config, error) {
    cfg := &Config{}
    if err := env.Parse(cfg); err != nil {
        return nil, fmt.Errorf("parse config: %w", err)
    }
    return cfg, nil
}
```

---

## 6. Error Handling Architecture

### Error Propagation

- [ ] Errors are propagated to appropriate layer
- [ ] Errors are wrapped with context
- [ ] Errors are translated at boundaries

```go
// Repository layer
func (r *Repository) GetByID(ctx context.Context, id string) (*User, error) {
    var user User
    err := r.db.GetContext(ctx, &user, "SELECT * FROM users WHERE id = ?", id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, ErrUserNotFound
        }
        return nil, fmt.Errorf("query user: %w", err)
    }
    return &user, nil
}

// Service layer
func (s *Service) GetUser(ctx context.Context, id string) (*User, error) {
    user, err := r.repo.GetByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("get user %s: %w", id, err)
    }
    return user, nil
}

// Handler layer
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
    user, err := h.service.GetUser(r.Context(), id)
    if err != nil {
        if errors.Is(err, ErrUserNotFound) {
            respondError(w, http.StatusNotFound, "user not found")
            return
        }
        respondError(w, http.StatusInternalServerError, "internal error")
        return
    }
    respondJSON(w, http.StatusOK, user)
}
```

---

## Architecture Review Priority Matrix

| Issue Type | Severity | Action |
|------------|----------|--------|
| Import Cycles | Critical | Block merge |
| Missing Dependency Injection | High | Must fix |
| Layer Violation | High | Must fix |
| Large Interface | Medium | Suggest refactoring |
| Missing Middleware | Medium | Flag for review |
| Hardcoded Config | Medium | Flag for follow-up |
