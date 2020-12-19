package rabbitmq

import (
	"fmt"
	"github.com/spjiang/go-test/rabbitmq/config"
	"github.com/gogf/gf/frame/g"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

//type RabbitMQ struct {
//	Conn    *amqp.Connection
//	Channel *amqp.Channel
//}

//我们还需要一个辅助函数来检查每个amqp调用的返回值：
//func failOnError(err error, msg string) {
//	if err != nil {
//		log.Fatalf("%s: %s", msg, err)
//	}
//}

//func New() (*RabbitMQ, error) {
//	mq := &RabbitMQ{}
//	conn, err := RabbitMQConn()
//	if err != nil {
//		return nil, err
//	}
//	failOnError(err, "RabbitMQ连接失败")
//	mq.Conn = conn
//
//	ch, err := RabbitMQConnChannel(conn)
//	if err != nil {
//		return nil, err
//	}
//	failOnError(err, "RabbitMQ打开通道失败")
//	mq.Channel = ch
//	return mq, nil
//}

func Connection() (*amqp.Connection, error) {
	cfg := config.GetRabbitmqCfg()
	addr := fmt.Sprintf("amqp://%s:%s@%s:%d/%s", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.Vhost)
	conn, err := amqp.Dial(addr)
	if err != nil {
		g.Log().Info("Failed to connect to RabbitMQ", zap.Error(err))
		return nil, err
	}
	return conn, nil
}

func Channel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		g.Log().Info("Failed to open a channel to RabbitMQ", zap.Error(err))
		return nil, err
	}
	return ch, nil
}

func ConnectionClose(conn *amqp.Connection) error {
	err := conn.Close()
	if err != nil {
		g.Log().Info("Failed to connect close to RabbitMQ", zap.Error(err))
		return err
	}
	return nil
}

func ChannelClose(ch *amqp.Channel) error {
	err := ch.Close()
	if err != nil {
		g.Log().Info("Failed to channel close to RabbitMQ", zap.Error(err))
		return err
	}
	return nil
}

func ExchangeDeclare(ch *amqp.Channel) {

}

func Panic(err error, msg string) {
	if err != nil {
		g.Log().Info("Failed to connect to RabbitMQ", zap.Error(err), zap.String("message", msg))
		panic("RabbitMQPanic")
	}
}
