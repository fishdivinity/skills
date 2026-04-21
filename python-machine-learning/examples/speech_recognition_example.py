"""语音识别示例代码，用于识别语音记账内容。

此模块提供了以下功能：
1. 使用 Google Speech-to-Text API 识别音频文件中的语音
2. 从识别的文本中提取金额、分类、日期和描述信息
3. 打印提取的信息

依赖项：
- speech_recognition: 用于语音识别
- re: 用于正则表达式匹配

使用示例：
    python speech_recognition_example.py

注意：
- 需要将音频文件命名为 "voice_note.wav" 并放在同一目录下
- 需要联网才能使用 Google Speech-to-Text API
"""


import re
import speech_recognition as sr

def recognize_speech(audio_path: str) -> str:
    """识别语音

    Args:
        audio_path: 音频文件路径

    Returns:
        识别的文本
    """
    # 初始化识别器
    r = sr.Recognizer()
    
    # 加载音频文件
    with sr.AudioFile(audio_path) as source:
        audio = r.record(source)
    
    try:
        # 使用Google Speech-to-Text
        text = r.recognize_google(audio, language='zh-CN')
        return text
    except sr.UnknownValueError:
        return "无法识别音频"
    except sr.RequestError as e:
        return f"API请求失败: {e}"


def process_text(text: str) -> dict[str, str | None]:
    """处理识别的文本

    Args:
        text: 识别的文本

    Returns:
        提取的信息字典
    """
    information: dict[str, str | None] = {
        "amount": None,
        "category": None,
        "date": None,
        "description": None
    }
    
    # 提取金额
    amount_pattern = r"(\d+(\.\d{1,2})?)"
    match = re.search(amount_pattern, text)
    if match:
        information["amount"] = match.group(1)
    
    # 提取分类
    categories = ["餐饮", "交通", "购物", "娱乐", "医疗", "教育", "其他"]
    for category in categories:
        if category in text:
            information["category"] = category
            break
    
    # 提取日期
    date_pattern = r"(今天|昨天|前天|\d{4}-\d{2}-\d{2}|\d{2}/\d{2})"
    match = re.search(date_pattern, text)
    if match:
        information["date"] = match.group(1)
    
    # 提取描述
    information["description"] = text
    
    return information


def main() -> None:
    """主函数"""

    # 音频路径
    audio_path = "voice_note.wav"
    
    # 识别语音
    text = recognize_speech(audio_path)
    print(f"识别的文本: {text}")
    
    # 处理文本
    information = process_text(text)
    
    # 打印结果
    print("提取的信息:")
    print(f"金额: {information['amount']}")
    print(f"分类: {information['category']}")
    print(f"日期: {information['date']}")
    print(f"描述: {information['description']}")

if __name__ == "__main__":
    main()