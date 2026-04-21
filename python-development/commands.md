# Python 开发常用命令

## 环境管理

### 创建虚拟环境
```powershell
# 使用 venv
python -m venv venv

# 使用 conda
conda create -n project-name python=3.13.13
```

### 激活虚拟环境
```powershell
# Windows
venv\Scripts\activate

# Linux/Mac
source venv/bin/activate

# conda
conda activate project-name
```

### 退出虚拟环境
```powershell
# venv
deactivate

# conda
conda deactivate
```

## 依赖管理

### 安装依赖
```powershell
# 安装单个包
pip install package-name

# 安装指定版本
pip install package-name==1.0.0

# 安装从文件
pip install -r requirements.txt

# 安装开发依赖
pip install -r requirements-dev.txt
```

### 升级依赖
```powershell
# 升级单个包
pip install --upgrade package-name

# 升级所有包
pip freeze --local | grep -v '^-e' | cut -d = -f 1 | xargs -n1 pip install -U
```

### 移除依赖
```powershell
# 移除单个包
pip uninstall package-name

# 移除未使用的包
pip autoremove
```

### 查看依赖
```powershell
# 查看已安装包
pip list

# 查看包详情
pip show package-name

# 检查依赖冲突
pip check
```

## 代码质量

### 代码格式化
```powershell
# 使用 ruff 格式化
ruff format .

# 使用 black 格式化
black .

# 排序导入
isort .
```

### 代码检查
```powershell
# 使用 ruff 检查
ruff check .

# 使用 flake8 检查
flake8 .

# 类型检查
mypy .

# 安全检查
bandit -r .
safety check
```

## 测试

### 运行测试
```powershell
# 运行所有测试
pytest

# 运行指定测试
pytest tests/test_module.py

# 运行指定函数
pytest tests/test_module.py::test_function

# 显示详细输出
pytest -v

# 显示覆盖率
pytest --cov=src
```

### 测试覆盖率
```powershell
# 安装覆盖率工具
pip install pytest-cov

# 运行测试并生成覆盖率报告
pytest --cov=src --cov-report=html
```

## 构建与部署

### 构建包
```powershell
# 构建源码分发包
python setup.py sdist

# 构建 wheel 包
python setup.py bdist_wheel

# 使用 build 模块
pip install build
python -m build
```

### 发布包
```powershell
# 上传到 PyPI
pip install twine
twine upload dist/*

# 上传到测试 PyPI
twine upload --repository testpypi dist/*
```

## 项目管理

### 初始化项目
```powershell
# 创建新项目目录
mkdir project-name
cd project-name

# 初始化 Git
git init

# 创建虚拟环境
python -m venv venv

# 激活虚拟环境
venv\Scripts\activate

# 安装基础依赖
pip install pytest ruff mypy

# 创建项目结构
mkdir -p src/package_name tests docs examples scripts

# 创建 __init__.py
echo "" > src/package_name/__init__.py

# 创建 pyproject.toml
New-Item -ItemType File -Path pyproject.toml

# 创建 README.md
New-Item -ItemType File -Path README.md
```

### 版本控制
```powershell
# 查看状态
git status

# 添加文件
git add .

# 提交变更
git commit -m "feat: initial commit"

# 推送到远程
git push origin main
```

## 常用工具

### 交互式解释器
```powershell
# 启动 Python 解释器
python

# 启动 IPython
pip install ipython
ipython

# 启动 Jupyter Notebook
pip install jupyter
jupyter notebook
```

### 代码分析
```powershell
# 性能分析
python -m cProfile script.py

# 内存分析
pip install memory_profiler
python -m memory_profiler script.py

# 行级性能分析
pip install line_profiler
kernprof -l -v script.py
```

### 包管理
```powershell
# 搜索包
pip search package-name

# 查看包信息
pip show package-name

# 列出过时的包
pip list --outdated
```

### 环境变量
```powershell
# Windows 设置环境变量
$env:API_KEY = "your-api-key"

# Linux/Mac 设置环境变量
export API_KEY="your-api-key"

# 使用 python-dotenv
pip install python-dotenv
# 在代码中
from dotenv import load_dotenv
load_dotenv()
```

## 调试

### 内置调试器
```powershell
# 启动调试器
python -m pdb script.py

# 在代码中设置断点
import pdb; pdb.set_trace()

# Python 3.7+ 使用 breakpoint()
breakpoint()
```

### 日志
```powershell
# 配置日志
import logging
logging.basicConfig(level=logging.DEBUG, format='%(asctime)s - %(name)s - %(levelname)s - %(message)s')

# 记录日志
logging.debug('Debug message')
logging.info('Info message')
logging.warning('Warning message')
logging.error('Error message')
logging.critical('Critical message')
```

## 其他命令

### 查看 Python 版本
```powershell
python --version
python -V
```

### 查看 pip 版本
```powershell
pip --version
pip -V
```

### 升级 pip
```powershell
pip install --upgrade pip
```

### 清理缓存
```powershell
# 清理 pip 缓存
pip cache purge

# 清理构建文件
rm -rf build dist *.egg-info
```

### 运行脚本
```powershell
# 直接运行
python script.py

# 作为模块运行
python -m package.module

# 执行字符串
python -c "print('Hello, World!')"
```

### 查看模块路径
```powershell
python -c "import sys; print(sys.path)"

# 查看包位置
python -c "import package; print(package.__file__)"
```