# Go 编码规范

## 格式与命名

```bash
gofmt -w .
goimports -w .
```

- 命名采用驼峰（CamelCase）

---

## 错误处理（Error Handling）（必须遵循）

```go
// 必须检查所有错误
if err != nil {
    return fmt.Errorf("failed to do something: %w", err)
}

// Go 1.26: 使用 errors.As 进行类型断言（Type Assertion）
var customErr *CustomError
if errors.As(err, &customErr) {
    // 处理特定错误类型
}

// 使用 %w 包装底层错误（Wrap Error）
return fmt.Errorf("process request %s: %w", req.ID, err)
```

---

## 日志规范（Logging）

- 使用结构化日志（Structured Logging: logrus/zap）
- 所有日志必须包含请求ID（Request ID）

```go
log.WithFields(log.Fields{
    "request_id": requestID,
    "user_id":    userID,
}).Info("processing request")
```

---

## 配置管理（Configuration Management）

- 环境变量优先
- 配置文件（YAML/JSON/TOML）辅助
- **禁止**硬编码（Hardcoding）

---

## Go 1.26+ 新特性

### 已启用特性
- 指针初始化（Pointer Initialization）: `new(expr)` 
- 泛型递归约束（Generic Recursive Constraints）
- 加密（Cryptography）: "无Reader"接口
- 默认启用 Green Tea GC
- 大JSON处理（Large JSON Processing）: `encoding/json/v2` 流式API（Streaming API）
- CGO开销降低30%

### Range Over Int（Go 1.22+）

```go
// 传统写法
for i := 0; i < n; i++ {
    // ...
}

// 推荐写法（Go 1.22+）
for i := range n {
    // i 从 0 到 n-1
}

// 示例：遍历切片
items := []string{"a", "b", "c"}
for i := range len(items) {
    fmt.Println(items[i])
}
```

### 实验性特性（Experimental Features）
- SIMD加速（SIMD Acceleration）
- goroutine泄漏检测（Goroutine Leak Detection）

---

## 禁止模式（Anti-patterns）

**绝对禁止**：
- `panic` 处理业务逻辑
- 全局变量（除配置常量外）
- 复杂的 `init` 函数
- 裸SQL（Raw SQL，必须参数化）
- 硬编码敏感信息
- 循环中使用 `defer`
- 无限制解析ZIP/表单
- 使用旧的 `errors.As`（应使用 `errors.As` 类型断言）
- 返回明文敏感信息
