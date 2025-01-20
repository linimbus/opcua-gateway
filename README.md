# opcua-gateway

## 项目背景

## 功能介绍

### MYSQL数据存储

- 支持自动建表建库
- 支持历史数据清理

### OPCUA客户端

- 支持多个OPCUA数据订阅配置
- 支持数据存储到MYSQL数据库

### OPCUA服务端

- 支持OPCUA服务网关代理查询
- 支持OPCUA服务自定义端口

## 编译过程

### 准备环境

- Windows10/11 64位
- [Golang SDK](https://studygolang.com/dl/golang/go1.23.3.windows-amd64.msi)
- [GCC 编译器](https://jmeubank.github.io/tdm-gcc/download/)

### 环境变量

- GOPROXY=<https://goproxy.cn,direct>
- GOPATH=D:\workspace\golang
- PATH=C:\TDM-GCC-64\bin:D:\workspace\golang\bin

### 准备工具

```
go install github.com/akavel/rsrc@latest
go install github.com/jteeuwen/go-bindata/...@latest
```

### 编译

```
git clone xxx
cd opua-gateway
.\build.bat
```

### 测试
