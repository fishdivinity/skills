---
name: "go-development"
description: "Go 1.26+ development skill with Docker, APIs, security, and best practices. Invoke when user asks for Go development help or code review."
---

# Go 开发技能

Go 1.26+ 开发指南，用于构建生产级（Production-Ready）应用程序。

## 技术栈要求

- **Go版本**: 1.26.0+
- **部署**: 容器化（Docker/Podman）
- **环境**: Windows PowerShell

---

## 0. 工作流程（重要）

**当用户提出 Go 开发需求时，按以下步骤执行：**

### 步骤1：需求确认（反复提问）

**重要：在开始开发之前，必须通过反复提问确认所有细节，避免返工。**

按 `docs/questionnaire.md` 中的问题清单逐轮确认：

| 轮次 | 确认内容 | 问题数量 |
|------|---------|---------|
| 第一轮 | 基础信息（项目类型、技术栈、部署环境） | 3-4个 |
| 第二轮 | 功能细节（根据项目类型针对性提问） | 3-5个 |
| 第三轮 | 安全/性能/可观测性（Observability）细节（如需要） | 2-3个 |

**确认完成标志：**
- [ ] 项目类型和用途明确
- [ ] 技术栈确定
- [ ] 部署环境确定
- [ ] 所有功能模块明确
- [ ] 用户确认需求理解正确

---

### 步骤2：复杂度评估

| 复杂度 | 特征 | 处理方式 |
|--------|------|---------|
| 简单 | 单功能、无外部依赖、无安全要求 | 直接开发 |
| 中等 | 多功能、有数据库/缓存、基础安全 | 建议使用 `/plan` |
| 复杂 | 微服务架构（Microservices）、多数据源、高并发、严格安全 | **必须使用 `/plan`** |
| 超复杂 | 分布式系统（Distributed System）、多团队协作、长期维护 | 拆分为多个阶段 |

**重要：中等及以上复杂度项目，必须提示用户使用 `/plan` 命令。**

---

### 步骤3：开发执行

根据确认的需求和复杂度评估结果：
- **简单项目**：直接编写代码
- **中等及以上**：先制定计划，再按计划执行

---

## 1. 架构规范

### API网关（API Gateway）
- 统一路由入口，实施限流（Rate Limiting）、请求大小限制
- 身份验证（JWT/OAuth2）

### 本地优先（Local-First）
- 支持离线操作，冲突解决机制（版本向量/CRDT）

---

## 2. 依赖管理（Dependency Management）

- 使用 Go Modules，明确版本号
- **禁止** `go get` 未版本控制的依赖
- 确保CI/CD中依赖来源可信

---

## 3. 核心规范速查

详细规范请查阅对应文档：

| 规范类型 | 文档 | 核心要点 |
|---------|------|---------|
| 编码规范 | `docs/coding-standards.md` | gofmt、错误处理（Error Handling）、日志、Go 1.26+特性 |
| 安全规范 | `docs/security.md` | JWT/OAuth2、TLS 1.3、参数化查询（Parameterized Query）、优雅关闭（Graceful Shutdown） |
| 性能规范 | `docs/performance.md` | errgroup、连接池（Connection Pool）、HTTP超时、缓存（Caching） |
| 部署规范 | `docs/deployment.md` | Docker多阶段构建（Multi-stage Build）、健康检查（Health Check）、版本注入 |

---

## 4. 禁止模式（Anti-patterns）

**绝对禁止**：
- `panic` 处理业务逻辑
- 全局变量（除配置常量外）
- 裸SQL（Raw SQL，必须参数化）
- 硬编码敏感信息
- 循环中使用 `defer`
- 返回明文敏感信息

---

## 5. 代码审查清单（Code Review Checklist）

- [ ] 所有错误都已检查并适当处理
- [ ] 使用 `errors.As` 进行错误类型断言
- [ ] 日志包含请求ID（Request ID）
- [ ] 无硬编码敏感信息
- [ ] 所有API（除登录/注册）都有认证
- [ ] 使用参数化查询（Parameterized Query）
- [ ] 实现优雅关闭（Graceful Shutdown）
- [ ] 实现健康检查端点（Health Check Endpoints: /livez、/readyz）
- [ ] 使用 `gofmt`+`goimports` 格式化

---

## 文件索引

```
go-development/
├── SKILL.md              # 技能描述文件（本文件）
├── docs/
│   ├── questionnaire.md  # 需求确认问题清单
│   ├── coding-standards.md # 编码规范
│   ├── security.md       # 安全规范
│   ├── performance.md    # 性能规范
│   └── deployment.md     # 部署规范
├── examples/             # 示例代码
│   ├── main.go           # 完整示例应用
│   ├── Dockerfile        # 多阶段构建配置
│   ├── go.mod            # 依赖管理
│   └── .env.example      # 环境变量示例
└── templates/            # 模板文件
    └── project_structure.md  # 标准项目结构
```
