// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: hello.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Hello service

type HelloService interface {
	Info(ctx context.Context, in *InfoRequest, opts ...client.CallOption) (*InfoResponse, error)
}

type helloService struct {
	c    client.Client
	name string
}

func NewHelloService(name string, c client.Client) HelloService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "hello"
	}
	return &helloService{
		c:    c,
		name: name,
	}
}

func (c *helloService) Info(ctx context.Context, in *InfoRequest, opts ...client.CallOption) (*InfoResponse, error) {
	req := c.c.NewRequest(c.name, "Hello.Info", in)
	out := new(InfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hello service

type HelloHandler interface {
	Info(context.Context, *InfoRequest, *InfoResponse) error
}

func RegisterHelloHandler(s server.Server, hdlr HelloHandler, opts ...server.HandlerOption) error {
	type hello interface {
		Info(ctx context.Context, in *InfoRequest, out *InfoResponse) error
	}
	type Hello struct {
		hello
	}
	h := &helloHandler{hdlr}
	return s.Handle(s.NewHandler(&Hello{h}, opts...))
}

type helloHandler struct {
	HelloHandler
}

func (h *helloHandler) Info(ctx context.Context, in *InfoRequest, out *InfoResponse) error {
	return h.HelloHandler.Info(ctx, in, out)
}
