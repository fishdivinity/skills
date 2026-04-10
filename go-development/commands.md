# Go 常用命令清单

## 基础命令

### 模块管理
- `go mod init <module>` - 初始化新的 Go 模块
- `go mod tidy` - 整理依赖，添加缺失的模块并移除未使用的模块
- `go mod download` - 下载模块依赖
- `go mod vendor` - 将依赖复制到 vendor 目录
- `go mod edit` - 编辑 go.mod 文件
- `go mod graph` - 打印模块依赖图
- `go mod why <module>` - 解释为什么需要某个模块

### 构建与运行
- `go build` - 构建当前目录下的包
- `go run <file>` - 直接运行 Go 文件
- `go install` - 编译并安装包
- `go clean` - 清理构建产物

### 测试
- `go test` - 运行测试
- `go test -v` - 详细输出测试结果
- `go test -run <pattern>` - 只运行匹配模式的测试
- `go test ./...` - 运行所有子目录的测试

### 工具
- `go fmt` - 格式化代码
- `go vet` - 静态代码分析
- `go fix` - Go 1.26+ 完全重写，用于自动化代码现代化（Modernizers）
- `go doc` - 查看包文档
- `go list` - 列出包信息
- `go env` - 查看 Go 环境变量

## 常用第三方工具

### 代码质量
- `golangci-lint run` - 运行多种 linter
- `gosec` - 安全代码扫描
- `staticcheck` - 静态代码分析

### 代码生成
- `swag init` - 生成 Swagger 文档
- `mockgen` - 生成 mock 代码
- `stringer` - 为枚举类型生成 String() 方法

### 性能分析
- `go test -bench=.` - 运行基准测试
- `go tool pprof` - 性能分析工具

### 依赖管理
- `go get -u <package>` - 更新指定包
- `go get -u all` - 更新所有依赖
- `go get <package>@<version>` - 获取指定版本的包

### 其他常用工具
- `shadow` - 检测变量 shadowing
- `goplay` - 在线 Go  playground 客户端
- `gore` - Go 交互式 REPL

## 项目相关命令

### 初始化项目
- `mkdir <project> && cd <project> && go mod init <module>` - 创建并初始化新项目

### 依赖管理最佳实践
- `go mod tidy` - 保持依赖整洁
- `go list -m all` - 查看所有依赖版本
- `go mod verify` - 验证依赖完整性

### 构建优化
- `go build -ldflags "-s -w"` - 减少可执行文件大小
- `go build -race` - 启用数据竞争检测

## 常见问题解决

### 依赖冲突
- `go mod tidy` - 自动解决依赖冲突
- `go get <package>@latest` - 强制更新到最新版本

### 模块路径问题
- 确保 GOPROXY 设置正确：`export GOPROXY=https://goproxy.cn,direct`
- 检查 go.mod 文件中的模块路径是否正确

### 版本管理
- 使用 `go get <package>@v1.2.3` 指定具体版本
- 使用 `go get <package>@master` 获取最新开发版本

## 快捷键与技巧

- `go doc <package>` - 查看包文档
- `go doc <package>.<function>` - 查看具体函数文档
- `go list ./... | grep -E "(test|spec)"` - 查找测试文件
- `go build -o <output>` - 指定输出文件名

## 环境变量

- `GOPATH` - Go 工作目录
- `GOROOT` - Go 安装目录
- `GOPROXY` - Go 模块代理
- `GO111MODULE` - 模块模式开关（auto/on/off）

## 参考资源

- [Go 官方文档](https://golang.org/doc/)
- [Go 命令文档](https://golang.org/cmd/go/)
- [Go 模块文档](https://golang.org/ref/mod/)