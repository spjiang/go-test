package subscribe_motor_vehicle

import (
	"fmt"
	"github.com/spjiang/go-test/rabbitmq"
	"github.com/gogf/gf/frame/g"
	"log"
	"time"
)

type MotorVehicle struct {
	ExchangeName string
	ExchangeType string
	QueueName    string
	BindingKey   []string
}

func (m *MotorVehicle) Handle() {
	m.process()
}

func (m *MotorVehicle) process() {
	defer func() {
		if err := recover(); err != nil {
			time.Sleep(3 * time.Second)
			g.Log().Debug("RabbitMQ MotorVehicle process panic")
			fmt.Println("RabbitMQ MotorVehicle process panic", err)
			m.process()
		}
	}()

	conn, err := rabbitmq.Connection()
	rabbitmq.Panic(err, "MotorVehicle Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	rabbitmq.Panic(err, "MotorVehicle Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	rabbitmq.Panic(err, "MotorVehicle Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		m.QueueName, // name
		false,       // durable
		false,       // delete when unused
		true,        // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	rabbitmq.Panic(err, "MotorVehicle Failed to declare a queue")

	for _, routingKey := range m.BindingKey {
		err = ch.QueueBind(
			q.Name,         // queue name
			routingKey,     // routing key
			m.ExchangeName, // exchange
			false,
			nil)
		rabbitmq.Panic(err, "MotorVehicle Failed to bind a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	rabbitmq.Panic(err, "MotorVehicle Failed to register a consumer")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("MotorVehicle Received a message: %s", d.Body)
			m.handleMsg(d.Body)
		}
	}()
	<-forever
}

func (m *MotorVehicle) handleMsg(msg []byte) {
	fmt.Println(msg)
}
