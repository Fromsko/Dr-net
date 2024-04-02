# Dr-net

Dr.com 校园网绕过（模拟登录）

## 开发

- 拉取依赖

  ```shell
  go mod tidy
  ```

- 获取 `wails`

  ```shell
  go install github.com/wailsapp/wails/v2/cmd/wails@latest
  ```

- 编译构建

  ```shell
  wails build -upx -nsis
  ```

- 可选依赖
  - [`UPX`](https://upx.github.io/) 用于压缩您的应用程序。
  - [`NSIS`](https://wails.io/docs/guides/windows-installer/) 用于生成 Windows 安装程序。

## 界面

- `主界面`
  ![](./backend/res/index.png)
- `管理页面`
  ![](./backend/res/config.png)

## 使用

1. 开始扫描
2. 点击使用(同时只能有一个被使用)
3. 设置 - 检测(查询网络状态)
4. 清空(移除所有扫描结果|恢复网络)

## 贡献

- 开发模板

  ```shell
  wails init -n myproject -t vue-ts
  ```

- 开发模式

  ```shell
  wails dev
  ```

- 构建模式(携带`debug`)

  > -windowsconsole 保留控制台
  >
  > --debug 启用 debug 模式

  ```shell
  wails build --debug
  ```

- 构建模式(打包+压缩)

  > -s 跳过前端构建
  >
  > -upx 压缩
  >
  > -nsis `Win` 打包程序

  ```shell
  wails build -upx -nsis
  ```

## 鸣谢

- 雏形 <a href="https://github.com/Fromsko/Dr-net">https://github.com/Fromsko/Dr-net</a>
