---
name: "modern-frontend-development"
description: "Modern frontend development skill for production-ready React + TypeScript + Vite + Tailwind CSS projects. Invoke when user needs modern frontend development help, code review, or best practices guidance."
---

# 现代前端开发 Skill

## 核心定位

本技能专注于**现代生产环境前端代码**，目标是创建基于 React + TypeScript + Vite + Tailwind CSS 技术栈的规范、可维护、高性能的前端项目。

### 与其他前端技能的区别

| 特性 | 本技能（现代前端开发） | 前端小Demo技能 | 原型设计技能 |
|------|-------------------|---------------|-------------|
| **目的** | 生产环境代码 | 小Demo开发、快速验证 | 快速可视化、需求沟通 |
| **技术栈** | React + TypeScript + Vite + Tailwind CSS | 纯 HTML/CSS/JS | Tailwind CDN + 内联样式 |
| **代码质量** | 完整、规范、可维护 | 简单、直接、可运行 | 简化、演示性质 |
| **查看方式** | 需要开发服务器 | 双击 HTML 直接打开 | 双击 HTML 直接打开 |

**调用场景：** 现代前端项目开发、代码审查、性能优化、架构设计、企业级应用开发

**不应调用：** 小Demo开发、快速原型绘制、概念可视化演示

---

## 1. 工作流程（重要）

### 步骤1：需求确认

**必须通过反复提问确认所有细节，避免返工。** 按 `docs/questionnaire.md` 逐轮确认。

**确认完成标志：**
- [ ] 项目类型和用途明确
- [ ] 技术栈确定
- [ ] 设计风格确定
- [ ] 所有功能模块明确
- [ ] 用户确认需求理解正确

### 步骤2：复杂度评估

| 复杂度 | 特征 | 处理方式 |
|--------|------|---------|
| 简单 | 单页面应用、功能单一、无后台对接 | 直接开发 |
| 中等 | 多页面应用、有后台对接、基础交互 | 建议使用 `/plan` |
| 复杂 | 多模块、复杂交互、性能要求高 | **必须使用 `/plan`** |
| 超复杂 | 大型项目、多团队协作、长期维护 | 拆分为多个阶段 |

**重要：中等及以上复杂度项目，必须提示用户使用 `/plan` 命令。**

### 步骤3：技术方案确定

- React + TypeScript + Vite：现代前端开发标准配置
- Tailwind CSS：实用优先的CSS框架
- 状态管理：根据项目复杂度选择合适的状态管理方案
- 路由：React Router或其他现代路由方案
- API调用：Axios或Fetch API

---

## 2. 核心规范速查

### 项目结构规范
- 组件化设计：按功能模块组织组件
- 文件命名： PascalCase 用于组件，camelCase 用于工具函数
- 目录结构：清晰的模块化组织

### TypeScript规范
- 严格模式：启用 `strict` 编译选项
- 类型定义：为所有组件和函数提供类型定义
- 接口命名：使用 PascalCase，前缀为 `I`

### React规范
- 函数组件：优先使用函数组件和Hooks
- 状态管理：合理使用 useState、useEffect、useContext 等Hooks
- 组件设计：单一职责原则， props 类型化

### Tailwind CSS规范
- 配置：自定义主题和工具类
- 命名：使用语义化的类名组合
- 组织：合理组织Tailwind类，保持代码整洁

### 性能优化
- 代码分割：使用动态导入减少初始加载时间
- 组件优化：使用 React.memo、useMemo、useCallback 等
- 网络优化：合理使用缓存和请求优化

### 可访问性
- 符合 WCAG 2.1 标准
- 键盘可导航，提供 alt 属性
- 语义化标签使用

---

## 3. 后台对接规范

### API 调用
- 集中管理 API 地址和请求逻辑
- 使用 TypeScript 类型定义 API 响应
- 统一错误处理和 loading 状态管理

### 不同后台技术栈
- **Go 后台**：RESTful API，注意 JSON 序列化特性
- **Python 后台**：Django/Flask，注意 CSRF 保护
- **Node.js 后台**：Express/NestJS，注意异步处理

### 数据交互流程
初始化加载 → 用户操作 → API 请求 → 数据处理 → UI 更新 → 错误处理

---

## 4. 部署规范

### 开发环境
- Vite 开发服务器 + 热重载
- 配置开发环境 API 地址
- ESLint + Prettier 代码质量保证

### 生产环境
- Vite 构建优化
- 静态资源压缩和缓存
- 配置生产环境 API 地址
- HTTPS 安全连接
- CI/CD 集成

---

## 5. 文件索引

```
modern-frontend-development/
├── docs/
│   ├── questionnaire.md    # 需求确认问题清单
│   ├── design-system.md    # 配色方案、设计规范
│   └── solutions.md        # 功能方案（状态管理、路由、API调用）
├── templates/
│   ├── project_structure.md      # 项目结构模板
│   ├── vite.config.ts            # Vite 配置模板
│   ├── tailwind.config.js        # Tailwind 配置模板
│   └── component_templates/      # React 组件模板
└── SKILL.md               # 技能描述文件（本文件）
```

**详细内容请查阅对应文档：**
- 配色方案：`docs/design-system.md`
- 功能方案（状态管理、路由、API调用）：`docs/solutions.md`
- 需求确认问题：`docs/questionnaire.md`