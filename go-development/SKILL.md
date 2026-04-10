---
name: go-development
description: Go 1.26+ development skill with Docker, APIs, security, and best practices. Invoke when user asks for Go development help or code review.
---

# Go 开发技能

## 概述
Go 1.26+ 开发指南，用于构建生产级（Production-Ready）应用程序。

## 技术栈
- **Go版本**: 1.26.2
- **部署**: 容器化（Docker/Podman）
- **环境**: Windows PowerShell

## 工作流程
1. **需求确认**: 调用`askuserquestion`，按 `docs/questionnaire.md` 逐轮确认
2. **复杂度评估**: 简单项目直接开发，中等及以上使用 `/plan`
3. **开发执行**: 按计划实施

## 核心规范索引
- **编码规范**: `docs/coding-standards.md`
- **安全规范**: `docs/security.md`
- **性能规范**: `docs/performance.md`
- **部署规范**: `docs/deployment.md`

## 禁止模式
- `panic` 处理业务逻辑
- 全局变量（除配置常量外）
- 裸SQL（Raw SQL，必须参数化）
- 硬编码敏感信息
- 循环中使用 `defer`
- 返回明文敏感信息

## 代码审查清单
- 所有错误都已检查并适当处理
- 使用 `errors.As` 进行错误类型断言
- 日志包含请求ID（Request ID）
- 无硬编码敏感信息
- 所有API（除登录/注册）都有认证
- 使用参数化查询（Parameterized Query）
- 实现优雅关闭（Graceful Shutdown）
- 实现健康检查端点
- 使用 `gofmt`+`goimports` 格式化

## 资源索引
- **示例代码**: `examples/`
- **模板文件**: `templates/`
- **命令清单**: `commands.md`

## 文件结构
```
go-development/
├── SKILL.md              # 技能描述文件（本文件）
├── docs/                 # 详细文档
├── examples/             # 示例代码
├── templates/            # 模板文件
└── commands.md           # 常用命令清单
```
