# 前端项目结构模板

## 基本目录结构

```
project/
├── index.html          # 首页
├── css/                # CSS 目录
│   ├── style.css       # 主样式文件
│   ├── components.css  # 组件样式
│   └── variables.css   # 变量定义
├── js/                 # JavaScript 目录
│   ├── main.js         # 主脚本文件
│   └── components.js   # 组件脚本
└── assets/             # 资源目录
    ├── images/         # 图片
    └── icons/          # 图标
```

## 目录说明

- **index.html**：项目的首页，包含页面的基本结构
- **css/**：存放 CSS 样式文件
  - **style.css**：主样式文件，包含全局样式
  - **components.css**：组件样式文件，包含所有组件的样式
  - **variables.css**：变量定义文件，包含颜色、字体等通用样式变量
- **js/**：存放 JavaScript 脚本文件
  - **main.js**：主脚本文件，包含页面的主要逻辑
  - **components.js**：组件脚本文件，包含所有组件的逻辑
- **assets/**：存放静态资源文件
  - **images/**：存放图片文件
  - **icons/**：存放图标文件

## 组件目录结构

```
project/
└── components/         # 组件目录
    ├── button/         # 按钮组件
    │   ├── button.html # 按钮组件 HTML
    │   ├── button.css  # 按钮组件 CSS
    │   └── button.js   # 按钮组件 JavaScript
    ├── form/           # 表单组件
    │   ├── form.html   # 表单组件 HTML
    │   ├── form.css    # 表单组件 CSS
    │   └── form.js     # 表单组件 JavaScript
    └── navigation/     # 导航组件
        ├── navigation.html # 导航组件 HTML
        ├── navigation.css  # 导航组件 CSS
        └── navigation.js   # 导航组件 JavaScript
```

## 组件规范

1. **组件命名**：组件名称使用小写字母，单词间用连字符 `-` 分隔
2. **组件结构**：每个组件包含 HTML、CSS 和 JavaScript 文件
3. **组件样式**：组件样式使用 BEM 命名规范
4. **组件复用**：组件应具有可复用性，参数可配置
5. **组件文档**：每个组件应有清晰的文档说明

## 示例组件

### 按钮组件

**button.html**
```html
<!-- 按钮组件 -->
<button class="btn btn--primary">主要按钮</button>
<button class="btn btn--secondary">次要按钮</button>
<button class="btn btn--danger">危险按钮</button>
```

**button.css**
```css
/* 按钮样式 */
.btn {
  display: inline-block;
  padding: 8px 16px;
  border: none;
  border-radius: 4px;
  font-size: 14px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.btn--primary {
  background-color: #007bff;
  color: white;
}

.btn--secondary {
  background-color: #6c757d;
  color: white;
}

.btn--danger {
  background-color: #dc3545;
  color: white;
}

.btn:hover {
  opacity: 0.8;
}
```

**button.js**
```javascript
// 按钮组件逻辑
class Button {
  constructor(element) {
    this.element = element;
    this.init();
  }

  init() {
    this.element.addEventListener('click', this.handleClick.bind(this));
  }

  handleClick(event) {
    // 处理按钮点击事件
    console.log('Button clicked:', this.element.textContent);
  }
}

// 初始化按钮组件
document.querySelectorAll('.btn').forEach(button => {
  new Button(button);
});
```

### 表单组件

**form.html**
```html
<!-- 表单组件 -->
<form class="form">
  <div class="form__group">
    <label class="form__label" for="name">姓名</label>
    <input class="form__input" type="text" id="name" name="name" required>
  </div>
  <div class="form__group">
    <label class="form__label" for="email">邮箱</label>
    <input class="form__input" type="email" id="email" name="email" required>
  </div>
  <div class="form__group">
    <button class="btn btn--primary" type="submit">提交</button>
  </div>
</form>
```

**form.css**
```css
/* 表单样式 */
.form {
  max-width: 600px;
  margin: 0 auto;
}

.form__group {
  margin-bottom: 16px;
}

.form__label {
  display: block;
  margin-bottom: 8px;
  font-weight: 600;
}

.form__input {
  width: 100%;
  padding: 8px 12px;
  border: 1px solid #ced4da;
  border-radius: 4px;
  font-size: 14px;
}

.form__input:focus {
  outline: none;
  border-color: #007bff;
  box-shadow: 0 0 0 0.2rem rgba(0, 123, 255, 0.25);
}
```

**form.js**
```javascript
// 表单组件逻辑
class Form {
  constructor(element) {
    this.element = element;
    this.init();
  }

  init() {
    this.element.addEventListener('submit', this.handleSubmit.bind(this));
  }

  handleSubmit(event) {
    event.preventDefault();
    // 处理表单提交
    const formData = new FormData(this.element);
    const data = Object.fromEntries(formData);
    console.log('Form submitted:', data);
  }
}

// 初始化表单组件
document.querySelectorAll('.form').forEach(form => {
  new Form(form);
});
```

## 使用指南

1. **初始化项目**：根据上述目录结构创建项目文件夹
2. **配置样式**：在 `variables.css` 中定义颜色、字体等通用样式
3. **开发组件**：按照组件规范开发可复用组件
4. **编写代码**：遵循 HTML、CSS 规范编写代码
5. **测试与优化**：测试页面在不同设备和浏览器中的表现，进行性能优化

## 最佳实践

- 使用版本控制系统（如 Git）管理代码
- 定期备份项目文件
- 保持代码简洁明了，避免冗余代码
- 注释清晰，说明代码的功能和目的
- 遵循 Web 标准，确保页面在不同浏览器中正常显示
- 考虑用户体验，确保页面加载速度快，交互流畅