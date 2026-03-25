# 构建阶段
FROM golang:1.26-alpine3.22 AS builder

# 设置工作目录
WORKDIR /app

# 复制go.mod和go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-s -w -X main.version=$(git describe --tags 2>/dev/null || echo 'dev') -X main.commit=$(git rev-parse HEAD 2>/dev/null || echo 'unknown') -X main.buildTime=$(date -u '+%Y-%m-%d_%H:%M:%S')" \
    -o /app/server

# 运行阶段
FROM alpine:3.22

# 安装证书
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /app

# 从构建阶段复制可执行文件
COPY --from=builder /app/server .

# 暴露端口
EXPOSE 8080

# 设置环境变量
ENV GO_ENV=production

# 运行应用
CMD ["./server"]
