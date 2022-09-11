## 简单使用

### 环境准备

- 安装nodejs运行前端
- 安装mysql存储必要的数据
- 安装go语言做为后端
- 安装Git进行命令输入(非必要, 但推荐安装)

### 开始使用

- 在lessons-table目录下使用Git Bash, 并输入以下命令($符后面的内容)
    ```Bash
    # 需要先通过npm安装好tyarn
    $ npm install -g tyarn --registry https://registry.npm.taobao.org
    # 由于vite已经创建过, 这里就不需要再create vite
    $ tyarn
    $ tyarn dev
    ```
    根据Bash命令行提示的url访问前端
    注意: 在运行前端时不要直接退出Bash, 应先按`Ctrl`+`c`结束进程再退出, 否则nodejs可能会一直占用后台资源

- 在demo目录下打开Bash或VS Code的终端, 确认处于demo目录下, 执行以下命令
    ```Bash
    # 这里的demo就是指demo目录这个模块
    $ go mod init demo
    # 如果目录下已经存在go.mod, 则不需要init
    # 可以通过下面的命令直接自动引用和删除依赖
    $ go mod tidy
    # 运行后端
    $ go run main.go
    ```
    后端的端口不能直接去访问, 而是根据前端已经编写好的代码, 通过用户与前端的交互实现访问

## 特别鸣谢[精弘网络](https://github.com/zjutjh/Summer-Lessons-2022)

