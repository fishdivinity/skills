# Python 性能规范

## 代码优化

### 算法选择
- 选择合适的算法和数据结构
- 了解时间复杂度
- 避免不必要的计算

### 示例
```python
# 低效的代码
def find_duplicates(lst):
    duplicates = []
    for i in range(len(lst)):
        for j in range(i + 1, len(lst)):
            if lst[i] == lst[j]:
                duplicates.append(lst[i])
    return duplicates

# 高效的代码
def find_duplicates(lst):
    seen = set()
    duplicates = set()
    for item in lst:
        if item in seen:
            duplicates.add(item)
        else:
            seen.add(item)
    return list(duplicates)
```

## 内存优化

### 生成器
- 使用生成器处理大文件
- 避免一次性加载全部数据
- 利用 `yield` 关键字

### 示例
```python
# 内存密集型
def read_file(filename):
    with open(filename, 'r') as f:
        return f.readlines()  # 加载全部行到内存

# 内存高效
def read_file(filename):
    with open(filename, 'r') as f:
        for line in f:  # 逐行读取
            yield line
```

### 数据结构
- 使用合适的数据结构
- 字典用于快速查找
- 集合用于去重和成员检查
- 列表用于有序数据

### 示例
```python
# 低效的成员检查
valid_users = ['user1', 'user2', 'user3']
if username in valid_users:  # O(n) 时间复杂度
    # 处理...

# 高效的成员检查
valid_users = {'user1', 'user2', 'user3'}
if username in valid_users:  # O(1) 时间复杂度
    # 处理...
```

## 计算优化

### 避免重复计算
- 缓存计算结果
- 使用 `functools.lru_cache`
- 预计算常量

### 示例
```python
from functools import lru_cache

@lru_cache(maxsize=None)
def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)
```

### 向量化操作
- 使用 NumPy 进行向量化计算
- 避免显式循环
- 利用 C 扩展的性能

### 示例
```python
import numpy as np

# 低效的循环
result = []
for i in range(1000):
    result.append(i * 2)

# 高效的向量化
result = np.arange(1000) * 2
```

## I/O 优化

### 文件操作
- 使用上下文管理器
- 批量读取和写入
- 避免频繁的 I/O 操作

### 示例
```python
# 低效的文件写入
with open('output.txt', 'w') as f:
    for i in range(1000):
        f.write(f"Line {i}\n")  # 1000次写入操作

# 高效的文件写入
lines = [f"Line {i}\n" for i in range(1000)]
with open('output.txt', 'w') as f:
    f.writelines(lines)  # 1次写入操作
```

### 网络操作
- 使用连接池
- 批量请求
- 异步 I/O

### 示例
```python
import aiohttp
import asyncio

async def fetch(session, url):
    async with session.get(url) as response:
        return await response.text()

async def main():
    urls = ["http://example.com" for _ in range(100)]
    async with aiohttp.ClientSession() as session:
        tasks = [fetch(session, url) for url in urls]
        results = await asyncio.gather(*tasks)
    return results
```

## 并行处理

### 多线程
- 使用 `threading` 模块
- 适合 I/O 密集型任务
- 注意 GIL 限制

### 示例
```python
import threading

def process_data(data):
    # 处理数据...
    pass

data_chunks = [data[i:i+1000] for i in range(0, len(data), 1000)]
threads = []

for chunk in data_chunks:
    thread = threading.Thread(target=process_data, args=(chunk,))
    threads.append(thread)
    thread.start()

for thread in threads:
    thread.join()
```

### 多进程
- 使用 `multiprocessing` 模块
- 适合 CPU 密集型任务
- 绕过 GIL 限制

### 示例
```python
from multiprocessing import Pool

def process_item(item):
    # 处理单个项目...
    return result

items = range(1000)

with Pool(processes=4) as pool:
    results = pool.map(process_item, items)
```

## 异步编程

### asyncio
- 使用 `async` 和 `await`
- 适合 I/O 密集型任务
- 提高并发性能

### 示例
```python
import asyncio

async def async_task(name, delay):
    print(f"Task {name} started")
    await asyncio.sleep(delay)
    print(f"Task {name} completed")

async def main():
    tasks = [
        async_task("A", 1),
        async_task("B", 2),
        async_task("C", 3)
    ]
    await asyncio.gather(*tasks)

asyncio.run(main())
```

## 性能分析

### 内置工具
- `cProfile`: 函数级性能分析
- `line_profiler`: 行级性能分析
- `memory_profiler`: 内存使用分析

### 示例
```python
import cProfile

def slow_function():
    # 慢代码...
    pass

cProfile.run('slow_function()')
```

### 第三方工具
- **py-spy**: 采样分析器
- **scalene**: CPU, 内存和GPU分析器
- **pyflame**: 火焰图生成

### 示例
```powershell
# 使用 py-spy
pip install py-spy
py-spy record -o profile.svg -- python script.py
```

## 代码优化技巧

### 字符串操作
- 使用 `join()` 连接字符串
- 避免重复字符串拼接
- 使用 f-strings （Python 3.6+）

### 循环优化
- 避免在循环中进行昂贵操作
- 使用列表推导式
- 减少循环内部的函数调用

### 示例
```python
# 低效的字符串拼接
s = ""
for i in range(1000):
    s += str(i)  # 创建新字符串

# 高效的字符串拼接
s = "".join(str(i) for i in range(1000))
```

## 缓存策略

### 内存缓存
- 使用 `functools.lru_cache`
- 实现自定义缓存
- 注意缓存大小和失效策略

### 示例
```python
from functools import lru_cache

@lru_cache(maxsize=1000)
def compute_expensive_value(n):
    # 计算密集型操作
    result = 0
    for i in range(n):
        result += i
    return result
```

### 外部缓存
- 使用 Redis
- 实现缓存失效策略
- 缓存热点数据

### 示例
```python
import redis

r = redis.Redis(host='localhost', port=6379, db=0)

def get_data(key):
    # 尝试从缓存获取
    cached = r.get(key)
    if cached:
        return cached.decode('utf-8')
    
    # 缓存未命中，计算数据
    data = compute_data(key)
    
    # 存入缓存
    r.set(key, data, ex=3600)  # 1小时过期
    
    return data
```

## 最佳实践总结

1. **先分析，后优化**：使用性能分析工具识别瓶颈
2. **选择合适的算法和数据结构**：了解时间复杂度
3. **避免不必要的计算**：缓存结果，预计算常量
4. **优化 I/O 操作**：批量处理，使用异步
5. **利用并行和异步**：根据任务类型选择合适的并发模型
6. **内存管理**：使用生成器，避免一次性加载大量数据
7. **代码简洁**：简洁的代码往往更高效
8. **定期测试**：持续监控性能变化