# 项目结构模板

```
project/
├── cmd/
│   └── server/
│       └── main.go         # 应用入口
├── internal/
│   ├── api/               # API层
│   │   ├── handlers/      # 请求处理器
│   │   ├── middlewares/   # 中间件
│   │   └── routes.go      # 路由配置
│   ├── config/            # 配置管理
│   │   └── config.go      # 配置结构和加载
│   ├── models/            # 数据模型
│   ├── services/          # 业务逻辑
│   └── utils/             # 工具函数
├── pkg/                   # 可共享的包
│   ├── auth/              # 认证相关
│   ├── database/          # 数据库连接
│   └── logger/            # 日志工具
├── scripts/               # 脚本文件
│   ├── build.sh           # 构建脚本
│   └── deploy.sh          # 部署脚本
├── .env.example           # 环境变量示例
├── .gitignore             # Git忽略文件
├── Dockerfile             # Docker构建文件
├── go.mod                 # Go模块文件
└── README.md              # 项目说明
```

## 说明

- **cmd/**: 应用入口点，包含main函数
- **internal/**: 内部代码，不对外暴露
- **pkg/**: 可共享的代码包
- **scripts/**: 构建和部署脚本

## 使用方法

1. 复制此结构到新的项目目录
2. 修改go.mod文件，更新模块名
3. 根据需要调整目录结构
