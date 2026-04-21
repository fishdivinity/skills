#!/usr/bin/env python3
"""模型训练脚本"""

import argparse
import configparser
import os
import logging
import tensorflow as tf
from tensorflow.keras.models import Sequential
from tensorflow.keras.layers import Dense, Conv2D, MaxPooling2D, Flatten
from tensorflow.keras.preprocessing.image import ImageDataGenerator

# 配置日志
logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')
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
        'data_dir': config['Data']['data_dir'],
        'model_dir': config['Model']['model_dir'],
        'epochs': int(config['Training']['epochs']),
        'batch_size': int(config['Training']['batch_size']),
        'learning_rate': float(config['Training']['learning_rate']),
        'image_size': tuple(map(int, config['Data']['image_size'].split(','))),
        'num_classes': int(config['Model']['num_classes'])
    }

def create_model(image_size, num_classes):
    """创建模型
    
    Args:
        image_size: 图像大小 (width, height)
        num_classes: 类别数量
        
    Returns:
        模型
    """
    model = Sequential([
        Conv2D(32, (3, 3), activation='relu', input_shape=(*image_size, 3)),
        MaxPooling2D((2, 2)),
        Conv2D(64, (3, 3), activation='relu'),
        MaxPooling2D((2, 2)),
        Conv2D(128, (3, 3), activation='relu'),
        MaxPooling2D((2, 2)),
        Flatten(),
        Dense(128, activation='relu'),
        Dense(num_classes, activation='softmax')
    ])
    
    return model

def train_model(config):
    """训练模型
    
    Args:
        config: 配置字典
    """
    # 创建模型目录
    os.makedirs(config['model_dir'], exist_ok=True)
    
    # 创建数据生成器
    train_datagen = ImageDataGenerator(
        rescale=1./255,
        shear_range=0.2,
        zoom_range=0.2,
        horizontal_flip=True,
        validation_split=0.2
    )
    
    # 加载训练数据
    train_generator = train_datagen.flow_from_directory(
        config['data_dir'],
        target_size=config['image_size'],
        batch_size=config['batch_size'],
        class_mode='categorical',
        subset='training'
    )
    
    # 加载验证数据
    validation_generator = train_datagen.flow_from_directory(
        config['data_dir'],
        target_size=config['image_size'],
        batch_size=config['batch_size'],
        class_mode='categorical',
        subset='validation'
    )
    
    # 创建模型
    model = create_model(config['image_size'], config['num_classes'])
    
    # 编译模型
    model.compile(
        optimizer=tf.keras.optimizers.Adam(learning_rate=config['learning_rate']),
        loss='categorical_crossentropy',
        metrics=['accuracy']
    )
    
    # 打印模型摘要
    model.summary()
    
    # 训练模型
    logging.info('开始训练模型...')
    history = model.fit(
        train_generator,
        steps_per_epoch=train_generator.samples // config['batch_size'],
        epochs=config['epochs'],
        validation_data=validation_generator,
        validation_steps=validation_generator.samples // config['batch_size']
    )
    
    # 保存模型
    model_path = os.path.join(config['model_dir'], 'model.h5')
    model.save(model_path)
    logging.info(f'模型保存到: {model_path}')
    
    return model, history

def main():
    """主函数"""
    # 解析命令行参数
    parser = argparse.ArgumentParser(description='模型训练脚本')
    parser.add_argument('--config', type=str, default='configs/train_config.ini', help='配置文件路径')
    args = parser.parse_args()
    
    # 加载配置
    config = load_config(args.config)
    
    # 训练模型
    model, history = train_model(config)
    
    # 打印训练结果
    logging.info('训练完成!')
    logging.info(f'训练准确率: {history.history["accuracy"][-1]:.4f}')
    logging.info(f'验证准确率: {history.history["val_accuracy"][-1]:.4f}')

if __name__ == "__main__":
    main()