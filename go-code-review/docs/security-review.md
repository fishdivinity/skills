# Security Review Checklist

## 1. Authentication & Authorization

### Authentication

- [ ] All API endpoints (except login/register) have authentication
- [ ] JWT tokens are properly validated (signature, expiration, issuer)
- [ ] Passwords are hashed with bcrypt/argon2 (not stored in plain text)
- [ ] Session tokens are cryptographically random
- [ ] Failed login attempts are rate-limited

```go
// GOOD: Proper JWT validation
func validateToken(tokenString string) (*Claims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
        }
        return jwtSecret, nil
    })
    // ...
}

// BAD: Missing algorithm validation
token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
    return jwtSecret, nil  // Vulnerable to algorithm confusion
})
```

### Authorization

- [ ] Role-based access control is implemented
- [ ] Resource ownership is checked before operations
- [ ] Privilege escalation is prevented
- [ ] Authorization checks happen after authentication

```go
// GOOD: Check ownership
func (s *Service) DeleteOrder(ctx context.Context, userID, orderID string) error {
    order, err := s.repo.GetOrder(ctx, orderID)
    if err != nil {
        return err
    }
    if order.OwnerID != userID {
        return ErrUnauthorized
    }
    return s.repo.DeleteOrder(ctx, orderID)
}
```

---

## 2. Input Validation

### SQL Injection Prevention

- [ ] All SQL queries use parameterized statements
- [ ] No string concatenation for SQL
- [ ] ORM/query builder used correctly

```go
// GOOD: Parameterized query
db.Query("SELECT * FROM users WHERE id = $1", userID)

// BAD: SQL injection vulnerable
db.Query("SELECT * FROM users WHERE id = " + userID)
```

### Input Sanitization

- [ ] All user input is validated
- [ ] Input length limits are enforced
- [ ] Special characters are properly escaped
- [ ] Content-Type is validated for file uploads

```go
// GOOD: Validate input
type CreateUserRequest struct {
    Email    string `validate:"required,email,max=255"`
    Username string `validate:"required,alphanum,min=3,max=50"`
    Password string `validate:"required,min=8,max=72"`
}

func (s *Service) CreateUser(ctx context.Context, req *CreateUserRequest) error {
    if err := validator.Struct(req); err != nil {
        return fmt.Errorf("validation failed: %w", err)
    }
    // ...
}
```

---

## 3. Secrets Management

### No Hardcoded Secrets

- [ ] No passwords in source code
- [ ] No API keys in source code
- [ ] No private keys in source code
- [ ] Secrets loaded from environment or secret manager

```go
// GOOD: Load from environment
dbPassword := os.Getenv("DB_PASSWORD")
apiKey := config.Get("API_KEY")

// BAD: Hardcoded secret
const dbPassword = "super-secret-123"
```

### Secret Handling

- [ ] Secrets are not logged
- [ ] Secrets are not returned in API responses
- [ ] Secrets are masked in error messages
- [ ] Secret rotation is supported

```go
// GOOD: Mask sensitive data in logs
log.WithField("password", "***").Info("user authenticated")

// BAD: Logging secrets
log.WithField("password", password).Info("user authenticated")
```

---

## 4. Cryptography

### Encryption

- [ ] TLS 1.2+ for all network communication
- [ ] AES-256-GCM for symmetric encryption
- [ ] RSA-2048+ or ECDSA for asymmetric encryption
- [ ] Proper IV/nonce management

### Hashing

- [ ] bcrypt/argon2 for passwords
- [ ] SHA-256+ for data integrity
- [ ] HMAC for message authentication

```go
// GOOD: bcrypt for passwords
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// BAD: Fast hash for passwords
hashedPassword := sha256.Sum256([]byte(password))
```

---

## 5. HTTP Security

### Headers

- [ ] Content-Security-Policy header set
- [ ] X-Content-Type-Options: nosniff
- [ ] X-Frame-Options: DENY or SAMEORIGIN
- [ ] Strict-Transport-Security (HSTS) enabled

### Request Handling

- [ ] Request size limits enforced
- [ ] Request timeout configured
- [ ] Rate limiting implemented
- [ ] CORS properly configured

```go
// GOOD: Security middleware
func SecurityHeaders(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        next.ServeHTTP(w, r)
    })
}
```

---

## 6. Dependency Security

- [ ] Dependencies are from trusted sources
- [ ] Dependencies are version-pinned
- [ ] No known vulnerabilities in dependencies
- [ ] Regular dependency updates

```bash
# Check for vulnerabilities
go list -m -json all | nancy sleuth
govulncheck ./...
```

---

## 7. Error Handling

- [ ] Errors don't expose internal details
- [ ] Stack traces not returned to clients
- [ ] Sensitive errors are logged, not displayed
- [ ] Generic error messages for authentication failures

```go
// GOOD: Generic error message
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
    // ...
    if err != nil {
        log.WithError(err).Error("login failed")
        http.Error(w, "invalid credentials", http.StatusUnauthorized)
        return
    }
}

// BAD: Exposes internal details
if err == ErrUserNotFound {
    http.Error(w, "user not found", http.StatusNotFound)
}
```

---

## Security Review Priority Matrix

| Issue Type | Severity | Action |
|------------|----------|--------|
| SQL Injection | Critical | Block merge |
| Hardcoded Secrets | Critical | Block merge |
| Missing Auth | Critical | Block merge |
| Weak Crypto | High | Must fix before merge |
| Missing Rate Limit | Medium | Flag for follow-up |
| Missing Security Headers | Medium | Flag for follow-up |
