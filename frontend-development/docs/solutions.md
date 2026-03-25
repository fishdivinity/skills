# 功能方案

## 1. 用户头像方案

### 1.1 推荐方案：后端存储 MD5，前端生成 Identicon

**工作流程：**
1. 后端根据用户信息（如用户ID、邮箱）生成 MD5 哈希值
2. 后端存储该 MD5 值，返回给前端
3. 前端根据 MD5 值动态生成 Identicon 头像

**方案优势：**
- 后端只需存储简短的 MD5 字符串，节省存储空间
- 前端动态生成，减轻后端图片处理压力
- 同一用户每次生成相同头像，保证一致性
- 无需存储图片文件

### 1.2 后端生成 MD5

**Go 实现：**
```go
import (
    "crypto/md5"
    "encoding/hex"
)

func GenerateAvatarMD5(userID string) string {
    hash := md5.Sum([]byte(userID))
    return hex.EncodeToString(hash[:])
}
```

**Python 实现：**
```python
import hashlib

def generate_avatar_md5(user_id: str) -> str:
    return hashlib.md5(user_id.encode()).hexdigest()
```

### 1.3 前端生成 Identicon

**推荐库：**
- `jdenticon` - 轻量级，支持多种风格
- `identicon.js` - 简单易用

**使用 jdenticon：**
```html
<script src="/static/js/jdenticon.min.js"></script>
```

```html
<canvas data-jdenticon-value="用户MD5值" width="40" height="40"></canvas>
```

或使用 SVG：
```html
<svg data-jdenticon-value="用户MD5值" width="40" height="40"></svg>
```

**使用 identicon.js：**
```javascript
import Identicon from 'identicon.js';

function generateAvatar(md5Hash, size = 40) {
    const data = new Identicon(md5Hash, {
        size: size,
        format: 'svg'
    });
    return `data:image/svg+xml;base64,${data.toString()}`;
}

// 使用
const avatarUrl = generateAvatar(user.avatarMd5);
document.querySelector('.avatar img').src = avatarUrl;
```

### 1.4 头像显示规范

```html
<div class="avatar">
  <img 
    id="avatar-{userId}" 
    alt="用户头像" 
    class="avatar__image"
  >
</div>
```

```javascript
// 初始化头像
function initAvatar(userId, avatarMd5, hasCustomAvatar) {
    const img = document.getElementById(`avatar-${userId}`);
    
    if (hasCustomAvatar) {
        img.src = `/api/avatar/${userId}`;
        img.onerror = () => generateIdenticon(img, avatarMd5);
    } else {
        generateIdenticon(img, avatarMd5);
    }
}

function generateIdenticon(img, md5Hash) {
    img.src = `data:image/svg+xml;base64,${new Identicon(md5Hash, {size: 40}).toString()}`;
}
```

```css
.avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  overflow: hidden;
  background-color: var(--neutral-200);
}

.avatar__image {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.avatar--sm { width: 32px; height: 32px; }
.avatar--md { width: 40px; height: 40px; }
.avatar--lg { width: 56px; height: 56px; }
.avatar--xl { width: 80px; height: 80px; }
```

### 1.5 API 返回格式

```json
{
  "userId": "12345",
  "username": "张三",
  "avatarMd5": "e99a18c428cb38d5f260853678922e03",
  "hasCustomAvatar": false
}
```

---

## 2. SVG 图标管理方案

### 2.1 方案对比

| 方案 | 优点 | 缺点 | 适用场景 |
|------|------|------|---------|
| **SVG Sprite** | HTTP 请求少、可缓存、支持多色 | 需要构建工具 | 推荐方案 |
| Icon Font | 兼容性好、文件小 | 只支持单色、存在渲染问题 | 旧项目兼容 |
| 内联 SVG | 无额外请求 | HTML 体积大、难以维护 | 简单项目 |
| JS 动态加载 | 按需加载 | 首屏闪烁 | 大型项目 |

### 2.2 推荐方案：SVG Sprite + CSS 类名

**目录结构：**
```
assets/
├── icons/
│   ├── svg/              # 原始 SVG 文件
│   │   ├── home.svg
│   │   ├── user.svg
│   │   └── settings.svg
│   └── sprite.svg        # 合并后的 Sprite 文件
└── fonts/                # 字体文件（如需要）
```

**单色图标处理：**

1. 使用构建工具将 SVG 转换为 Sprite：
```bash
npx svg-sprite-generator -d assets/icons/svg -o assets/icons/sprite.svg
```

2. CSS 样式定义：
```css
.icon {
  display: inline-block;
  width: 1em;
  height: 1em;
  fill: currentColor;
  vertical-align: middle;
}

.icon--sm { font-size: 16px; }
.icon--md { font-size: 24px; }
.icon--lg { font-size: 32px; }
```

3. HTML 使用：
```html
<svg class="icon icon--md" aria-hidden="true">
  <use href="#icon-home"></use>
</svg>
```

**多色图标处理：**

多色图标推荐使用 **内联 SVG** 或 **JS 动态加载**：

```javascript
const iconCache = new Map();

async function loadIcon(name) {
  if (iconCache.has(name)) return iconCache.get(name);
  
  const response = await fetch(`/assets/icons/svg/${name}.svg`);
  const svgText = await response.text();
  iconCache.set(name, svgText);
  return svgText;
}
```

### 2.3 图标命名规范

| 类型 | 命名前缀 | 示例 |
|------|---------|------|
| 操作图标 | `icon-` | `icon-edit`, `icon-delete` |
| 状态图标 | `status-` | `status-success`, `status-error` |
| 导航图标 | `nav-` | `nav-home`, `nav-settings` |
| 品牌图标 | `brand-` | `brand-wechat`, `brand-github` |

---

## 3. 图片资源管理

### 3.1 图片格式选择

| 格式 | 适用场景 | 压缩工具 |
|------|---------|---------|
| WebP | 照片、复杂图形 | `cwebp` |
| PNG | 透明图片、简单图形 | `pngquant`, `optipng` |
| SVG | 图标、矢量图形 | `svgo` |
| JPEG | 照片（不支持透明） | `jpegoptim` |

### 3.2 响应式图片

```html
<picture>
  <source srcset="image.webp" type="image/webp">
  <source srcset="image.jpg" type="image/jpeg">
  <img src="image.jpg" alt="描述" loading="lazy">
</picture>
```

### 3.3 懒加载

```javascript
const observer = new IntersectionObserver((entries) => {
  entries.forEach(entry => {
    if (entry.isIntersecting) {
      const img = entry.target;
      img.src = img.dataset.src;
      observer.unobserve(img);
    }
  });
});

document.querySelectorAll('img[data-src]').forEach(img => observer.observe(img));
```

---

## 4. CSS 框架使用规范

### 4.1 重要：生产环境禁止使用外部 CDN

生产环境**严禁**使用任何外部 CDN，原因：
- 内网环境可能无法访问外网
- 外部 CDN 可能不稳定，影响用户体验
- 存在安全风险和隐私问题
- 无法保证版本一致性

### 4.2 Tailwind CSS 生产环境方案

| 环境 | 使用方式 | 说明 |
|------|---------|------|
| 开发环境 | Tailwind CLI 或构建工具 | 本地构建，生成 CSS 文件 |
| 生产环境 | **必须本地构建** | 使用 `tailwindcss --minify` 生成压缩 CSS |

**生产环境配置步骤：**
1. 安装：`npm install -D tailwindcss`
2. 初始化：`npx tailwindcss init`
3. 配置 `purge` 选项移除未使用的样式
4. 构建：`npx tailwindcss -i ./src/input.css -o ./dist/output.css --minify`
5. 将生成的 CSS 文件部署到服务器

**禁止事项：**
- ❌ 禁止使用 `<script src="https://cdn.tailwindcss.com"></script>`
- ❌ 禁止使用任何外部 CDN 链接

### 4.3 字体方案

生产环境禁止使用外部字体 CDN（如 Google Fonts），使用系统字体栈：

```css
:root {
  --font-sans: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, 'Noto Sans SC', sans-serif;
  --font-mono: 'SF Mono', Monaco, 'Cascadia Code', 'Roboto Mono', Consolas, monospace;
}
```