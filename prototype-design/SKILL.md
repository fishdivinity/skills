---
name: "prototype-design"
description: "Professional prototype design skill for B-end enterprise applications. Invoke when user needs to create wireframes, mockups, or UI prototypes. Supports PC and mobile with premium visual design."
---

# 原型设计 Skill

## 核心定位

专注于**快速生成高质量原型**（双击HTML即可预览），而非生产级代码。

- **样式方案**: Tailwind CDN + 内联样式
- **查看方式**: 直接打开HTML文件，无需开发服务器

---

## 1. 主题配色

### 默认主题
| 变量 | 色值 | 用途 |
|------|------|------|
| primary-500 | #0ea5e9 | 主色 |
| neutral-900 | #171717 | 侧边栏 |

### 南方航空主题
| 变量 | 色值 | 用途 |
|------|------|------|
| csa-500 | #0066cc | 主色，按钮/链接 |
| csa-700 | #003d7a | 侧边栏背景 |
| csaRed | #c8102e | 南航红，强调 |

**完整配置**: `templates/design-system.html`

---

## 2. 页面类型与模板

| 页面类型 | 推荐模板 |
|---------|---------|
| 任务管理列表 | `templates/pages/task-management.html` |
| 甘特图/任务编排 | `templates/gantt-advanced.html` |
| 数据看板 | `templates/pages/dashboard.html` |
| 详情页面 | `templates/pages/employee-detail.html` |
| 表单页面 | `templates/components/form-components.html` |

---

## 3. 甘特图专项（⭐核心）

### 核心CSS类
| 元素 | CSS类 | 说明 |
|------|-------|------|
| 停场条 | `.parking-bar` | 蓝色渐变，维修任务 |
| 航班条 | `.flight-bar` | 灰色，航班起降 |
| 任务块 | `.task-block` | 金色，停场内任务 |
| 当前时间线 | `.current-time-line` | 红色竖线 |

### 时间粒度
| 粒度 | 宽度 |
|------|------|
| 分钟 | 15px |
| 小时 | 75px（默认） |
| 天 | 180px |

### 粘性定位
```css
.time-header { position: sticky; top: 0; z-index: 60; }
.fixed-left-column { position: sticky; left: 0; z-index: 40; }
```

### 拖拽实现
```javascript
// 拖拽数据
e.dataTransfer.setData('application/json', JSON.stringify(taskData));

// 放置处理
parkingBar.addEventListener('drop', function(e) {
  if (isOccupied) { showToast('该停场条已被占用', 'error'); return; }
  // 处理放置...
});
```

---

## 4. 组件速查

### 基础组件
- 表单: `templates/components/form-components.html`
- 表格: `templates/components/table-components.html`

### 甘特图组件
- 高级甘特图: `templates/gantt-advanced.html`
- 时间轴: `templates/components/time-axis.html`
- 拖拽交互: `templates/modules/drag-drop.html`

### 交互组件
- 模态框: `templates/components/modals.html`
- 筛选系统: `templates/components/filters.html`
- 状态管理: `templates/modules/state-management.html`

### 布局
- 侧边栏: `templates/layouts/sidebar-layout.html`
- 顶部导航: `templates/layouts/topnav-layout.html`

---

## 5. 工作流程

1. **需求确认** — 参考 `docs/questionnaire.md` 提问确认
2. **分析需求** — 确定页面类型、主题、布局、复杂度
3. **读取模板** — 按页面类型读取对应模板文件
4. **组合调整** — 复制模板，调整字段/数据/颜色
5. **输出文件** — 生成HTML，确保双击可预览

### 复杂度策略
| 复杂度 | 策略 |
|--------|------|
| <300行 | 单文件，简单分区 |
| 300-600行 | 按模块拆分 |
| >600行 | 完整模块化+详细注释 |
| >1000行 | 多文件通过链接关联 |

---

## 6. 文件索引

```
templates/
├── gantt-advanced.html      # 高级甘特图（核心模板）
├── components/
│   ├── time-axis.html      # 时间轴组件
│   ├── modals.html         # 模态框系统
│   └── filters.html        # 筛选系统
├── modules/
│   ├── drag-drop.html      # 拖拽交互
│   └── state-management.html
├── layouts/                # 布局模板
└── pages/                  # 页面模板

docs/
└── questionnaire.md         # 需求确认问题清单
```

---

## 7. 第三方CDN

- Font Awesome 图标
- Flatpickr 日期选择
- Choices.js 下拉多选

CDN配置详见: `templates/design-system.html`
