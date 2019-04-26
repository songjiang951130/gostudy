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
go get -u -v github.com/nsf/gocode
//下载失败
go get -u -v github.com/tools/imports
go get -u -v github.com/sqs/goreturns
go get -u golang.org/x/lint/golint
//手动安装
go install golang.org/x/lint/golint

//调试工具安装 executables built by Go 1.11 or later need Delve built by Go 1.11 or later
go get -u github.com/go-delve/delve/cmd/dlv
//vs code 代码分析工具
gopkgs
  go-outline
  guru
  gorename
  godef
```
