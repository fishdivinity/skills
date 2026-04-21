#!/usr/bin/env python3
"""模型导出脚本"""

import argparse
import configparser
import os
import logging
import tensorflow as tf

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
        'model_path': config['Model']['model_path'],
        'export_dir': config['Export']['export_dir'],
        'export_format': config['Export']['export_format'],
        'quantize': config['Export'].getboolean('quantize', False)
    }

def export_model(config):
    """导出模型
    
    Args:
        config: 配置字典
    """
    # 创建导出目录
    os.makedirs(config['export_dir'], exist_ok=True)
    
    # 加载模型
    logging.info(f'加载模型: {config["model_path"]}')
    model = tf.keras.models.load_model(config['model_path'])
    
    # 导出为不同格式
    if config['export_format'] == 'tflite':
        # 转换为TensorFlow Lite模型
        converter = tf.lite.TFLiteConverter.from_keras_model(model)
        
        # 量化
        if config['quantize']:
            converter.optimizations = [tf.lite.Optimize.DEFAULT]
        
        # 转换模型
        logging.info('转换为TensorFlow Lite模型...')
        tflite_model = converter.convert()
        
        # 保存模型
        export_path = os.path.join(config['export_dir'], 'model.tflite')
        with open(export_path, 'wb') as f:
            f.write(tflite_model)
        logging.info(f'TensorFlow Lite模型保存到: {export_path}')
        
    elif config['export_format'] == 'pb':
        # 导出为SavedModel格式
        export_path = os.path.join(config['export_dir'], 'saved_model')
        logging.info(f'导出为SavedModel格式到: {export_path}')
        model.save(export_path, save_format='tf')
        logging.info('SavedModel导出完成!')
        
    elif config['export_format'] == 'onnx':
        # 导出为ONNX格式
        try:
            import tf2onnx
            
            # 转换为ONNX格式
            export_path = os.path.join(config['export_dir'], 'model.onnx')
            logging.info(f'导出为ONNX格式到: {export_path}')
            
            # 创建输入签名
            input_signature = [tf.TensorSpec([None, 224, 224, 3], tf.float32, name='input')]
            
            # 转换模型
            tf2onnx.convert.from_keras(model, input_signature=input_signature, output_path=export_path)
            logging.info('ONNX模型导出完成!')
            
        except ImportError:
            logging.error('tf2onnx库未安装，请运行: pip install tf2onnx')
            return
    
    else:
        logging.error(f'不支持的导出格式: {config["export_format"]}')
        return

def main():
    """主函数"""
    # 解析命令行参数
    parser = argparse.ArgumentParser(description='模型导出脚本')
    parser.add_argument('--config', type=str, default='configs/export_config.ini', help='配置文件路径')
    args = parser.parse_args()
    
    # 加载配置
    config = load_config(args.config)
    
    # 导出模型
    export_model(config)

if __name__ == "__main__":
    main()