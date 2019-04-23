
``` bash
#GOROOT 为 /libexec 这文件夹
export GOROOT=/usr/local/Cellar/go/1.12.4/libexec
export GOPATH=$HOME/gopath
#默认 mod 配置，
export GO111MODULE=auto
#设置代理下载代码
export GOPROXY=https://goproxy.io
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```