# 基于 DOR.COM 的校园网认证

**适用:**
Dr.COM Eportal管理系统

本项目基于 Go + Python 重写实现
>
> 1. 使用 `Python` 作为服务端 `Go`  作为客户端
> 2. `recory` 用作恢复 | `check` 用于检测是否连接
> 3. `main.go` 为入口

## 一、环境适配

提前配置 `Python3.8+` 环境

    ./win-install.bat

## 二、启动服务

```shell
#=> 使用管理员权限运行 客户端程序

# 启动服务端
python main.py

# 编译客户端
go build -o client.exe main.go
go build -o recory.exe recory.go

# 启动客户端
./client.exe

# 恢复 DHCP 模式
./recory.exe
```

## 三、接口格式

1. 文件路径：
    + `simulate\server\code\scan_start.py`

2. 数据格式：
    + 路径：`http://localhost:9000/api/v1/ip`
    + 类型：`json`
    + 示例：

    + ```json
        {
            "142": [
                { "scan_ip": "10.27.142.31", "ttl": "0.178400" },
                { "scan_ip": "10.27.142.218", "ttl": "0.032700" }
            ],
        }
        ```

3. 注意
    > 出于安全考虑 移除了 mac 信息
