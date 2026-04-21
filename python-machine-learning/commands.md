# Python 机器学习常用命令

## 环境管理

### 创建虚拟环境
```powershell
# 使用 venv
python -m venv venv

# 使用 conda
conda create -n ml-project python=3.13.13
```

### 激活虚拟环境
```powershell
# Windows
venv\Scripts\activate

# Linux/Mac
source venv/bin/activate

# conda
conda activate ml-project
```

### 退出虚拟环境
```powershell
# venv
deactivate

# conda
conda deactivate
```

## 依赖管理

### 安装依赖
```powershell
# 安装机器学习库
pip install tensorflow tensorflow-gpu
pip install torch torchvision
pip install keras

# 安装数据处理库
pip install numpy pandas scikit-learn

# 安装图像处理库
pip install opencv-python pillow

# 安装音频处理库
pip install librosa pydub sounddevice

# 安装OCR库
pip install pytesseract

# 安装语音识别库
pip install SpeechRecognition

# 安装模型部署库
pip install tensorflow-lite
pip install onnx onnxruntime

# 安装其他工具
pip install matplotlib seaborn jupyter
```

### 导出依赖
```powershell
# 生成依赖文件
pip freeze > requirements.txt

# 安装依赖
pip install -r requirements.txt
```

## 数据处理

### 数据下载
```powershell
# 使用wget下载数据
wget https://example.com/dataset.zip

# 使用curl下载数据
curl -O https://example.com/dataset.zip

# 解压数据
unzip dataset.zip
```

### 数据预处理
```powershell
# 运行数据预处理脚本
python scripts/preprocess_data.py

# 使用pandas处理数据
python -c "import pandas as pd; df = pd.read_csv('data.csv'); print(df.head())"
```

## 模型训练

### 训练模型
```powershell
# 运行训练脚本
python scripts/train_model.py --config configs/train_config.ini

# 使用TensorBoard监控训练
tensorboard --logdir=logs
```

### 评估模型
```powershell
# 运行评估脚本
python scripts/evaluate_model.py --config configs/eval_config.ini

# 使用scikit-learn评估模型
python -c "from sklearn.metrics import accuracy_score; y_true = [0, 1, 1, 0]; y_pred = [0, 1, 0, 0]; print(accuracy_score(y_true, y_pred))"
```

## 模型部署

### 导出模型
```powershell
# 运行导出脚本
python scripts/export_model.py --config configs/export_config.ini

# 转换为TensorFlow Lite模型
python -c "import tensorflow as tf; model = tf.keras.models.load_model('models/model.h5'); converter = tf.lite.TFLiteConverter.from_keras_model(model); tflite_model = converter.convert(); open('models/model.tflite', 'wb').write(tflite_model)"
```

### 模型优化
```powershell
# 量化模型
python -c "import tensorflow as tf; model = tf.keras.models.load_model('models/model.h5'); converter = tf.lite.TFLiteConverter.from_keras_model(model); converter.optimizations = [tf.lite.Optimize.DEFAULT]; tflite_model = converter.convert(); open('models/model_quantized.tflite', 'wb').write(tflite_model)"

# 模型剪枝
python -c "import tensorflow as tf; from tensorflow_model_optimization.sparsity import keras as sparsity; model = tf.keras.models.load_model('models/model.h5'); pruning_schedule = sparsity.PolynomialDecay(initial_sparsity=0.0, final_sparsity=0.5, begin_step=0, end_step=1000); pruned_model = sparsity.prune_low_magnitude(model, pruning_schedule); pruned_model.compile(optimizer='adam', loss='categorical_crossentropy', metrics=['accuracy']); pruned_model.save('models/model_pruned.h5')"
```

## 性能分析

### 分析模型性能
```powershell
# 使用cProfile分析训练脚本
python -m cProfile -o train_profile.out scripts/train_model.py

# 查看分析结果
python -c "import pstats; p = pstats.Stats('train_profile.out'); p.sort_stats('cumulative').print_stats(10)"

# 分析模型推理时间
python -c "import time; import tensorflow as tf; model = tf.keras.models.load_model('models/model.h5'); import numpy as np; input_data = np.random.random((1, 224, 224, 3)); start = time.time(); model.predict(input_data); end = time.time(); print(f'Inference time: {end - start:.4f} seconds')"
```

## 安全检查

### 依赖安全检查
```powershell
# 使用safety检查依赖
pip install safety
safety check

# 使用bandit检查代码
pip install bandit
bandit -r scripts/
```

## 版本控制

### Git命令
```powershell
# 初始化Git仓库
git init

# 添加文件
git add .

# 提交更改
git commit -m "feat: initial commit"

# 推送到远程
git push origin main

# 分支管理
git checkout -b feature/ocr-implementation
git merge main
```

## 容器化

### Docker命令
```powershell
# 构建镜像
docker build -t ml-model .

# 运行容器
docker run -p 8000:8000 ml-model

# 推送镜像
docker push username/ml-model:latest
```

## 云服务

### AWS命令
```powershell
# 安装AWS CLI
pip install awscli

# 配置AWS CLI
aws configure

# 上传模型到S3
aws s3 cp models/model.h5 s3://my-bucket/models/
```

### GCP命令
```powershell
# 安装gcloud CLI
# 参考: https://cloud.google.com/sdk/docs/install

# 配置gcloud
 gcloud init

# 上传模型到GCS
gcloud storage cp models/model.h5 gs://my-bucket/models/
```

## 其他命令

### 查看Python版本
```powershell
python --version
python -V
```

### 查看CUDA版本
```powershell
nvcc --version
nvidia-smi
```

### 清理缓存
```powershell
# 清理pip缓存
pip cache purge

# 清理TensorFlow缓存
rm -rf ~/.cache/tensorflow

# 清理PyTorch缓存
rm -rf ~/.cache/torch
```

### 运行Jupyter Notebook
```powershell
# 启动Jupyter Notebook
jupyter notebook

# 启动Jupyter Lab
jupyter lab
```

### 查看GPU使用情况
```powershell
# Windows
nvidia-smi

# Linux
top -p $(pgrep -d',' python)

# 监控GPU使用
watch -n 1 nvidia-smi
```