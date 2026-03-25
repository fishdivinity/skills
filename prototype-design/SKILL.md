---
name: "prototype-design"
description: "Professional prototype design skill for B-end enterprise applications. Invoke when user needs to create wireframes, mockups, or UI prototypes. Supports PC and mobile with premium visual design."
---

# 原型设计 Skill

## 核心定位

本技能专注于**快速生成高质量原型**，而非生产级代码。目标是创建视觉精美、可直接双击查看的 HTML 原型文件。

### 与前端开发技能的区别

| 特性 | 本技能（原型设计） | 前端开发技能 |
|------|-------------------|-------------|
| **目的** | 快速可视化、需求沟通 | 生产环境代码 |
| **样式方案** | Tailwind CDN + 内联样式 | 构建工具 + CSS 变量 |
| **查看方式** | 双击 HTML 直接打开 | 需要开发服务器 |

---

## 1. 主题配色方案

### 1.1 默认主题（通用B端）

| 颜色变量 | 主色值 | 用途 |
|---------|--------|------|
| `primary-500` | #0ea5e9 | 主色，按钮、链接 |
| `primary-600` | #0284c7 | 主色hover状态 |
| `neutral-900` | #171717 | 侧边栏背景 |

### 1.2 南方航空主题（机务维修系统）

| 颜色变量 | 色值 | 用途 |
|---------|------|------|
| `csa-500` | #0066cc | 主色，按钮、链接、选中状态 |
| `csa-600` | #0052a3 | 主色hover状态 |
| `csa-700` | #003d7a | 侧边栏背景 |
| `csaRed` | #c8102e | 南航红，强调、重要操作 |
| `csaGold` | #d4a843 | 南航金，特殊标记 |

**完整配置代码见**: `templates/design-system.html`

---

## 2. 基础模板结构

每个原型 HTML 文件必须包含：Tailwind CDN、Inter字体、主题配置。

**完整模板见**: `templates/design-system.html`

---

## 3. 设计规范速查

### 颜色使用
- **主色**: `primary-500/600` 或 `csa-500/600` 用于按钮、链接、选中状态
- **背景**: `bg-white` 卡片, `bg-neutral-100` 页面背景, `bg-neutral-900` 或 `bg-csa-700` 侧边栏
- **文字**: `text-neutral-900` 标题, `text-neutral-600` 正文, `text-neutral-500` 辅助
- **状态**: `success` 成功, `warning` 警告, `danger` 危险, `info` 信息

### 间距规范
- 卡片内边距: `p-6` (24px)
- 表单项间距: `space-y-6`
- 按钮内边距: `px-4 py-2` 或 `px-6 py-2.5`

### 圆角规范
- 小组件: `rounded` (4px)
- 卡片: `rounded-lg` (8px) 或 `rounded-xl` (12px)
- 模态框: `rounded-xl` (12px)

### 阴影规范
- 卡片: `shadow` 或 `shadow-sm`
- 悬浮/下拉: `shadow-lg`
- 模态框: `shadow-xl`

---

## 4. 页面类型与模板映射

根据需求快速定位模板：

| 页面类型 | 特征 | 推荐模板 |
|---------|------|---------|
| **任务管理列表** | 搜索筛选+数据表格+批量操作 | `templates/pages/task-management.html` |
| **甘特图/任务编排** | 左右分栏+时间轴+拖拽排程 | `templates/pages/gantt-scheduling.html` |
| **数据看板** | 统计卡片+图表+快捷操作 | `templates/pages/dashboard.html` |
| **详情页面** | 头部信息+分组信息+操作区 | `templates/pages/employee-detail.html` |
| **表单页面** | 表单字段+提交按钮 | `templates/components/form-components.html` |

---

## 5. 布局类型

### 侧边栏布局 (B端管理系统)
- 侧边栏: `w-60 bg-neutral-900` 或 `w-60 bg-csa-700` (南航主题)
- 主内容: `flex-1 ml-60`
- 顶部栏: `h-16 bg-white shadow-sm sticky top-0`

**完整模板**: `templates/layouts/sidebar-layout.html`

### 顶部导航布局 (轻量级系统)
- 顶部导航: `bg-white border-b sticky top-0`
- 内容区: `max-w-7xl mx-auto py-6 px-4`

**完整模板**: `templates/layouts/topnav-layout.html`

---

## 6. 常用组件速查

| 组件 | 说明 | 模板文件 |
|------|------|---------|
| 表单 | 输入框、下拉、日期、开关等 | `templates/components/form-components.html` |
| 表格 | 数据表格、分页、工具栏 | `templates/components/table-components.html` |
| 高级组件 | 下拉多选、级联、树形、富文本 | `templates/components/advanced-components.html` |

---

## 7. 移动端原型

移动端原型使用手机壳容器，在 PC 端展示移动端效果。

**完整模板**: `templates/mobile/phone-frame.html`

---

## 8. 第三方库 CDN

如需高级功能：Font Awesome图标、Flatpickr日期选择、Choices.js下拉多选、Quill富文本。

**完整CDN引用见**: `templates/design-system.html`

---

## 9. 工作流程（重要）

**当用户提出原型需求时，按以下步骤执行：**

### 步骤1：需求确认（反复提问）

**重要：在生成原型之前，必须通过反复提问确认所有细节，避免返工。**

按 `docs/questionnaire.md` 中的问题清单逐轮确认：

| 轮次 | 确认内容 | 问题数量 |
|------|---------|---------|
| 第一轮 | 基础信息（页面类型、主题、布局、用户） | 3-4个 |
| 第二轮 | 功能细节（根据页面类型针对性提问） | 3-5个 |
| 第三轮 | 交互/数据/视觉细节（如需要） | 2-3个 |

**确认完成标志：**
- [ ] 页面类型和用途明确
- [ ] 主题和布局确定
- [ ] 所有字段/列明确
- [ ] 主要交互方式确定
- [ ] 特殊效果/动画要求明确
- [ ] 用户确认需求理解正确

---

### 步骤2：分析需求
- 确定页面类型（列表页、详情页、表单页、看板页）
- 确定主题（默认主题或南航主题）
- 确定布局（侧边栏或顶部导航）
- **评估复杂度**（简单/中等/复杂/超复杂）

### 步骤3：读取模板
根据页面类型，**必须读取对应的模板文件**：
- 任务管理列表 → 读取 `templates/pages/task-management.html`
- 甘特图/任务编排 → 读取 `templates/pages/gantt-scheduling.html`
- 数据看板 → 读取 `templates/pages/dashboard.html`
- 详情页面 → 读取 `templates/pages/employee-detail.html`
- 表单页面 → 读取 `templates/components/form-components.html`
- **复杂原型参考** → 读取 `templates/modular-example.html`

### 步骤4：组合调整
- 复制模板中的基础结构
- 根据需求调整字段、数据
- 应用正确的主题色
- **复杂原型按模块化策略组织代码**

### 步骤5：输出文件
- 生成完整的 HTML 文件
- 保存到指定位置
- **确保双击可预览**

### 复杂原型生成策略

| 复杂度 | 生成策略 |
|--------|---------|
| 简单 | 直接输出完整HTML |
| 中等 | 按模块分区注释，便于维护 |
| 复杂 | 完整模块化拆分（CSS/HTML/JS分离） |
| 超复杂 | 拆分为多个HTML文件，通过链接关联 |

---

## 10. 复杂原型模块化设计

### 10.1 设计原则

| 特性 | 说明 |
|------|------|
| **单文件输出** | 最终输出仍是单个HTML，双击即可预览 |
| **模块化组织** | 内部按功能拆分为多个标记清晰的区块 |
| **按需组合** | AI根据需求选择所需模块组合 |

### 10.2 模块拆分规范

**CSS模块**：按功能拆分为多个`<style>`块（基础样式、布局样式、组件样式、动画样式）

**HTML模块**：使用注释标记区块
```html
<!-- ==================== 侧边栏模块 ==================== -->
<!-- ========== 表格模块 ========== -->
```

**JavaScript模块**：按功能拆分为多个`<script>`块（数据定义、工具函数、渲染函数、事件处理、初始化）

### 10.3 复杂度分级

| 复杂度 | 特征 | 模块化建议 |
|--------|------|-----------|
| 简单 | <300行 | 单文件，简单注释分区 |
| 中等 | 300-600行 | 按模块拆分CSS/HTML/JS |
| 复杂 | >600行 | 完整模块化+详细注释标记 |
| 超复杂 | >1000行 | 拆分为多个HTML文件，通过链接关联 |

**完整示例见**: `templates/modular-example.html`

---

## 11. 文件索引

```
templates/
├── layouts/           # 布局模板
├── components/        # 组件模板
├── pages/             # 页面模板
├── mobile/            # 移动端模板
├── design-system.html # 设计系统（主题配置、CDN引用）
└── modular-example.html # 模块化示例

docs/
└── questionnaire.md   # 需求确认问题清单
```
