> 1. go get github.com/micro/micro
> 2. go get github.com/micro/go-micro
> 3. go get github.com/micro/protoc-gen-micro


#PROTOBUF
###命令

```protobuf
protoc --micro_out=. --go_out=. api.proto
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


#安装mac环境micro
###1. 安装micro和所需依赖
`go get -u -v github.com/micro/micro`
###2.安装micro
`go install github.com/micro/micro`


#服务启动
##API网关服务启动(默认mDNS)
`micro api --handler=rpc`
``
2021-03-23 17:06:23.417088 I | [api] Registering API RPC Handler at /
2021-03-23 17:06:23.417987 I | [api] HTTP API Listening on [::]:8080
2021-03-23 17:06:23.418617 I | [api] Transport [http] Listening on [::]:56351
2021-03-23 17:06:23.419382 I | [api] Broker [http] Connected to [::]:56352
2021-03-23 17:06:23.419704 I | [api] Registry [mdns] Registering node: go.micro.api-0792356c-ca66-4c48-ab6c-d6b4595a7469
``

##API网关服务启动(推荐使用ETCD)
``micro --registry=etcd api --handler=rpc``


### 安装ETCD服务端
>1、brew search etcd
>2、brew install etcd
### 安装ETCDweb端
>下载源代码
`git clone https://github.com/henszey/etcd-browser.git`
>
>cd etcd-browser/
> 
>vim server.js  
>
`var etcdHost = process.env.ETCD_HOST || '127.0.0.1';  # etcd 主机IP`

`var etcdPort = process.env.ETCD_PORT || 4001;          # etcd 主机端口`

`var serverPort = process.env.SERVER_PORT || 8000;      # etcd-browser 监听端口`