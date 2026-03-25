# 设计系统

## 1. 主题配色方案

### 1.1 默认现代主题

| 颜色变量 | 色值 | 用途 |
|---------|------|------|
| `--primary-50` | #f0f9ff | 主色浅背景 |
| `--primary-100` | #e0f2fe | 主色背景 |
| `--primary-500` | #0ea5e9 | 主色，按钮、链接 |
| `--primary-600` | #0284c7 | 主色hover状态 |
| `--primary-700` | #0369a1 | 主色active状态 |
| `--neutral-50` | #fafafa | 页面背景 |
| `--neutral-100` | #f5f5f5 | 卡片背景 |
| `--neutral-200` | #e5e5e5 | 边框颜色 |
| `--neutral-700` | #404040 | 正文文字 |
| `--neutral-900` | #171717 | 标题文字 |
| `--success-500` | #22c55e | 成功状态 |
| `--warning-500` | #f59e0b | 警告状态 |
| `--danger-500` | #ef4444 | 危险状态 |
| `--info-500` | #3b82f6 | 信息状态 |

### 1.2 暖色专业主题

适合商务、企业类项目。

| 颜色变量 | 色值 | 用途 |
|---------|------|------|
| `--primary-500` | #f59e0b | 主色 |
| `--primary-600` | #d97706 | hover状态 |
| `--neutral-50` | #fafaf9 | 页面背景 |
| `--neutral-900` | #1c1917 | 标题文字 |

### 1.3 冷色科技主题

适合科技、SaaS类项目。

| 颜色变量 | 色值 | 用途 |
|---------|------|------|
| `--primary-500` | #8b5cf6 | 主色 |
| `--primary-600` | #7c3aed | hover状态 |
| `--neutral-50` | #fafafa | 页面背景 |
| `--neutral-900` | #18181b | 标题文字 |

## 2. CSS 变量配置

```css
:root {
  --primary-50: #f0f9ff;
  --primary-100: #e0f2fe;
  --primary-500: #0ea5e9;
  --primary-600: #0284c7;
  --primary-700: #0369a1;
  
  --neutral-50: #fafafa;
  --neutral-100: #f5f5f5;
  --neutral-200: #e5e5e5;
  --neutral-700: #404040;
  --neutral-900: #171717;
  
  --success-500: #22c55e;
  --warning-500: #f59e0b;
  --danger-500: #ef4444;
  --info-500: #3b82f6;
  
  --font-sans: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'Noto Sans SC', sans-serif;
  --font-mono: 'SF Mono', Monaco, 'Cascadia Code', 'Roboto Mono', Consolas, monospace;
  
  --spacing-xs: 4px;
  --spacing-sm: 8px;
  --spacing-md: 16px;
  --spacing-lg: 24px;
  --spacing-xl: 32px;
  
  --radius-sm: 4px;
  --radius-md: 8px;
  --radius-lg: 12px;
  --radius-xl: 16px;
  
  --shadow-sm: 0 1px 2px 0 rgb(0 0 0 / 0.05);
  --shadow-md: 0 4px 6px -1px rgb(0 0 0 / 0.1);
  --shadow-lg: 0 10px 15px -3px rgb(0 0 0 / 0.1);
}
```

## 3. 设计规范速查

### 颜色使用
- **主色**: `var(--primary-500)` 用于按钮、链接、选中状态
- **背景**: `var(--neutral-50)` 页面背景, `var(--neutral-100)` 卡片背景
- **文字**: `var(--neutral-900)` 标题, `var(--neutral-700)` 正文
- **状态**: `var(--success-500)` 成功, `var(--warning-500)` 警告, `var(--danger-500)` 危险

### 间距规范
- 卡片内边距: `var(--spacing-lg)` (24px)
- 表单项间距: `var(--spacing-md)` (16px)
- 按钮内边距: `var(--spacing-sm) var(--spacing-md)` (8px 16px)

### 圆角规范
- 小组件: `var(--radius-sm)` (4px)
- 卡片: `var(--radius-md)` (8px)
- 模态框: `var(--radius-lg)` (12px)

### 阴影规范
- 卡片: `var(--shadow-sm)`
- 悬浮/下拉: `var(--shadow-md)`
- 模态框: `var(--shadow-lg)`