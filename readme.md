### 学习 go 语言的项目

环境: `mac pro` , 编译器 `VS CODE` , 安装器 `homebrew`

go version : 1.12.4，
mac os 上 brew 安装
```bash
brew install go@1.12
```

构建方式 : mod

go [环境变量配置](config/bash_profile.md)

```bash
#GOROOT 为 /libexec 这文件夹
export GOROOT=/usr/local/Cellar/go/1.12.4/libexec
export GOPATH=$HOME/gopath
#默认 mod 配置
export GO111MODULE=auto
#设置代理下载代码
export GOPROXY=https://goproxy.io
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

### ide 辅助安装

```bash
go get -u -v golang.org/x/tools/cmd/gopls
go get github.com/nsf/gocode
go get github.com/tools/imports
go get github.com/sqs/goreturns
```
