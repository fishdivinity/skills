# 现代前端项目结构模板

## 基本项目结构

```
project-root/
├── public/                 # 静态资源
│   └── favicon.ico
├── src/                    # 源代码
│   ├── assets/             # 图片、字体等资源
│   ├── components/         # 通用组件
│   │   ├── common/         # 公共组件
│   │   │   ├── button/     # 按钮组件
│   │   │   ├── card/       # 卡片组件
│   │   │   └── form/       # 表单组件
│   │   └── layout/         # 布局组件
│   │       ├── Header.tsx  # 头部组件
│   │       ├── Footer.tsx  # 底部组件
│   │       └── Layout.tsx  # 布局组件
│   ├── pages/              # 页面组件
│   │   ├── Home.tsx        # 首页
│   │   ├── About.tsx       # 关于页
│   │   └── Contact.tsx     # 联系页
│   ├── services/           # API 服务
│   │   ├── api.ts          # API 基础配置
│   │   └── userService.ts  # 用户相关 API
│   ├── hooks/              # 自定义 Hooks
│   │   └── useAuth.ts      # 认证相关 Hook
│   ├── context/            # React Context
│   │   └── AuthContext.tsx # 认证上下文
│   ├── utils/              # 工具函数
│   │   ├── format.ts       # 格式化工具
│   │   └── validation.ts   # 验证工具
│   ├── types/              # TypeScript 类型定义
│   │   └── index.ts        # 类型定义文件
│   ├── App.tsx             # 应用根组件
│   ├── main.tsx            # 应用入口
│   └── vite-env.d.ts       # Vite 类型声明
├── .eslintrc.cjs           # ESLint 配置
├── .gitignore              # Git 忽略文件
├── index.html              # HTML 模板
├── package.json            # 项目依赖
├── postcss.config.js       # PostCSS 配置
├── tailwind.config.js      # Tailwind CSS 配置
├── tsconfig.json           # TypeScript 配置
├── tsconfig.node.json      # TypeScript 节点配置
└── vite.config.ts          # Vite 配置
```

## 目录说明

### public/
- 存放静态资源文件，如 favicon.ico、robots.txt 等
- 这些文件会被直接复制到构建输出目录

### src/
- **assets/**: 存放图片、字体、图标等静态资源
- **components/**: 存放可复用的 React 组件
  - **common/**: 通用基础组件，如按钮、卡片、表单等
  - **layout/**: 布局相关组件，如头部、底部、侧边栏等
- **pages/**: 存放页面级组件，每个文件对应一个路由页面
- **services/**: 存放 API 调用相关代码
- **hooks/**: 存放自定义 React Hooks
- **context/**: 存放 React Context 相关代码
- **utils/**: 存放工具函数
- **types/**: 存放 TypeScript 类型定义
- **App.tsx**: 应用根组件，通常包含路由配置
- **main.tsx**: 应用入口，负责渲染根组件
- **vite-env.d.ts**: Vite 类型声明文件

## 配置文件说明

### package.json
- 项目依赖管理
- 脚本命令配置

### tsconfig.json
- TypeScript 编译配置

### vite.config.ts
- Vite 构建工具配置

### tailwind.config.js
- Tailwind CSS 配置

### postcss.config.js
- PostCSS 配置

### .eslintrc.cjs
- ESLint 代码质量检查配置

### .gitignore
- Git 版本控制忽略文件配置

## 项目初始化步骤

1. **创建项目**
   ```bash
   npm create vite@latest project-name -- --template react-ts
   cd project-name
   ```

2. **安装依赖**
   ```bash
   npm install
   npm install tailwindcss postcss autoprefixer
   npx tailwindcss init -p
   ```

3. **配置 Tailwind CSS**
   - 修改 `tailwind.config.js` 文件，添加内容路径
   - 在 `src/index.css` 文件中添加 Tailwind 指令

4. **安装其他依赖**
   ```bash
   # 路由
   npm install react-router-dom
   
   # HTTP 客户端
   npm install axios
   
   # 状态管理（可选）
   npm install @reduxjs/toolkit react-redux
   
   # 表单处理（可选）
   npm install react-hook-form
   
   # 图标库
   npm install @heroicons/react
   ```

5. **创建目录结构**
   - 按照上述项目结构创建相应的目录和文件

6. **配置 ESLint 和 Prettier**
   ```bash
   npm install --save-dev eslint prettier eslint-plugin-react eslint-plugin-react-hooks eslint-plugin-prettier @typescript-eslint/eslint-plugin @typescript-eslint/parser
   ```

7. **启动开发服务器**
   ```bash
   npm run dev
   ```

## 最佳实践

1. **组件设计**
   - 遵循单一职责原则
   - 使用 TypeScript 类型定义 Props
   - 合理使用 React Hooks
   - 组件命名使用 PascalCase

2. **代码风格**
   - 使用 ESLint 和 Prettier 保持代码风格一致
   - 遵循 TypeScript 最佳实践
   - 合理使用注释

3. **性能优化**
   - 使用 React.memo、useMemo、useCallback 优化组件渲染
   - 合理使用代码分割
   - 优化图片和静态资源

4. **可访问性**
   - 符合 WCAG 2.1 标准
   - 提供适当的 alt 属性
   - 确保键盘可导航

5. **测试**
   - 编写单元测试
   - 编写集成测试
   - 使用测试覆盖率工具

## 部署流程

1. **构建项目**
   ```bash
   npm run build
   ```

2. **部署到静态网站托管服务**
   - Vercel
   - Netlify
   - GitHub Pages
   - AWS S3 + CloudFront

3. **配置 CI/CD**
   - GitHub Actions
   - GitLab CI/CD
   - Jenkins

## 常见问题

1. **TypeScript 类型错误**
   - 确保所有组件和函数都有正确的类型定义
   - 使用类型断言时要谨慎

2. **Tailwind CSS 类名过长**
   - 使用 `@apply` 提取复用样式
   - 合理组织类名顺序

3. **性能问题**
   - 使用 React DevTools 分析组件渲染
   - 优化 API 请求
   - 合理使用状态管理

4. **路由配置**
   - 使用 React Router v6 的新特性
   - 合理组织路由结构

5. **API 调用**
   - 集中管理 API 地址
   - 统一错误处理
   - 使用 TypeScript 定义 API 响应类型