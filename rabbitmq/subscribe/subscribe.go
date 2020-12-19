package subscribe

import "github.com/spjiang/go-test/rabbitmq/entity"

type Subscribe interface {
	Handle()
}

func SubscribeHandle() {
	MotorVehicle()
}

func MotorVehicle() {
	m := &entity.MotorVehicle{
		ExchangeName: "motor.vehicle",
		ExchangeType: "topic",
		QueueName:    "motor.vehicle",
		BindingKey: []string{
			"motor.vehicle.new",
		},
	}
	m.Handle()
}
