---
name: "frontend-development"
description: "Frontend development best practices skill for production-ready HTML/CSS/JS projects. Invoke when user needs frontend development help, code review, or best practices guidance. NOT for prototyping or wireframing."
---

# 前端开发最佳实践 Skill

## 核心定位

本技能专注于**生产环境可用的前端代码**，目标是创建规范、可维护、高性能的前端项目。

### 与原型设计技能的区别

| 特性 | 本技能（前端开发） | 原型设计技能 |
|------|-------------------|-------------|
| **目的** | 生产环境代码 | 快速可视化、需求沟通 |
| **样式方案** | BEM 规范 + CSS 变量 | Tailwind CDN + 内联样式 |
| **代码质量** | 完整、规范、可维护 | 简化、演示性质 |
| **查看方式** | 需要开发服务器 | 双击 HTML 直接打开 |

**调用场景：** 生产环境前端代码开发、代码审查、性能优化、架构设计

**不应调用：** 快速原型绘制、概念可视化演示

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
| 简单 | 单页面、功能单一、无后台对接 | 直接开发 |
| 中等 | 多页面、有后台对接、基础交互 | 建议使用 `/plan` |
| 复杂 | 多模块、复杂交互、性能要求高 | **必须使用 `/plan`** |
| 超复杂 | 大型项目、多团队协作、长期维护 | 拆分为多个阶段 |

**重要：中等及以上复杂度项目，必须提示用户使用 `/plan` 命令。**

### 步骤3：技术方案确定

- 纯 HTML/CSS/JS：使用 BEM 规范 + CSS 变量
- Tailwind 项目：配置 purge，本地构建
- 有后台对接：确定 API 规范和数据格式

---

## 2. 核心规范速查

### HTML 规范
- `<!DOCTYPE html>` 开头，指定 `lang` 属性
- 语义化标签：`header`, `nav`, `main`, `section`, `article`, `footer`
- 缩进 2 空格，属性值双引号，避免内联样式

### CSS 规范
- BEM 命名：`block__element--modifier`
- CSS 变量定义通用样式
- **生产环境禁止外部 CDN**（详见 `docs/solutions.md`）

### 命名规范
- 文件：小写字母，连字符分隔（如 `user-profile.html`）
- 类名：语义化，反映功能

### 性能优化
- 减少 HTTP 请求，压缩文件
- 使用适当图片格式，合理缓存

### 可访问性
- 符合 WCAG 2.1 标准
- 键盘可导航，提供 alt 属性

---

## 3. 后台对接规范

### API 调用
- 集中管理 API 地址
- 使用 JSON 格式交互
- 统一错误处理

### 不同后台技术栈
- **Go 后台**：RESTful API，注意 JSON 序列化特性
- **Python 后台**：Django/Flask，注意 CSRF 保护

### 数据交互流程
初始化加载 → 用户操作 → API 请求 → 数据处理 → UI 更新 → 错误处理

---

## 4. 部署规范

### 开发环境
- 本地服务器 + 热重载
- 配置开发环境 API 地址

### 生产环境
- 压缩 CSS/JS，优化图片
- 配置生产环境 API 地址
- HTTPS 安全连接

---

## 5. 文件索引

```
frontend-development/
├── docs/
│   ├── questionnaire.md    # 需求确认问题清单
│   ├── design-system.md    # 配色方案、设计规范
│   └── solutions.md        # 功能方案（头像、图标、CDN规范）
├── templates/
│   ├── project_structure.md      # 项目结构模板
│   └── component_templates/      # 组件模板
└── SKILL.md               # 技能描述文件（本文件）
```

**详细内容请查阅对应文档：**
- 配色方案：`docs/design-system.md`
- 功能方案（头像、图标、CDN）：`docs/solutions.md`
- 需求确认问题：`docs/questionnaire.md`