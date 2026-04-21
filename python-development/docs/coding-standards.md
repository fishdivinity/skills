# Python 编码规范

## PEP8 规范

### 缩进
- 使用4个空格进行缩进
- 不要使用制表符（Tab）
- 续行应该对齐，或者使用悬挂缩进

### 行长度
- 每行不超过79个字符
- 长表达式可以使用反斜杠换行，或利用括号自动换行

### 命名规范
- **变量和函数**: snake_case（小写字母和下划线）
- **类**: PascalCase（首字母大写）
- **常量**: UPPER_SNAKE_CASE（全大写加下划线）
- **私有成员**: 前缀下划线（如 `_internal_method`）
- **模块**: 小写字母，可使用下划线

### 空格使用
- 运算符两侧加空格
- 逗号后加空格
- 冒号后加空格
- 函数参数列表中，等号两侧不加空格

### 导入组织
- 分组导入：标准库、第三方库、本地模块
- 使用绝对导入，避免相对导入
- 每行一个导入

## 类型提示

### 基本类型
```python
from typing import List, Dict, Optional, Union, Tuple, Callable

def calculate_average(numbers: List[float]) -> Optional[float]:
    pass
```

### 类型别名
```python
from typing import TypeAlias

UserId: TypeAlias = str
UserDict: TypeAlias = Dict[str, str]
```

## 文档字符串

### Google 风格
```python
def process_data(data: List[str], timeout: int = 30) -> bool:
    """处理数据并返回结果。

    Args:
        data: 要处理的数据列表
        timeout: 超时时间（秒）

    Returns:
        处理是否成功

    Raises:
        ValueError: 当数据格式错误时
    """
    pass
```

## 代码结构

### 模块划分
- 每个模块职责单一
- 避免过大的模块（建议不超过500行）
- 相关功能放在同一模块

### 函数设计
- 函数职责单一
- 函数长度适中（建议不超过50行）
- 使用默认参数和关键字参数提高可读性

### 类设计
- 类职责单一
- 使用属性装饰器管理属性
- 实现 `__repr__` 和 `__str__` 方法

## 错误处理

### 异常处理
- 使用特定异常类型
- 创建自定义异常类
- 异常处理要具体，避免捕获所有异常

### 日志记录
- 使用 `logging` 模块
- 日志级别合理使用
- 日志信息包含足够上下文

## 性能优化

### 数据结构
- 使用合适的数据结构
- 字典和集合用于快速查找
- 列表推导式比显式循环更高效

### 内存管理
- 生成器用于处理大文件
- 上下文管理器管理资源
- 避免不必要的对象创建

## 工具配置

### pyproject.toml
```toml
[tool.ruff]
line-length = 79
target-version = "py313"

[tool.ruff.lint]
select = [
    "E",    # pycodestyle errors
    "W",    # pycodestyle warnings
    "F",    # pyflakes
    "I",    # isort
    "B",    # flake8-bugbear
    "C4",   # flake8-comprehensions
    "UP",   # pyupgrade
    "SIM",  # flake8-simplify
]

[tool.ruff.format]
quote-style = "double"
indent-style = "space"

[tool.mypy]
python_version = "3.13"
strict = true
warn_return_any = true
warn_unused_ignores = true
disallow_untyped_defs = true
disallow_incomplete_defs = true
```