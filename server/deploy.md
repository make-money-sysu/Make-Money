首先先确定golang已经成功安装，不同的系统对应不同的安装方法，具体安装教程见[go安装教程](<https://golang.org/doc/install>)

然后执行以下代码把项目下载到本地：

```
go get github.com/make-money-sysu/server
```

然后执行以下代码安装bee工具：

```
go get github.com/beego/bee
```

请务必将$GOPATH/bin加入到环境变量中，保证bee工具能正确运行

然后运行如下命令安装所有的依赖(部分依赖可能需要翻墙)：

```
go get ./...
```

然后进入$GOPATH/src/github.com/make-money/server文件夹（也就是源码的存储位置），运行如下命令：

```
bee run -gendoc=true -downdoc=true
```

这样后端程序就已经运行起来了