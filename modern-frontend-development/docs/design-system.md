# 现代前端设计系统

## 1. 主题配色方案

### 1.1 默认现代主题

| 颜色变量 | 色值 | 用途 | Tailwind 对应 |
|---------|------|------|--------------|
| `--primary-50` | #f0f9ff | 主色浅背景 | `bg-primary-50` |
| `--primary-100` | #e0f2fe | 主色背景 | `bg-primary-100` |
| `--primary-500` | #0ea5e9 | 主色，按钮、链接 | `bg-primary-500` |
| `--primary-600` | #0284c7 | 主色hover状态 | `bg-primary-600` |
| `--primary-700` | #0369a1 | 主色active状态 | `bg-primary-700` |
| `--neutral-50` | #fafafa | 页面背景 | `bg-neutral-50` |
| `--neutral-100` | #f5f5f5 | 卡片背景 | `bg-neutral-100` |
| `--neutral-200` | #e5e5e5 | 边框颜色 | `border-neutral-200` |
| `--neutral-700` | #404040 | 正文文字 | `text-neutral-700` |
| `--neutral-900` | #171717 | 标题文字 | `text-neutral-900` |
| `--success-500` | #22c55e | 成功状态 | `bg-success-500` |
| `--warning-500` | #f59e0b | 警告状态 | `bg-warning-500` |
| `--danger-500` | #ef4444 | 危险状态 | `bg-danger-500` |
| `--info-500` | #3b82f6 | 信息状态 | `bg-info-500` |

### 1.2 暖色专业主题

适合商务、企业类项目。

| 颜色变量 | 色值 | 用途 | Tailwind 对应 |
|---------|------|------|--------------|
| `--primary-500` | #f59e0b | 主色 | `bg-primary-500` |
| `--primary-600` | #d97706 | hover状态 | `bg-primary-600` |
| `--neutral-50` | #fafaf9 | 页面背景 | `bg-neutral-50` |
| `--neutral-900` | #1c1917 | 标题文字 | `text-neutral-900` |

### 1.3 冷色科技主题

适合科技、SaaS类项目。

| 颜色变量 | 色值 | 用途 | Tailwind 对应 |
|---------|------|------|--------------|
| `--primary-500` | #8b5cf6 | 主色 | `bg-primary-500` |
| `--primary-600` | #7c3aed | hover状态 | `bg-primary-600` |
| `--neutral-50` | #fafafa | 页面背景 | `bg-neutral-50` |
| `--neutral-900` | #18181b | 标题文字 | `text-neutral-900` |

## 2. Tailwind CSS 配置

```js
// tailwind.config.js
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: {
          50: '#f0f9ff',
          100: '#e0f2fe',
          500: '#0ea5e9',
          600: '#0284c7',
          700: '#0369a1',
        },
        neutral: {
          50: '#fafafa',
          100: '#f5f5f5',
          200: '#e5e5e5',
          700: '#404040',
          900: '#171717',
        },
        success: {
          500: '#22c55e',
        },
        warning: {
          500: '#f59e0b',
        },
        danger: {
          500: '#ef4444',
        },
        info: {
          500: '#3b82f6',
        },
      },
      fontFamily: {
        sans: ['-apple-system', 'BlinkMacSystemFont', 'Segoe UI', 'Roboto', 'Helvetica Neue', 'Arial', 'Noto Sans SC', 'sans-serif'],
        mono: ['SF Mono', 'Monaco', 'Cascadia Code', 'Roboto Mono', 'Consolas', 'monospace'],
      },
      spacing: {
        'xs': '0.25rem', // 4px
        'sm': '0.5rem',  // 8px
        'md': '1rem',    // 16px
        'lg': '1.5rem',  // 24px
        'xl': '2rem',    // 32px
      },
      borderRadius: {
        'sm': '0.25rem', // 4px
        'md': '0.5rem',  // 8px
        'lg': '0.75rem', // 12px
        'xl': '1rem',    // 16px
      },
      boxShadow: {
        'sm': '0 1px 2px 0 rgb(0 0 0 / 0.05)',
        'md': '0 4px 6px -1px rgb(0 0 0 / 0.1)',
        'lg': '0 10px 15px -3px rgb(0 0 0 / 0.1)',
      },
    },
  },
  plugins: [],
}
```

## 3. 设计规范速查

### 颜色使用
- **主色**: `bg-primary-500` 用于按钮、链接、选中状态
- **背景**: `bg-neutral-50` 页面背景, `bg-neutral-100` 卡片背景
- **文字**: `text-neutral-900` 标题, `text-neutral-700` 正文
- **状态**: `bg-success-500` 成功, `bg-warning-500` 警告, `bg-danger-500` 危险

### 间距规范
- 卡片内边距: `p-lg` (24px)
- 表单项间距: `space-y-md` (16px)
- 按钮内边距: `px-sm py-sm` (8px 16px)

### 圆角规范
- 小组件: `rounded-sm` (4px)
- 卡片: `rounded-md` (8px)
- 模态框: `rounded-lg` (12px)

### 阴影规范
- 卡片: `shadow-sm`
- 悬浮/下拉: `shadow-md`
- 模态框: `shadow-lg`

### 响应式设计
- 移动端: `sm:` 前缀
- 平板: `md:` 前缀
- 桌面: `lg:` 前缀
- 大屏: `xl:` 前缀

## 4. React 组件设计规范

### 组件命名
- 组件名: PascalCase (如 `Button`, `UserProfile`)
- 文件命名: 与组件名一致 (如 `Button.tsx`)
- 目录组织: 按功能模块组织

### 组件结构
- 单一职责: 每个组件只负责一个功能
- Props 类型化: 使用 TypeScript 定义 Props 类型
- 状态管理: 合理使用 Hooks 管理状态

### 样式组织
- 使用 Tailwind CSS 类
- 复杂样式: 可使用 `@apply` 提取复用样式
- 主题变量: 使用 Tailwind 配置的主题变量

### 组件示例

```tsx
// Button.tsx
import React from 'react';

interface ButtonProps {
  children: React.ReactNode;
  variant?: 'primary' | 'secondary' | 'danger';
  size?: 'sm' | 'md' | 'lg';
  disabled?: boolean;
  onClick?: () => void;
}

const Button: React.FC<ButtonProps> = ({
  children,
  variant = 'primary',
  size = 'md',
  disabled = false,
  onClick,
}) => {
  const variantClasses = {
    primary: 'bg-primary-500 hover:bg-primary-600 text-white',
    secondary: 'bg-neutral-100 hover:bg-neutral-200 text-neutral-900',
    danger: 'bg-danger-500 hover:bg-danger-600 text-white',
  };

  const sizeClasses = {
    sm: 'px-sm py-xs text-sm',
    md: 'px-md py-sm',
    lg: 'px-lg py-md text-lg',
  };

  return (
    <button
      className={`
        ${variantClasses[variant]}
        ${sizeClasses[size]}
        rounded-md
        font-medium
        transition-colors
        focus:outline-none
        focus:ring-2
        focus:ring-primary-500
        focus:ring-offset-2
        ${disabled ? 'opacity-50 cursor-not-allowed' : ''}
      `}
      disabled={disabled}
      onClick={onClick}
    >
      {children}
    </button>
  );
};

export default Button;
```