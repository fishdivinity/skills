# Go 安全规范

## 身份验证（Authentication）

- 除登录/注册接口外，所有API必须JWT/OAuth2认证

```go
func (h *Handler) requireAuth(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        // 验证token...
        next.ServeHTTP(w, r)
    })
}
```

---

## 数据加密（Data Encryption）

- 敏感数据必须加密存储
- 使用AES-256-GCM等现代加密算法（Modern Encryption Algorithms）
- 后量子加密：Go 1.26+ 可使用 `crypto/hpke` 实现 RFC 9180 标准的混合公钥加密

```go
import "crypto/hpke"

func example() {
    recipient, err := hpke.NewRecipientHPKE(
        hpke.AEADChaCha20Poly1305,
        hpke.KDFX963SHA256,
        hpke.KEMX25519,
        publicKey,
    )
}
```

---

## 传输加密（Transport Encryption）

- TLS 1.3 强制启用

```go
srv := &http.Server{
    TLSConfig: &tls.Config{
        MinVersion: tls.VersionTLS13,
    },
}
```

- Go 1.26+ 默认启用混合后量子密钥交换 `SecP256r1MLKEM768` 和 `SecP384r1MLKEM1024`

---

## 输入验证（Input Validation）

- 限制POST键值对数量
- 限制请求体大小（Request Body Size）
- 限制ZIP解压大小
- 防止SQL注入（SQL Injection）、命令注入（Command Injection）

---

## SQL安全（SQL Security）

- 必须使用参数化查询（Parameterized Query）或ORM

```go
// 参数化查询示例
rows, err := db.QueryContext(ctx,
    "SELECT * FROM users WHERE id = $1", userID)
```

---

## 文件路径（File Path Security）

- 使用 `filepath.Join` 拼接路径，防止路径遍历攻击（Path Traversal Attack）

```go
path := filepath.Join(config.BaseDir, "uploads", filename)
```

---

## 优雅关闭（Graceful Shutdown）

- 必须支持Graceful Shutdown

```go
srv := &http.Server{Addr: ":8080"}
go func() {
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
    <-sigCh

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    srv.Shutdown(ctx)
}()
```

---

## 敏感数据处理（实验性）

- Go 1.26+ 可使用 `runtime/secret` 包安全擦除临时敏感数据
- 启用方式：`GOEXPERIMENT=runtimesecret`
- 支持架构：amd64 和 arm64 (Linux)

```go
import "golang.org/x/exp/runtime/secret"

func processSensitiveData(data []byte) {
    secret.Erase(data)
}
```
