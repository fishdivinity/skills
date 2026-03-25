# Go 性能与并发规范

## Goroutine管理（Goroutine Management）

- 合理使用goroutine，避免goroutine泄漏（Goroutine Leak）
- 使用 `errgroup` 管理并发任务（Concurrent Tasks）
- 使用 `context` 传递取消信号（Cancellation Signal）

```go
g := &errgroup.Group{}
g.Go(func() error {
    return task1()
})
g.Go(func() error {
    return task2()
})
if err := g.Wait(); err != nil {
    // 处理错误
}
```

---

## 数据库连接池（Database Connection Pool）

```go
db.SetMaxOpenConns(25)
db.SetMaxIdleConns(5)
db.SetConnMaxLifetime(5 * time.Minute)
```

---

## HTTP超时（HTTP Timeout）

```go
client := &http.Client{
    Timeout: 10 * time.Second,
}
```

---

## 缓存策略（Caching Strategy）

- 使用Redis
- 注意缓存一致性（Cache Consistency）
- 实现缓存失效策略（Cache Invalidation Strategy）
