package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	"log"
	"mymicro/01/proto"
)

// 申明一个结构体用于实现接口
type Hello struct {
}

func (h *Hello) Info(ctx context.Context, in *proto.InfoRequest, out *proto.InfoResponse) (err error) {
	out.Msg = "您好，" + in.Name
	return
}
func main() {
	registre := etcdv3.NewRegistry()
	// 1.得到服务端实例
	service := micro.NewService(
		micro.Registry(registre),
		// 设置微服务名，用来访问
		// micro call hello Hello.Info{"name":"zhansan"}
		micro.Name("hello"),
	)
	// 2.初始化
	service.Init()
	// 3.服务注册
	err := proto.RegisterHelloHandler(service.Server(), new(Hello))
	if err != nil {
		fmt.Println(err)
	}
	// 4.启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
