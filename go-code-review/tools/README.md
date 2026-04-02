# Tools 目录

此目录用于存放构建后的代码审查工具二进制文件。

## 构建说明

在 `review/` 目录下运行以下命令构建工具：

```bash
# 构建并安装到此目录
make install
```

## 文件说明

- `review.exe` (Windows) / `review` (Linux/Mac) - 代码审查工具的可执行文件

## 注意事项

此目录下的二进制文件（.exe）不会被提交到 Git 仓库。如果需要分发工具，请从源码重新构建。
