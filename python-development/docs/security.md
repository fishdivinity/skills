# Python 安全规范

## 输入验证

### 所有输入都不可信
- 验证所有用户输入
- 使用类型检查和范围检查
- 避免使用 `eval()` 和 `exec()`

### 示例
```python
# 不安全的代码
def process_input(user_input):
    result = eval(user_input)  # 危险！
    return result

# 安全的代码
def process_input(user_input):
    if not isinstance(user_input, str):
        raise ValueError("Input must be a string")
    # 安全处理...
    return result
```

## 密码处理

### 密码存储
- 使用 `bcrypt` 或 `argon2` 哈希密码
- 不要明文存储密码
- 不要使用 MD5 或 SHA1 等弱哈希算法

### 示例
```python
import bcrypt

def hash_password(password: str) -> bytes:
    salt = bcrypt.gensalt()
    return bcrypt.hashpw(password.encode('utf-8'), salt)

def check_password(password: str, hashed: bytes) -> bool:
    return bcrypt.checkpw(password.encode('utf-8'), hashed)
```

## 敏感信息管理

### 环境变量
- 使用 `.env` 文件和 `python-dotenv`
- 不要将敏感信息硬编码到代码中
- 将 `.env` 文件添加到 `.gitignore`

### 示例
```python
from dotenv import load_dotenv
import os

load_dotenv()

API_KEY = os.getenv('API_KEY')
DATABASE_URL = os.getenv('DATABASE_URL')
```

## SQL 注入防护

### 参数化查询
- 使用 ORM（如 SQLAlchemy）
- 使用参数化查询
- 避免拼接 SQL 语句

### 示例
```python
# 不安全的代码
def get_user(username):
    query = f"SELECT * FROM users WHERE username = '{username}'"  # 危险！
    # 执行查询...

# 安全的代码
def get_user(username):
    query = "SELECT * FROM users WHERE username = %s"
    # 使用参数化查询执行...
```

## 跨站脚本（XSS）防护

### 输出转义
- 转义所有用户生成的内容
- 使用安全的模板引擎
- 实施内容安全策略（CSP）

### 示例
```python
from markupsafe import escape

def render_user_input(user_input):
    return escape(user_input)
```

## 跨站请求伪造（CSRF）防护

### 令牌验证
- 使用 CSRF 令牌
- 验证请求来源
- 实施 SameSite cookie 属性

### 示例
```python
from flask_wtf.csrf import CSRFProtect

csrf = CSRFProtect(app)
```

## 依赖安全

### 依赖检查
- 定期更新依赖
- 使用 `pip-audit` 或 `safety` 检查安全漏洞
- 锁定依赖版本

### 示例
```powershell
# 检查依赖安全
pip install pip-audit
pip-audit

# 或使用 safety
pip install safety
safety check
```

## 错误处理

### 安全的错误处理
- 不要向用户暴露详细错误信息
- 记录详细错误信息
- 使用通用错误消息

### 示例
```python
try:
    # 操作...
except Exception as e:
    logger.error(f"Error: {e}")
    return {"error": "An internal error occurred"}
```

## 文件操作

### 安全的文件操作
- 验证文件路径
- 限制文件访问权限
- 避免路径遍历攻击

### 示例
```python
import os

def safe_file_path(base_dir, filename):
    # 解析路径
    full_path = os.path.join(base_dir, filename)
    # 确保路径在基础目录内
    if not full_path.startswith(os.path.abspath(base_dir)):
        raise ValueError("Invalid file path")
    return full_path
```

## 网络安全

### HTTPS
- 使用 HTTPS
- 验证 SSL 证书
- 避免明文传输敏感信息

### 示例
```python
import requests

response = requests.get('https://api.example.com', verify=True)
```

## 安全配置

### 最小权限原则
- 应用以最小必要权限运行
- 限制文件和目录权限
- 配置适当的防火墙规则

### 示例
```python
# 设置文件权限
import os

os.chmod('config.py', 0o600)  # 只有所有者可读写
```

## 安全审计

### 代码审查
- 定期进行安全代码审查
- 使用静态代码分析工具
- 实施安全测试

### 工具
- **Bandit**: 查找常见安全问题
- **Safety**: 检查依赖安全漏洞
- **Pylint**: 代码质量和安全检查

### 示例
```powershell
# 使用 Bandit
pip install bandit
bandit -r src/
```