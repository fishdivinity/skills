package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// 版本信息（通过ldflags注入）
var (
	version    = "dev"
	commit     = "unknown"
	buildTime  = "unknown"
)

// User 用户模型
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

// Handler 处理器
type Handler struct {
	users []User
}

// NewHandler 创建新的处理器
func NewHandler() *Handler {
	return &Handler{
		users: []User{
			{ID: 1, Name: "Alice", Email: "alice@example.com"},
			{ID: 2, Name: "Bob", Email: "bob@example.com"},
		},
	}
}

// requireAuth 认证中间件
func (h *Handler) requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		// 实际应用中应该验证JWT token
		c.Next()
	}
}

// getUsers 获取用户列表
func (h *Handler) getUsers(c *gin.Context) {
	requestID := c.Request.Header.Get("X-Request-ID")
	if requestID == "" {
		requestID = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	log.Printf("[requestID: %s] GET /api/users", requestID)
	c.JSON(http.StatusOK, h.users)
}

// getHealth 健康检查
func getHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// getReady 就绪检查
func getReady(c *gin.Context) {
	// 检查数据库连接等
	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// 设置Gin模式
	if os.Getenv("GO_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Gin引擎
	r := gin.Default()

	// 健康检查端点
	r.GET("/livez", getHealth)
	r.GET("/readyz", getReady)

	// 创建处理器
	handler := NewHandler()

	// API路由组
	api := r.Group("/api")
	api.Use(handler.requireAuth())
	{
		api.GET("/users", handler.getUsers)
	}

	// 启动服务器
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// 优雅关闭
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
		<-sigCh

		log.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}
	}()

	// 启动服务器
	log.Printf("Server starting on :8080 (version: %s, commit: %s, build: %s)", version, commit, buildTime)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("Server exited gracefully")
}
