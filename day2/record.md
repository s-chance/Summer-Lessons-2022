## 记录nodejs使用过程中的问题
前言:该文档主要目的是尽可能减少占用C盘空间。
前提环境(特别注意):
- node版本:node-v16.16.0-x64
- node安装在D盘并且npm全局安装路径指定D盘
- 通过npm install的方式来安装Yarn
### 问题1:Could not find file D:\Program
node安装的路径D:\Program Files\nodejs\存在空格,导致命令找不到,默认C盘安装时,也可能存在这个问题
解决:
- 重新安装node时,指定没有中文、没有空格、没有特殊字符的路径
- 其他方法参考:https://stackoverflow.com/questions/16395612/
### 问题2:node npm yarn安装后,无法正常使用,显示报错,不是内部或外部命令,也不是可运行的程序或批处理文件
- Yarn默认安装在C盘,如果指定nodejs安装路径为D盘,则需要修改Yarn的路径
解决:
```bash
# 查找yarn管理的文件路径
yarn global dir
# 找到该文件后移动到D盘或者删除

# 对yarn管理的文件安装、缓存路径进行修改
yarn config set global-folder "D:\Yarn\Data\global"
yarn config set cache-folder "D:\Yarn\Cache\v6"
# 之后通过yarn管理的文件都存放到上面指定的文件夹中
```

