# Python 机器学习编码规范

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

def train_model(data: List[Dict[str, float]], epochs: int = 10) -> Dict[str, float]:
    pass
```

### 类型别名
```python
from typing import TypeAlias

ImageData: TypeAlias = List[List[float]]
ModelConfig: TypeAlias = Dict[str, Union[str, int, float]]
```

## 文档字符串

### Google 风格
```python
def preprocess_image(image_path: str, target_size: Tuple[int, int]) -> ImageData:
    """预处理图像数据。

    Args:
        image_path: 图像文件路径
        target_size: 目标图像大小 (width, height)

    Returns:
        预处理后的图像数据

    Raises:
        FileNotFoundError: 当图像文件不存在时
        ValueError: 当图像格式不支持时
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

## 机器学习特定规范

### 数据处理
- 数据处理函数应该有明确的输入和输出类型
- 数据预处理和特征工程应该模块化
- 数据验证应该严格

### 模型训练
- 训练函数应该接受配置参数
- 训练过程应该有日志记录
- 训练结果应该可重现

### 模型评估
- 评估函数应该计算多种指标
- 评估结果应该可解释
- 评估过程应该可自动化

### 模型部署
- 部署代码应该与训练代码分离
- 部署模型应该经过优化
- 部署过程应该可自动化

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

### 代码优化
- 使用向量化操作
- 避免不必要的计算
- 使用合适的数据结构

### 内存管理
- 处理大数据时使用生成器
- 及时释放不再使用的资源
- 使用上下文管理器管理资源

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