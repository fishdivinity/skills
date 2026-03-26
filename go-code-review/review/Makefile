# Makefile for review tool

.PHONY: build clean test install

# 检测操作系统
ifeq ($(OS),Windows_NT)
    EXECUTABLE=review.exe
else
    EXECUTABLE=review
endif

# 构建工具
build:
	@mkdir -p bin
	go build -o bin/$(EXECUTABLE) ./cmd/review

# 安装到tools目录
install:
	@mkdir -p ../tools
	go build -o ../tools/$(EXECUTABLE) ./cmd/review

# 清理构建产物
clean:
	@rm -f bin/review bin/review.exe ../tools/review ../tools/review.exe 2>/dev/null || true

# 运行测试
test:
	go test ./...

# 安装到系统
tinstall:
	go install ./cmd/review
