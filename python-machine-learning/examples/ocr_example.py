"""OCR示例代码，用于识别支付截图中的金额、商家等信息。"""

import cv2
import pytesseract
import re
import numpy as np
from typing import List, Tuple, Dict, Optional


def preprocess_image(image_path: str) -> np.ndarray:
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


def detect_text_regions(image: np.ndarray) -> List[Tuple[int, int, int, int]]:
    """检测文本区域

    Args:
        image: 预处理后的图像

    Returns:
        文本区域边界框列表
    """
    # 使用Tesseract OCR的页面分割模式
    custom_config = r'--oem 3 --psm 6'

    # 获取OCR结果
    d = pytesseract.image_to_data(image, config=custom_config, output_type=pytesseract.Output.DICT)

    boxes: List[Tuple[int, int, int, int]] = []
    n_boxes = len(d['level'])

    for i in range(n_boxes):
        if d['conf'][i] > 60:  # 只考虑置信度大于60的结果
            (x, y, w, h) = (d['left'][i], d['top'][i], d['width'][i], d['height'][i])
            boxes.append((x, y, x + w, y + h))

    return boxes


def recognize_text(
    image: np.ndarray,
    boxes: List[Tuple[int, int, int, int]]
) -> List[Tuple[str, Tuple[int, int, int, int]]]:
    """识别文本

    Args:
        image: 原始图像
        boxes: 文本区域边界框列表

    Returns:
        识别结果列表
    """
    results: List[Tuple[str, Tuple[int, int, int, int]]] = []

    for (startX, startY, endX, endY) in boxes:
        # 裁剪文本区域
        roi = image[startY:endY, startX:endX]

        # 使用Tesseract OCR识别文本
        custom_config = r'--oem 3 --psm 6'
        text = pytesseract.image_to_string(roi, config=custom_config, lang='chi_sim+eng')

        # 过滤空文本
        if text.strip():
            results.append((text.strip(), (startX, startY, endX, endY)))

    return results


def extract_information(
    texts: List[Tuple[str, Tuple[int, int, int, int]]]
) -> Dict[str, Optional[str]]:
    """提取关键信息

    Args:
        texts: 识别的文本列表

    Returns:
        提取的信息字典
    """
    information: Dict[str, Optional[str]] = {
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
    merchant_keywords = ['商家', '商户', '店铺', '商店', '餐厅', '饭店']
    for text, _ in texts:
        for keyword in merchant_keywords:
            if keyword in text:
                # 提取关键词后面的文本作为商家名称
                index = text.find(keyword)
                if index != -1:
                    merchant = text[index + len(keyword):].strip()
                    if merchant:
                        information['merchant'] = merchant
                        break
        if information['merchant']:
            break

    # 提取日期
    date_pattern = r'(\d{4}-\d{2}-\d{2}|\d{2}/\d{2}/\d{4}|\d{4}/\d{2}/\d{2})'
    for text, _ in texts:
        match = re.search(date_pattern, text)
        if match:
            information['date'] = match.group(1)
            break

    # 提取分类
    categories = ['餐饮', '交通', '购物', '娱乐', '医疗', '教育', '其他']
    for text, _ in texts:
        for category in categories:
            if category in text:
                information['category'] = category
                break
        if information['category']:
            break

    return information


def main() -> None:
    """主函数"""
    # 图像路径
    image_path = 'payment_screenshot.jpg'

    # 预处理图像
    enhanced = preprocess_image(image_path)

    # 检测文本区域
    boxes = detect_text_regions(enhanced)

    # 识别文本
    original_image = cv2.imread(image_path)
    texts = recognize_text(original_image, boxes)

    # 提取信息
    information = extract_information(texts)

    # 打印结果
    print("提取的信息:")
    print(f"金额: {information['amount']}")
    print(f"商家: {information['merchant']}")
    print(f"日期: {information['date']}")
    print(f"分类: {information['category']}")


if __name__ == "__main__":
    main()
