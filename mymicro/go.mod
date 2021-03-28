module mymicro

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/golang/protobuf v1.4.3
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/etcdv3 v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro v1.18.0 // indirect
	github.com/micro/micro/v3 v3.1.1 // indirect
	google.golang.org/protobuf v1.25.0
)
