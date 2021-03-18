package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"tgrpc/proto"
)

func main() {
	// 地址
	addr := "127.0.0.1:8080"
	// 1、监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		// e:=fmt.Errorf("监听异常：%s\n",err)
		fmt.Printf("监听异常：%s\n", err)
	}
	fmt.Printf("监听端口：%s\n", addr)
	// 2、实例化gRPC
	s := grpc.NewServer()
	// 3、在gRPC上注册微服务
	proto.RegisterUserInfoServiceServer(s, &proto.U)
	// 4、启动服务端
	s.Serve(listener)

}
