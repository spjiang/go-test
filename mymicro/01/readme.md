> 1. go get github.com/micro/micro
> 2. go get github.com/micro/go-micro
> 3. go get github.com/micro/protoc-gen-micro


#PROTOBUF
###命令

```protobuf
protoc --micro_out=. --go_out=. hello.proto
```

> 1. proto_path：当前proto文件目录
> 2. go_out：生成go文件目录

#执行方式
###命令
```shell
 micro call hello Hello.Info{"name":"zhansan"}
```

#软件降级
###确定版本
```go version```
###安装版本
```brew install go@1.14```
###unlink
`brew unlink go`
###强制link
`brew link --force go@1.14`

#加入环境变量
``echo 'export PATH="/usr/local/opt/go@1.14/bin:$PATH"' >> ~/.zshrc``


#安装micro
###1. 安装micro和所需依赖
`go get -u -v github.com/micro/micro`
###2.安装micro
`go install github.com/micro/micro`
