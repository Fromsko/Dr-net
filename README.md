# 内网IP伪装

`net-tools` 是一款由 Go语言 编写的 内网伪装IP的工具 (GUI) 版

基于Arp协议的扫描工具, 能够扫描指定网段存活主机信息 [IP | MAC | TTL]

---

## 版本信息

> 工具: net-tools.exe
>
> 版本：V4.1
>
> 构建语言：go1.20.2
>
> 系统环境：windows/amd64

项目版本更迭：

+ `V1.0` Python 实现 基础功能
+ `v2.0` Python 实现 安全的多进程
+ `V3.0` Python + Go 1.18 实现 CLi版
+ `V4.0` Python + Go 20.2 实现 GUI + Gin 版

本人已离校, 该项目将不限期停更, 项目仅供参考

## 运行截屏

![img.png](res/img.png)

# 如何使用

> 本项目只支持在 `Win` 平台下使用

项目出发点：

+ 课表数据更新不及时
+ 电费数据难以获取
+ 内网增加了基于MAC的心跳检测机制

注意：

+ 不保证适用于所有系统
+ 项目只作为学习技术交流
+ 不提供任何违反规章制度的技术支援

搭建建议：

+ 使用本工具获取IP数据后 实现一个定时任务 进行自动替换IP地址
+ 通过修改 路由器的 IP地址 | 网关 可以实现多端使用
+ 对接青龙面板 和 port-forward 可以实现 安全的校内数据任务的爬取

适用范围：校园内网基于(Dr.COM)的认证系统

适用人群：喜欢Golang的朋友

## 编译

```shell
cd Dr-net
# 获取依赖
go mod tidy
# 安装 fyne 的命令行工具
go install fyne.io/fyne/v2/cmd/fyne@latest
# 编译并打包
fyne package -os windows -icon myapp.png -tags noconsole
```

# 鸣谢

+ [arp-scan](https://github.com/QbsuranAlang/arp-scan-windows-) - C实现的`arp`扫描
+ [scanPort](https://github.com/xs25cn/scanPort) - Go 实现的端口扫描工具
+ [fyne](https://github.com/fyne-io/fyne) - Go 流行的`GUI`框架
+ [qinglong](https://github.com/whyour/qinglong) - 定时任务管理平台
+ [port-forward](https://github.com/tavenli/port-forward) - Go实现的端口转发工具
