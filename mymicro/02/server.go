package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-plugins/registry/etcdv3"
	"log"
	"mymicro/02/proto"
)

// 申明一个结构体用于实现接口
type Example struct {
}

func (h *Example) Call(ctx context.Context, in *proto.CallRequest, out *proto.CallResponse) (err error) {
	log.Print("收到Example.call请求:"+in.Name)
	if len(in.Name) == 0 {
		return errors.BadRequest("go.micro.api.example", "no name")
	}
	out.Message = "Example.Call收到你的请求" + in.Name
	return
}

// 申明一个结构体用于实现接口
type Foo struct {
}

func (h *Foo) Bar(ctx context.Context, in *proto.EmptyRequest, out *proto.EmptyResponse) (err error) {
	log.Print("收到Foo.Bar请求")
	return
}

func main() {
	registre := etcdv3.NewRegistry()
	// 1.得到服务端实例
	service := micro.NewService(
		micro.Registry(registre),
		// 设置微服务名，用来访问
		// micro call hello Hello.Info{"name":"zhansan"}
		micro.Name("go.micro.api.example"),
	)
	// 2.初始化
	service.Init()
	// 3.服务注册
	err := proto.RegisterExampleHandler(service.Server(), new(Example))
	if err != nil {
		fmt.Println(err)
	}

	err = proto.RegisterFooHandler(service.Server(), new(Foo))
	if err != nil {
		fmt.Println(err)
	}
	// 4.启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
