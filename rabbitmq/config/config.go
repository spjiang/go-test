package config

import (
	"github.com/spjiang/go-test/rabbitmq/entity"
)

func GetRabbitmqCfg() *entity.RabbitmqCfg {
	cfg := &entity.RabbitmqCfg{
		UserName: "admin",
		Password: "123456",
		Host:     "127.0.0.1",
		Port:     15672,
		Vhost:    "/",
	}
	return cfg
}
