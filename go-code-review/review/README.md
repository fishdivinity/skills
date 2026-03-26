# Review Tool

一个轻量级的Go代码审查辅助工具，用于检测代码中的优化机会并生成分析报告。

## 功能

- **Map预分配检测**：检测未预分配容量的map创建
- **重复代码模式检测**：检测常见的重复代码模式，如切片遍历检查元素是否存在
- **应用程序规模检测**：基于文件数量自动检测应用程序规模
- **Markdown报告生成**：生成结构化的分析报告
- **基于规模的建议**：根据应用程序规模生成相应的优化建议

## 安装

### 前提条件
- Go 1.26+

### 构建方法

```bash
# 克隆仓库
cd review

# 构建工具
go build -o bin/review.exe ./cmd/review

# 或使用Makefile
make build
```

## 使用

### 分析代码

```bash
# 分析指定目录
review.exe analyze --path /path/to/project

# 分析git diff
review.exe analyze --path /path/to/project --diff

# 覆盖自动检测的规模
review.exe analyze --path /path/to/project --scale medium
```

### 生成报告

```bash
# 生成markdown格式报告
review.exe report --format markdown

# 生成文本格式报告
review.exe report --format text
```

## 示例输出

```
Analyzing path: /path/to/project
Analysis completed successfully!
Analyzed 83 files
Found 16 optimization opportunities
Detected application scale: large

Report:
# Code Review Report

## Summary
- **Files Analyzed**: 83
- **Issues Found**: 16
- **Scale**: large

## Optimization Opportunities

### MapPreallocation
- **Description**: Map created without preallocation
- **File**: /path/to/project/app/services/backup_service.go
- **Line**: 115
- **Severity**: Low

## Recommendations

- Prioritize performance and scalability
- Implement comprehensive monitoring
- Optimize all performance-critical paths
```

## 集成

### 与Go-Code-Review技能集成

该工具已集成到Go-Code-Review技能中，位于 `tools/review.exe`。技能会自动检查工具是否存在，如果不存在则构建它，然后使用它进行代码分析。

### 自定义分析规则

可以通过扩展 `internal/analyzer` 包来添加自定义分析规则：

1. 在 `internal/analyzer` 目录中创建新的分析器文件
2. 实现分析逻辑
3. 在 `analyzer.go` 的 `Analyze` 方法中调用新的分析器

## 项目结构

```
review/
├── bin/              # 构建产物
├── cmd/              # 命令行入口
│   └── review/       # 主命令
├── internal/         # 内部包
│   ├── analyzer/     # 代码分析
│   ├── reporter/     # 报告生成
│   └── priority/     # 优先级校准
├── go.mod            # Go模块定义
├── Makefile          # 构建脚本
└── README.md         # 本文档
```

## 许可证

MIT
