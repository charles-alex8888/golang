# 标准go工程的目录结构

- src --- 源代码
- bin --- 编译后生成的可执行文件
- pkg --- 编译时生成的中间文件

# 配置go mod

go env -w GOBIN=/usr/local/bin/go
go env -w GO111MODULE=on # on off auto
go env -w GOPROXY=https://goproxy.cn,direct // 使用七牛云的

# 使用go mod管理一个新项目
~~~
#  
mkdir project
cd project

# 
go mod init project

#
module project
~~~



# vscode插件
1. Go
2. Code Runner
3. Run in Terminal (setting设置)
4. Chinese (Simplified) Language Pack for Visual Studio Code
5. 配置代理
~~~ bash
echo "export GO111MODULE=on" >> ~/.profile
echo "export GOPROXY=https://goproxy.cn" >> ~/.profile
$ source ~/.profile
~~~

# 参考
> https://www.liwenzhou.com/

# 快捷键
~~~ 
1、注释：
　　a) 单行注释：[ctrl+k,ctrl+c] 或 ctrl+/

　　b) 取消单行注释：[ctrl+k,ctrl+u] (按下ctrl不放，再按k + u)

　　c) 多行注释：[alt+shift+A]

　　d) 多行注释：/**
2、移动行：alt+up/down，选中按TAB右移，按SHIFT+TAB左移
3、显示/隐藏左侧目录栏 ctrl + b
4、复制当前行：shift + alt +up/down
5、删除当前行：shift + ctrl + k
6、控制台终端显示与隐藏：ctrl + ~
7、查找文件/安装vs code 插件地址：ctrl + p
8、代码格式化：shift + alt +f
9、新建一个窗口 : ctrl + shift + n
10、行增加缩进: ctrl + [
11、行减少缩进: ctrl + ]
12、裁剪尾随空格(去掉一行的末尾那些没用的空格) : ctrl + shift + x
13、字体放大/缩小: ctrl + ( + 或 - )
14、拆分编辑器 : ctrl + 1/2/3
15、切换窗口 : ctrl + shift + left/right
16、关闭编辑器窗口 : ctrl + w
17、关闭所有窗口 : ctrl + k + w
18、切换全屏 : F11
19、自动换行 : alt + z
20、显示git : ctrl + shift + g
21、全局查找文件：ctrl + shift + f
22、显示相关插件的命令(如：git log)：ctrl + shift + p
23、选中文字：shift + left / right / up / down
24、折叠代码： ctrl + k + 0-9 (0是完全折叠)
25、*展开代码： ctrl + k + j (完全展开代码)
26、删除行 ： ctrl + shift + k
27、快速切换主题：ctrl + k / ctrl + t
28、快速回到顶部 ： ctrl + home
29、快速回到底部 : ctrl + end
30、格式化选定代码 ：ctrl + k / ctrl +f
~~~
