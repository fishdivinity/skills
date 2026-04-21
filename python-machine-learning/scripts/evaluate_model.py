#!/usr/bin/env python3
"""模型评估脚本"""

from __future__ import annotations

import argparse
import configparser
import os
import logging
import numpy as np
import tensorflow as tf
from tensorflow.keras.preprocessing.image import ImageDataGenerator
from sklearn.metrics import classification_report, confusion_matrix


# 配置日志
logging.basicConfig(
    level=logging.INFO, 
    format='%(asctime)s - %(levelname)s - %(message)s'
)


def load_config(config_path):
    """加载配置文件
    
    Args:
        config_path: 配置文件路径
        
    Returns:
        配置字典
    """
    config = configparser.ConfigParser()
    config.read(config_path)
    
    return {
        'test_data_dir': config['Data']['test_data_dir'],
        'model_path': config['Model']['model_path'],
        'batch_size': int(config['Evaluation']['batch_size']),
        'image_size': tuple(map(int, config['Data']['image_size'].split(',')))
    }


def evaluate_model(config):
    """评估模型
    
    Args:
        config: 配置字典
        
    Returns:
        评估结果
    """
    # 加载模型
    logging.info(f'加载模型: {config["model_path"]}')
    model = tf.keras.models.load_model(config['model_path'])
    
    # 创建测试数据生成器
    test_datagen = ImageDataGenerator(rescale=1./255)
    
    # 加载测试数据
    test_generator = test_datagen.flow_from_directory(
        config['test_data_dir'],
        target_size=config['image_size'],
        batch_size=config['batch_size'],
        class_mode='categorical',
        shuffle=False
    )
    
    # 评估模型
    logging.info('评估模型...')
    loss, accuracy = model.evaluate(test_generator)
    
    # 预测
    logging.info('生成预测...')
    predictions = model.predict(test_generator)
    y_pred = np.argmax(predictions, axis=1)
    y_true = test_generator.classes
    
    # 计算分类报告
    class_names = list(test_generator.class_indices.keys())
    report = classification_report(y_true, y_pred, target_names=class_names)
    
    # 计算混淆矩阵
    cm = confusion_matrix(y_true, y_pred)
    
    return {
        'loss': loss,
        'accuracy': accuracy,
        'classification_report': report,
        'confusion_matrix': cm,
        'class_names': class_names
    }


def main():
    """主函数"""
    # 解析命令行参数
    parser = argparse.ArgumentParser(description='模型评估脚本')
    parser.add_argument(
        '--config', 
        type=str, 
        default='configs/eval_config.ini', 
        help='配置文件路径'
    )
    args = parser.parse_args()
    
    # 加载配置
    config = load_config(args.config)
    
    # 评估模型
    results = evaluate_model(config)
    
    # 打印评估结果
    logging.info('评估完成!')
    logging.info(f'测试损失: {results["loss"]:.4f}')
    logging.info(f'测试准确率: {results["accuracy"]:.4f}')
    logging.info('\n分类报告:')
    logging.info(results['classification_report'])
    logging.info('\n混淆矩阵:')
    logging.info(results['confusion_matrix'])


if __name__ == "__main__":
    main()