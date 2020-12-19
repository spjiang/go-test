package subscribe

import  "github.com/spjiang/go-test/rabbitmq/subscribe/motor_vehicle"

type Subscribe interface {
	Handle()
}

func SubscribeHandle() {
	MotorVehicle()
}

func MotorVehicle() {
	m := &subscribe_motor_vehicle.MotorVehicle{
		ExchangeName: "motor.vehicle",
		ExchangeType: "topic",
		QueueName:    "motor.vehicle",
		BindingKey: []string{
			"motor.vehicle.new",
		},
	}
	m.Handle()
}
