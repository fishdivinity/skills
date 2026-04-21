# Python 开发技能

## 概述

Python 3.13+ 开发技能，遵循PEP8规范，使用严格模式，注重代码鲁棒性和解耦。本技能提供了一套完整的Python开发最佳实践，包括编码规范、环境管理、安全规范和性能优化。

## 目录结构

```
python-development/
├── SKILL.md              # 技能描述文件
├── README.md             # 本文件
├── docs/                 # 详细文档
│   ├── coding-standards.md  # 编码规范
│   ├── environment.md       # 环境管理
│   ├── security.md          # 安全规范
│   └── performance.md       # 性能规范
├── examples/             # 示例代码
│   ├── simple_module.py     # 简单模块示例
│   └── project_structure.md # 项目结构示例
├── scripts/              # 脚本工具
│   └── lint_check.py        # 代码质量检查脚本
├── templates/            # 模板文件
│   ├── pyproject.toml.template  # pyproject.toml模板
│   └── setup.py.template        # setup.py模板
└── commands.md           # 常用命令清单
```

## 核心特性

### 1. 编码规范
- 严格遵循PEP8规范
- 使用类型提示
- 模块化设计
- 清晰的命名规范

### 2. 环境管理
- 虚拟环境创建和管理
- 依赖管理
- 项目结构标准化

### 3. 安全规范
- 输入验证
- 密码处理
- 敏感信息管理
- SQL注入防护
- 依赖安全检查

### 4. 性能优化
- 算法和数据结构选择
- 内存优化
- I/O 优化
- 并行处理
- 性能分析

## 使用指南

### 1. 需求确认
- 调用`askuserquestion`确认项目规模和复杂度
- 确定是否需要创建虚拟环境

### 2. 环境准备
- 复杂项目创建虚拟环境
- 安装必要的依赖
- 配置开发工具

### 3. 开发执行
- 遵循编码规范
- 实现模块化设计
- 编写测试用例
- 进行代码审查

### 4. 质量检查
- 运行 lint 检查
- 运行类型检查
- 运行测试用例
- 进行性能分析

## 最佳实践

### 项目结构
- 使用标准项目结构
- 模块化设计
- 清晰的目录组织

### 代码质量
- 遵循PEP8规范
- 使用类型提示
- 编写文档字符串
- 进行单元测试

### 环境管理
- 始终使用虚拟环境
- 管理依赖版本
- 配置开发工具

### 安全措施
- 验证所有输入
- 安全处理敏感信息
- 定期检查依赖安全
- 实施适当的错误处理

## 工具推荐

### 代码质量工具
- **ruff**: 快速的代码检查和格式化工具
- **mypy**: 静态类型检查工具
- **pytest**: 测试框架
- **black**: 代码格式化工具
- **isort**: 导入排序工具

### 性能分析工具
- **cProfile**: 内置性能分析工具
- **memory_profiler**: 内存使用分析工具
- **line_profiler**: 行级性能分析工具
- **py-spy**: 采样分析器

### 安全工具
- **bandit**: 安全代码检查工具
- **safety**: 依赖安全检查工具

## 示例项目

查看 `examples/` 目录中的示例代码，了解如何组织Python项目和编写符合最佳实践的代码。

## 命令参考

查看 `commands.md` 文件，了解Python开发中常用的命令。

## 文档索引

- **编码规范**: `docs/coding-standards.md`
- **环境管理**: `docs/environment.md`
- **安全规范**: `docs/security.md`
- **性能规范**: `docs/performance.md`

## 许可证

MIT