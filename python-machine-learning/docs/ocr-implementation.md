# OCR 实现规范

## 概述

OCR（Optical Character Recognition，光学字符识别）是BeeCount ML项目的核心功能之一，用于自动识别支付截图中的金额、商家、分类等信息。本规范提供了OCR功能的实现指南，包括技术选择、模型训练、部署优化等方面。

## 技术选择

### OCR引擎
- **Tesseract OCR**: 开源OCR引擎，支持多种语言
- **EasyOCR**: 基于深度学习的OCR库，支持80+语言
- **PaddleOCR**: 百度开源的OCR库，性能优异
- **Google Cloud Vision OCR**: 云端OCR服务，准确性高

### 深度学习框架
- **TensorFlow**: 广泛使用的深度学习框架，支持模型部署
- **PyTorch**: 灵活的深度学习框架，适合研究和开发
- **Keras**: 高级神经网络API，易于使用

### 预处理工具
- **OpenCV**: 图像处理库，用于图像预处理
- **Pillow**: 图像处理库，用于图像操作
- **scikit-image**: 图像处理库，用于高级图像处理

## 实现流程

### 1. 图像预处理

#### 步骤
1. **图像加载**: 加载原始图像
2. **灰度转换**: 将彩色图像转换为灰度图像
3. **二值化**: 将灰度图像转换为二值图像
4. **噪声去除**: 去除图像噪声
5. **图像增强**: 增强图像对比度和清晰度
6. **图像分割**: 分割图像中的文本区域

#### 代码示例
```python
import cv2
import numpy as np

def preprocess_image(image_path):
    """预处理图像
    
    Args:
        image_path: 图像文件路径
        
    Returns:
        预处理后的图像
    """
    # 加载图像
    image = cv2.imread(image_path)
    
    # 转换为灰度图像
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    
    # 二值化
    _, binary = cv2.threshold(gray, 0, 255, cv2.THRESH_BINARY + cv2.THRESH_OTSU)
    
    # 噪声去除
    kernel = np.ones((1, 1), np.uint8)
    binary = cv2.morphologyEx(binary, cv2.MORPH_OPEN, kernel)
    binary = cv2.morphologyEx(binary, cv2.MORPH_CLOSE, kernel)
    
    # 图像增强
    clahe = cv2.createCLAHE(clipLimit=2.0, tileGridSize=(8, 8))
    enhanced = clahe.apply(gray)
    
    return enhanced
```

### 2. 文本检测

#### 方法
- **基于规则的方法**: 使用边缘检测、轮廓分析等方法
- **基于深度学习的方法**: 使用目标检测模型（如YOLO、Faster R-CNN）
- **专门的文本检测模型**: 使用CTPN、EAST、PSENet等模型

#### 代码示例
```python
import cv2
import numpy as np

def detect_text_regions(image):
    """检测文本区域
    
    Args:
        image: 预处理后的图像
        
    Returns:
        文本区域边界框列表
    """
    # 使用EAST文本检测器
    net = cv2.dnn.readNet('frozen_east_text_detection.pb')
    
    # 获取图像尺寸
    (H, W) = image.shape[:2]
    
    # 设置输入层
    blob = cv2.dnn.blobFromImage(image, 1.0, (W, H), (123.68, 116.78, 103.94), swapRB=True, crop=False)
    net.setInput(blob)
    
    # 前向传播
    (scores, geometry) = net.forward(["feature_fusion/Conv_7/Sigmoid", "feature_fusion/concat_3"])
    
    # 解码检测结果
    boxes = decode_predictions(scores, geometry)
    
    return boxes
```

### 3. 文本识别

#### 方法
- **基于传统OCR引擎**: 使用Tesseract OCR
- **基于深度学习的方法**: 使用CRNN、RNN-T等模型
- **端到端OCR**: 使用同时进行检测和识别的模型

#### 代码示例
```python
import pytesseract
from PIL import Image

def recognize_text(image, boxes):
    """识别文本
    
    Args:
        image: 原始图像
        boxes: 文本区域边界框列表
        
    Returns:
        识别结果列表
    """
    results = []
    
    for (startX, startY, endX, endY) in boxes:
        # 裁剪文本区域
        roi = image[startY:endY, startX:endX]
        
        # 使用Tesseract OCR识别文本
        text = pytesseract.image_to_string(roi, lang='chi_sim+eng')
        
        # 过滤空文本
        if text.strip():
            results.append((text.strip(), (startX, startY, endX, endY)))
    
    return results
```

### 4. 信息提取

#### 方法
- **基于规则的方法**: 使用正则表达式提取关键信息
- **基于模板的方法**: 使用预定义模板匹配信息
- **基于机器学习的方法**: 使用分类模型识别信息类型

#### 代码示例
```python
import re

def extract_information(texts):
    """提取关键信息
    
    Args:
        texts: 识别的文本列表
        
    Returns:
        提取的信息字典
    """
    information = {
        'amount': None,
        'merchant': None,
        'date': None,
        'category': None
    }
    
    # 提取金额
    amount_pattern = r'¥?\s*([0-9]+(\.[0-9]{1,2})?)'
    for text, _ in texts:
        match = re.search(amount_pattern, text)
        if match:
            information['amount'] = match.group(1)
            break
    
    # 提取商家
    # 这里可以根据具体场景定制规则
    
    # 提取日期
    date_pattern = r'(\d{4}-\d{2}-\d{2}|\d{2}/\d{2}/\d{4})'
    for text, _ in texts:
        match = re.search(date_pattern, text)
        if match:
            information['date'] = match.group(1)
            break
    
    return information
```

## 模型训练

### 数据集准备
- **公开数据集**: 使用公开OCR数据集（如ICDAR、COCO-Text）
- **自定义数据集**: 收集和标注支付截图数据
- **数据增强**: 对训练数据进行增强，提高模型泛化能力

### 模型选择
- **检测模型**: EAST、PSENet、TextFuseNet
- **识别模型**: CRNN、RNN-T、Transformer
- **端到端模型**: DTRB、TextBoxes++

### 训练流程
1. **数据预处理**: 对数据进行预处理和标注
2. **模型配置**: 配置模型参数和训练超参数
3. **模型训练**: 训练模型并监控训练过程
4. **模型评估**: 在验证集上评估模型性能
5. **模型优化**: 优化模型性能和推理速度

## 部署优化

### 模型压缩
- **知识蒸馏**: 使用知识蒸馏压缩模型
- **量化**: 对模型进行量化，减少模型大小
- **剪枝**: 对模型进行剪枝，减少参数量

### 推理优化
- **批处理**: 使用批处理提高推理速度
- **硬件加速**: 利用GPU、NPU等硬件加速
- **边缘优化**: 针对边缘设备进行优化

### 移动设备部署
- **TensorFlow Lite**: 使用TensorFlow Lite部署模型
- **PyTorch Mobile**: 使用PyTorch Mobile部署模型
- **模型大小**: 控制模型大小，适应移动设备内存

## 性能评估

### 评估指标
- **检测准确率**: 文本区域检测的准确率
- **识别准确率**: 文本识别的准确率
- **信息提取准确率**: 关键信息提取的准确率
- **推理速度**: 模型推理的速度
- **内存使用**: 模型运行时的内存使用

### 评估方法
- **测试集评估**: 在测试集上评估模型性能
- **真实场景测试**: 在真实场景中测试模型性能
- **A/B测试**: 与其他OCR方案进行比较

## 最佳实践

### 预处理优化
- **自适应预处理**: 根据图像质量自动调整预处理步骤
- **实时预处理**: 优化预处理速度，确保实时性
- **鲁棒性**: 提高预处理的鲁棒性，适应不同质量的图像

### 模型选择
- **场景适配**: 根据具体场景选择合适的模型
- **精度与速度平衡**: 平衡模型精度和推理速度
- **迁移学习**: 利用预训练模型进行迁移学习

### 部署策略
- **混合部署**: 结合云端和边缘部署
- **增量更新**: 支持模型增量更新
- **故障恢复**: 实现故障恢复机制

### 持续改进
- **数据收集**: 持续收集真实场景数据
- **模型迭代**: 基于新数据迭代更新模型
- **性能监控**: 监控模型在生产环境的性能