# Go 开发需求确认问题清单

## 概述

本文档用于 Go 开发需求确认，通过多轮提问确保需求清晰，避免返工。

---

## 第一轮：基础信息确认

### 1. 项目类型与用途
- 这是什么类型的项目？（REST API、gRPC服务、CLI工具、微服务等）
- 主要目标用户是谁？
- 项目的主要目的是什么？

### 2. 技术栈确认
- Go 版本要求？（最低 1.26.0）
- 是否需要对接数据库？使用什么数据库？（PostgreSQL、MySQL、MongoDB等）
- 是否需要缓存？（Redis、Memcached等）
- 是否需要消息队列？（Message Queue: Kafka、RabbitMQ等）

### 3. 部署环境（Deployment Environment）
- 部署方式？（Docker、Kubernetes、裸机 Bare Metal）
- 环境隔离需求？（开发、测试、生产）
- 是否需要 CI/CD 集成？

---

## 第二轮：功能细节确认

### API 服务类
- 需要哪些 API 端点（Endpoints）？
- 认证方式？（JWT、OAuth2、API Key）
- 是否需要限流（Rate Limiting）？
- 是否需要 API 文档？（Swagger/OpenAPI）

### 数据处理类
- 数据量级？
- 是否需要批量处理（Batch Processing）？
- 是否需要异步处理（Async Processing）？
- 数据一致性要求？（Data Consistency: 强一致性/最终一致性）

### CLI 工具类
- 需要哪些命令？
- 是否需要配置文件？
- 输出格式？（JSON、YAML、表格 Table）

---

## 第三轮：安全与性能细节

### 安全要求（Security Requirements）
- 是否处理敏感数据？
- 是否需要加密存储？
- 是否需要审计日志（Audit Log）？
- 是否需要输入验证（Input Validation）？

### 性能要求（Performance Requirements）
- 预期 QPS/TPS？
- 响应时间要求？
- 是否需要连接池（Connection Pool）？
- 是否需要缓存策略（Caching Strategy）？

### 可观测性（Observability）
- 是否需要日志收集（Log Collection）？
- 是否需要监控告警（Monitoring & Alerting）？
- 是否需要链路追踪（Distributed Tracing）？

---

## 确认完成标志

- [ ] 项目类型和用途明确
- [ ] 技术栈确定
- [ ] 部署环境确定
- [ ] 所有功能模块明确
- [ ] 安全要求明确
- [ ] 性能要求明确
- [ ] 可观测性需求明确
- [ ] 用户确认需求理解正确

---

## 复杂度评估

根据确认的需求评估项目复杂度：

| 复杂度 | 特征 | 建议处理方式 |
|--------|------|-------------|
| 简单 | 单功能、无外部依赖、无安全要求 | 直接开发 |
| 中等 | 多功能、有数据库/缓存、基础安全 | 建议使用 `/plan` |
| 复杂 | 微服务架构（Microservices）、多数据源、高并发、严格安全 | **必须使用 `/plan`** |
| 超复杂 | 分布式系统（Distributed System）、多团队协作、长期维护 | 拆分为多个阶段 |

**注意**：中等及以上复杂度的项目，建议用户先使用 `/plan` 命令制定开发计划。
