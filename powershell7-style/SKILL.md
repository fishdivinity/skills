---
name: "powershell7-style"
description: "Ensures command line operations use PowerShell 7 syntax and style. Invoke when executing any command line operations to maintain consistent PowerShell 7 compatibility."
---

# PowerShell 7 命令行风格

本技能确保所有命令行操作使用 PowerShell 7 语法和风格规范。

## 触发条件

**必须调用此技能的场景：**
- 执行任何命令行操作时
- 在终端运行脚本或命令时
- 创建批处理或自动化脚本时
- 使用 PowerShell 特定功能时

## AI 调用指南

### 调用时机
当需要执行命令行操作时，AI 应首先调用此技能获取正确的命令格式和风格指南。

### 调用方式
```
使用 Skill 工具调用：powershell7-style
```

### 调用后行为
1. 阅读并理解本技能提供的 PowerShell 7 规范
2. 按照规范格式化所有命令
3. 使用 PowerShell 7 语法执行命令

## PowerShell 7 语法规范

### 基本语法约定
- 使用 `;` 分隔同一行的多个命令
- 使用 `&&` 和 `||` 进行管道操作（仅 PowerShell 7+ 支持）
- 使用 `Get-ChildItem` 代替 `ls` 以确保跨平台兼容性
- 使用 `Select-Object` 代替 `select`
- 使用 `Where-Object` 代替 `where`

### 命令对照表

| 操作 | PowerShell 7 风格 | 避免使用 |
|------|-------------------|----------|
| 列出文件 | `Get-ChildItem -Path .` | `ls` |
| 过滤结果 | `Where-Object { $_.Name -like "*.ps1" }` | `where` |
| 选择属性 | `Select-Object Name, LastWriteTime` | `select` |
| 多命令 | `cmd1; cmd2` | `cmd1 && cmd2` (除非需要条件执行) |
| 条件执行 | `cmd1 && cmd2` 或 `cmd1 || cmd3` | bash 风格 |

### 常用命令模板

#### 文件操作
```powershell
# 列出当前目录所有文件
Get-ChildItem -Path . -File

# 递归查找特定文件
Get-ChildItem -Path . -Recurse -Filter "*.ps1" | Select-Object Name, FullName, LastWriteTime

# 创建目录
New-Item -Path ".\newFolder" -ItemType Directory -Force

# 复制文件
Copy-Item -Path ".\source.txt" -Destination ".\dest.txt"

# 删除文件
Remove-Item -Path ".\old.txt" -Force
```

#### 进程管理
```powershell
# 查看进程
Get-Process | Where-Object { $_.CPU -gt 100 } | Select-Object Name, CPU, Id

# 停止进程
Stop-Process -Name "processName" -Force
```

#### 网络操作
```powershell
# 测试网络连接
Test-NetConnection -ComputerName "example.com" -Port 443

# 下载文件
Invoke-WebRequest -Uri "https://example.com/file.zip" -OutFile ".\file.zip"
```

#### 字符串处理
```powershell
# 字符串分割
$string -split ","

# 字符串替换
$string -replace "old", "new"

# 字符串匹配
$string -match "pattern"
```

### 管道操作
```powershell
# PowerShell 7 风格
Get-ChildItem -Path . -Recurse -Filter "*.ps1" | 
    Where-Object { $_.Length -gt 1KB } | 
    Select-Object Name, Length, LastWriteTime | 
    Sort-Object Length -Descending
```

### 条件执行（PowerShell 7+ 特性）
```powershell
# 成功后执行
Write-Output "step1" && Write-Output "step2"

# 失败后执行
Write-Error "error" || Write-Output "fallback"
```

## RunCommand 工具使用规范

当使用 `RunCommand` 工具时：
- 确保 `command` 参数使用 PowerShell 7 语法
- 默认 shell 为 PowerShell 7
- 多行命令使用反引号 `` ` `` 进行续行
- 复杂命令先在 PowerShell 7 环境中测试

### 示例
```powershell
# 正确的多行命令格式
Get-ChildItem -Path . -Recurse `
    | Where-Object { $_.Extension -eq ".ps1" } `
    | Select-Object Name, FullName
```

## 最佳实践

1. **使用完整 cmdlet 名称**：提高可读性和可维护性
2. **利用管道能力**：PowerShell 7 的管道功能强大，应充分利用
3. **改进的错误处理**：使用 `try-catch-finally` 和 `$ErrorActionPreference`
4. **显式调用**：需要时可使用 `pwsh` 命令明确调用 PowerShell 7
5. **参数别名**：了解常用参数别名，但优先使用完整参数名

## 禁止使用的模式

- 不要使用 `cmd.exe` 特定语法
- 不要使用 Linux/Bash 特定命令（除非有 PowerShell 等效命令）
- 不要混用不同 shell 的语法

## 本技能帮助维护一致性和兼容性，专注于 PowerShell 7 的特性和语法。
