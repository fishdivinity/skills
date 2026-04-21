---
name: python-development
description: Python 3.13+ general development skill with PEP8 compliance, strict mode, and best practices. Invoke for general Python development tasks, web applications, data processing scripts, code review, and project structure setup.
---

# Python 开发技能

## 概述
Python 3.13+ 开发指南，用于构建生产级（Production-Ready）应用程序。遵循PEP8规范，使用严格模式，注重代码鲁棒性和解耦。

## 技术栈
- **Python版本**: 3.13.13
- **环境管理**: 虚拟环境（venv）
- **代码质量**: PEP8 + 严格类型检查

## 工作流程
1. **需求确认**: 调用`askuserquestion`，确认项目规模和复杂度
2. **环境准备**: 复杂项目创建虚拟环境
3. **开发执行**: 按最佳实践实施
4. **质量检查**: 运行lint和类型检查

## 核心规范索引
- **编码规范**: `docs/coding-standards.md`
- **安全规范**: `docs/security.md`
- **性能规范**: `docs/performance.md`
- **环境管理**: `docs/environment.md`

## 禁止模式
- 将所有逻辑写在一个Python文件中（非demo项目）
- 不使用虚拟环境管理依赖
- 忽略PEP8规范
- 硬编码敏感信息
- 不使用类型提示
- 捕获异常后不处理

## 代码审查清单
- 遵循PEP8规范
- 使用类型提示
- 函数职责单一
- 代码模块化
- 错误处理完善
- 无硬编码敏感信息
- 虚拟环境配置正确
- 测试覆盖率高

## 资源索引
- **示例代码**: `examples/`
- **模板文件**: `templates/`
- **脚本工具**: `scripts/`
- **命令清单**: `commands.md`

## 文件结构
```
python-development/
├── SKILL.md              # 技能描述文件（本文件）
├── docs/                 # 详细文档
├── examples/             # 示例代码
├── templates/            # 模板文件
├── scripts/              # 脚本工具
└── commands.md           # 常用命令清单
```