# Testing Review Checklist

## 1. Test Coverage

### Coverage Requirements

- [ ] Critical paths have 100% coverage
- [ ] Overall coverage meets project threshold (typically 80%+)
- [ ] Edge cases are tested
- [ ] Error paths are tested

```bash
# Run coverage
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## 2. Unit Tests

### Test Structure

- [ ] Tests follow Arrange-Act-Assert pattern
- [ ] Test names describe what is being tested
- [ ] Each test focuses on one behavior

```go
// GOOD: Clear test structure
func TestUserService_CreateUser_ValidInput(t *testing.T) {
    // Arrange
    mockRepo := &MockUserRepository{
        SaveFunc: func(ctx context.Context, user *User) error {
            return nil
        },
    }
    service := NewUserService(mockRepo)
    ctx := context.Background()
    
    // Act
    user, err := service.CreateUser(ctx, "test@example.com", "Test User")
    
    // Assert
    require.NoError(t, err)
    assert.Equal(t, "test@example.com", user.Email)
    assert.Equal(t, "Test User", user.Name)
    assert.NotEmpty(t, user.ID)
}

func TestUserService_CreateUser_DuplicateEmail(t *testing.T) {
    // Arrange
    mockRepo := &MockUserRepository{
        SaveFunc: func(ctx context.Context, user *User) error {
            return ErrDuplicateEmail
        },
    }
    service := NewUserService(mockRepo)
    
    // Act
    _, err := service.CreateUser(context.Background(), "existing@example.com", "Test")
    
    // Assert
    require.Error(t, err)
    assert.ErrorIs(t, err, ErrDuplicateEmail)
}
```

### Table-Driven Tests

- [ ] Use table-driven tests for multiple scenarios
- [ ] Test cases are clearly named
- [ ] Each case is independent

```go
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr bool
    }{
        {
            name:    "valid email",
            email:   "user@example.com",
            wantErr: false,
        },
        {
            name:    "missing @",
            email:   "userexample.com",
            wantErr: true,
        },
        {
            name:    "empty string",
            email:   "",
            wantErr: true,
        },
        {
            name:    "missing domain",
            email:   "user@",
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateEmail(tt.email)
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
            }
        })
    }
}
```

---

## 3. Mocking

### Mock Usage

- [ ] Interfaces are mocked, not concrete types
- [ ] Mocks verify expected behavior
- [ ] Mocks are reset between tests

```go
// Using testify/mock
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*User, error) {
    args := m.Called(ctx, id)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*User), args.Error(1)
}

func TestUserService_GetUser(t *testing.T) {
    mockRepo := new(MockUserRepository)
    service := NewUserService(mockRepo)
    
    expectedUser := &User{ID: "123", Email: "test@example.com"}
    mockRepo.On("GetByID", mock.Anything, "123").Return(expectedUser, nil)
    
    user, err := service.GetUser(context.Background(), "123")
    
    require.NoError(t, err)
    assert.Equal(t, expectedUser, user)
    mockRepo.AssertExpectations(t)
}
```

---

## 4. Integration Tests

### Integration Test Setup

- [ ] Uses test containers or test databases
- [ ] Tests are isolated from each other
- [ ] Resources are cleaned up after tests

```go
// Using testcontainers
func TestUserRepository_Integration(t *testing.T) {
    if testing.Short() {
        t.Skip("skipping integration test")
    }
    
    ctx := context.Background()
    
    // Setup container
    container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
        ContainerRequest: testcontainers.ContainerRequest{
            Image: "postgres:15",
            Env: map[string]string{
                "POSTGRES_DB":       "testdb",
                "POSTGRES_PASSWORD": "test",
            },
            ExposedPorts: []string{"5432/tcp"},
        },
    })
    require.NoError(t, err)
    defer container.Terminate(ctx)
    
    // Run tests
    // ...
}
```

### Build Tags

- [ ] Integration tests use build tags
- [ ] Unit tests run without external dependencies

```go
//go:build integration
// +build integration

package repository_test

import "testing"

func TestIntegration(t *testing.T) {
    // Integration test code
}
```

```bash
# Run unit tests only
go test ./...

# Run integration tests
go test -tags=integration ./...
```

---

## 5. HTTP Handler Tests

### Handler Testing

- [ ] Request/response is tested
- [ ] Status codes are verified
- [ ] Error responses are tested

```go
func TestCreateUserHandler(t *testing.T) {
    tests := []struct {
        name       string
        body       string
        mockReturn *User
        mockError  error
        wantStatus int
    }{
        {
            name:       "success",
            body:       `{"email":"test@example.com","name":"Test"}`,
            mockReturn: &User{ID: "123", Email: "test@example.com"},
            wantStatus: http.StatusCreated,
        },
        {
            name:       "invalid json",
            body:       `{invalid`,
            wantStatus: http.StatusBadRequest,
        },
        {
            name:       "service error",
            body:       `{"email":"test@example.com","name":"Test"}`,
            mockError:  errors.New("database error"),
            wantStatus: http.StatusInternalServerError,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mockService := &MockUserService{}
            if tt.mockReturn != nil || tt.mockError != nil {
                mockService.On("CreateUser", mock.Anything, mock.Anything).Return(tt.mockReturn, tt.mockError)
            }
            
            handler := NewHandler(mockService)
            req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(tt.body))
            req.Header.Set("Content-Type", "application/json")
            rec := httptest.NewRecorder()
            
            handler.ServeHTTP(rec, req)
            
            assert.Equal(t, tt.wantStatus, rec.Code)
        })
    }
}
```

---

## 6. Benchmark Tests

### Performance Testing

- [ ] Critical functions have benchmarks
- [ ] Benchmarks test realistic scenarios
- [ ] Allocations are measured

```go
func BenchmarkUserService_CreateUser(b *testing.B) {
    mockRepo := &MockUserRepository{
        SaveFunc: func(ctx context.Context, user *User) error {
            return nil
        },
    }
    service := NewUserService(mockRepo)
    ctx := context.Background()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, _ = service.CreateUser(ctx, "test@example.com", "Test")
    }
}

func BenchmarkValidateEmail(b *testing.B) {
    email := "test.user+tag@example.com"
    for i := 0; i < b.N; i++ {
        _ = ValidateEmail(email)
    }
}
```

---

## 7. Fuzz Tests (Go 1.18+)

### Fuzz Testing

- [ ] Input validation is fuzzed
- [ ] Parsing functions are fuzzed

```go
func FuzzValidateEmail(f *testing.F) {
    // Add seed corpus
    f.Add("user@example.com")
    f.Add("invalid-email")
    f.Add("")
    
    f.Fuzz(func(t *testing.T, email string) {
        // Should not panic for any input
        ValidateEmail(email)
    })
}

func FuzzParseUser(f *testing.F) {
    f.Add(`{"id":"1","email":"test@example.com"}`)
    
    f.Fuzz(func(t *testing.T, data string) {
        // Should not panic for any input
        _, _ = ParseUser([]byte(data))
    })
}
```

---

## Testing Review Priority Matrix

| Issue Type | Severity | Action |
|------------|----------|--------|
| Missing Critical Path Tests | Critical | Block merge |
| Flaky Tests | Critical | Block merge |
| Missing Error Path Tests | High | Must fix |
| Low Coverage (<50%) | High | Must fix |
| Missing Integration Tests | Medium | Flag |
| Missing Benchmarks | Low | Suggest |
