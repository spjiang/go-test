package registry

import "time"

type Options struct {
	// 地址
	Addrs []string
	// 超时时间
	Timeout time.Duration
	// 心跳时间
	HeartBeat int64
	// 注册地址
	RegistryPath string
}

// 定义函数类型的变量
type Option func(opts *Options)

func WithAddrs(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}


func WithHeartBeat(heartBeat int64) Option {
	return func(opts *Options) {
		opts.HeartBeat = heartBeat
	}
}

func WithRegistryPath(registryPath string) Option {
	return func(opts *Options) {
		opts.RegistryPath = registryPath
	}
}