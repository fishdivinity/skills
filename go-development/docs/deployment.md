# Go 构建与部署规范

## Docker多阶段构建（Multi-stage Build）

```dockerfile
# 构建阶段（Build Stage）
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /app/server

# 运行阶段（Runtime Stage）
FROM alpine:3.19
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
```

---

## 健康检查（Health Check）

实现以下端点（Endpoints）：
- `/livez` - 存活检查（Liveness Probe）
- `/readyz` - 就绪检查（Readiness Probe）

```go
http.HandleFunc("/livez", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
})

http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
    if !isReady() {
        http.Error(w, "not ready", http.StatusServiceUnavailable)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
})
```

---

## 版本注入（Version Injection）

```powershell
go build -ldflags="-X main.version=1.0.0 -X main.commit=$(git rev-parse HEAD) -X main.buildTime=$(date -u '+%Y-%m-%d_%H:%M:%S')"
```

---

## 环境隔离（Environment Isolation）

- 开发环境（Development）
- 测试环境（Testing/Staging）
- 生产环境（Production）

---

## 常用命令（Common Commands）

```powershell
# 格式化代码（Format Code）
gofmt -w .
goimports -w .

# 依赖检查（Dependency Check）
go mod tidy
go mod verify

# 构建（Build）
go build -ldflags="-X main.version=$(git describe --tags)" -o bin/server

# 运行测试（Run Tests）
go test -v ./...

# Docker构建（Docker Build）
docker build -t myapp:latest .
```
