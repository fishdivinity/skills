---
name: python-machine-learning
description: Python machine learning specialized skill with TensorFlow, PyTorch, and deep learning best practices. Invoke for machine learning development, OCR image recognition, speech recognition, model training, evaluation, and deployment.
---

# Python 机器学习技能

## 概述
Python 3.13+ 机器学习开发指南，用于构建生产级（Production-Ready）机器学习应用。专注于OCR、语音识别、模型训练和跨平台部署，为BeeCount ML项目提供全面支持。

## 技术栈
- **Python版本**: 3.13.13
- **机器学习框架**: TensorFlow 2.0+, PyTorch 1.7+
- **模型部署**: TensorFlow Lite, PyTorch Mobile, ONNX Runtime
- **辅助工具**: OpenVINO, Tesseract OCR, SpeechRecognition
- **环境管理**: 虚拟环境（venv）
- **代码质量**: PEP8 + 严格类型检查

## 工作流程
1. **需求确认**: 调用`askuserquestion`，确认项目需求和技术栈
2. **环境准备**: 创建虚拟环境，安装必要的依赖
3. **数据处理**: 数据收集、预处理和标注
4. **模型开发**: 模型设计、训练和评估
5. **模型部署**: 模型导出、优化和部署
6. **质量检查**: 运行lint检查和类型检查

## 核心规范索引
- **编码规范**: `docs/coding-standards.md`
- **模型训练**: `docs/model-training.md`
- **模型部署**: `docs/model-deployment.md`
- **OCR实现**: `docs/ocr-implementation.md`
- **语音识别**: `docs/speech-recognition.md`
- **性能优化**: `docs/performance.md`
- **安全规范**: `docs/security.md`

## 禁止模式
- 将所有逻辑写在一个Python文件中
- 不使用虚拟环境管理依赖
- 忽略PEP8规范
- 硬编码敏感信息
- 不使用类型提示
- 捕获异常后不处理
- 不进行模型评估和测试
- 不优化模型性能

## 代码审查清单
- 遵循PEP8规范
- 使用类型提示
- 函数职责单一
- 代码模块化
- 错误处理完善
- 无硬编码敏感信息
- 虚拟环境配置正确
- 测试覆盖率高
- 模型性能评估完整
- 部署配置合理

## 资源索引
- **示例代码**: `examples/`
- **模板文件**: `templates/`
- **脚本工具**: `scripts/`
- **命令清单**: `commands.md`

## 文件结构
```
python-machine-learning/
├── SKILL.md              # 技能描述文件（本文件）
├── docs/                 # 详细文档
├── examples/             # 示例代码
├── templates/            # 模板文件
├── scripts/              # 脚本工具
└── commands.md           # 常用命令清单
```