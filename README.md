# apiserver

一个基于Go语言和Gin框架构建的轻量级RESTful API服务器。

## 功能特性

- 基于Gin Web框架，高性能HTTP服务
- 配置管理（Viper支持YAML配置文件和环境变量）
- 日志轮转功能（使用lumberjack）
- 健康检查接口
- 系统监控接口（磁盘、CPU、内存）
- 安全HTTP头部设置
- CORS支持
- 配置热重载

## 项目结构

```
apiserver/
├── conf/          # 配置文件目录
├── config/        # 配置管理代码
├── handler/       # 请求处理函数
├── router/        # 路由配置
├── main.go        # 程序入口
└── go.mod         # Go模块定义
```

## 快速开始

### 环境要求

- Go 1.16+

### 安装步骤

```bash
# 克隆项目
git clone <repository-url>

# 进入项目目录
cd apiserver

# 下载依赖
go mod tidy

# 编译
go build -o apiserver

# 运行
./apiserver
```

### 配置

项目默认会在`conf/`目录下查找`config.yaml`配置文件。你也可以通过命令行参数指定配置文件：

```bash
./apiserver -c /path/to/config.yaml
```

配置项可以通过环境变量覆盖，环境变量前缀为`APISERVER_`，例如：
- `APISERVER_RUNMODE=release`
- `APISERVER_ADDR=:8080`

### API接口

#### 健康检查
- `GET /sd/health` - 服务健康检查
- `GET /sd/disk` - 磁盘空间检查
- `GET /sd/cpu` - CPU使用率检查
- `GET /sd/ram` - 内存使用率检查

## 开发指南

### 添加新的API接口

1. 在`handler/`目录下创建处理函数
2. 在`router/router.go`中添加路由配置

### 日志配置

日志支持文件输出和控制台输出，并且可以配置日志级别和轮转策略。

## 许可证

MIT