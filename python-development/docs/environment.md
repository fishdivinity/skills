# Python 环境管理

## 虚拟环境

### 为什么使用虚拟环境
- 隔离项目依赖
- 避免版本冲突
- 便于项目迁移和部署
- 保持系统环境清洁

### 创建虚拟环境

#### 使用 venv
```powershell
# 创建虚拟环境
python -m venv venv

# 激活虚拟环境（Windows）
venv\Scripts\activate

# 激活虚拟环境（Linux/Mac）
source venv/bin/activate
```

#### 使用 conda
```powershell
# 创建虚拟环境
conda create -n project-name python=3.13.13

# 激活虚拟环境
conda activate project-name
```

## 依赖管理

### requirements.txt
```powershell
# 生成依赖文件
pip freeze > requirements.txt

# 安装依赖
pip install -r requirements.txt
```

### 开发依赖
创建 `requirements-dev.txt` 用于开发依赖：
```
# 核心依赖
-r requirements.txt

# 开发依赖
pytest
ruff
mypy
black
isort
pre-commit
```

## 项目结构

### 标准项目结构
```
project-name/
├── src/
│   └── package_name/
│       ├── __init__.py
│       ├── core/
│       ├── utils/
│       └── exceptions.py
├── tests/
│   ├── unit/
│   ├── integration/
│   └── conftest.py
├── docs/
├── examples/
├── scripts/
├── requirements.txt
├── requirements-dev.txt
├── pyproject.toml
├── setup.py
└── README.md
```

### 模块划分
- **src/**: 源代码目录
- **tests/**: 测试代码目录
- **docs/**: 文档目录
- **examples/**: 示例代码目录
- **scripts/**: 脚本工具目录

## 工具配置

### pyproject.toml
```toml
[project]
name = "package-name"
version = "0.1.0"
description = "Description of the package"
authors = [
    { name = "Author Name", email = "author@example.com" },
]
dependencies = [
    "requests>=2.31.0",
    "pydantic>=2.0.0",
]

[build-system]
requires = ["setuptools", "wheel"]
build-backend = "setuptools.build_meta"

[tool.pytest.ini_options]
testpaths = ["tests"]
pythonpath = ["src"]

[tool.pre-commit]
repos = [
    {
        repo = "https://github.com/astral-sh/ruff-pre-commit",
        rev = "v0.4.1",
        hooks = [
            { id = "ruff", args = ["--fix"] },
            { id = "ruff-format" },
        ],
    },
    {
        repo = "https://github.com/pre-commit/mirrors-mypy",
        rev = "v1.9.0",
        hooks = [
            { id = "mypy" },
        ],
    },
]
```

## 常见问题

### 虚拟环境激活问题
- **Windows**: 确保使用正确的激活脚本路径
- **权限问题**: 以管理员身份运行终端
- **路径问题**: 检查Python和pip的路径设置

### 依赖冲突
- 使用 `pip check` 检查依赖冲突
- 锁定依赖版本
- 考虑使用 `pip-tools` 管理依赖

### 性能优化
- 使用 `pip install --no-cache-dir` 加速安装
- 考虑使用 `uv` 或 `pipenv` 作为替代工具