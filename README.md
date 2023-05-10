# 基于 DOR.COM 的校园网认证

**适用:**
Dr.COM Eportal管理系统

## 一、环境适配

提前配置 `Python3.8+` 环境

    ./win-install.bat

## 二、启动服务

    ./run.bat

## 三、接口格式

1. 文件路径：
    + `./school_network/util/server.py`

2. 数据格式：
    + 路径：`http://localhost/api/rand_ip`
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

1. 注意
    > 出于安全考虑 移除了 mac 信息
