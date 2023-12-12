package broker

import (
	"github.com/BaseMax/RabbitMQOrderGo/conf"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	QUEHE_NAME_ORDERS  = "orders"
	QUEHE_NAME_REFUNDS = "refunds"
)

var (
	rConn *amqp.Connection
	rCh   *amqp.Channel
)

func ConnectAndCreateQueue() error {
	var err error
	rConn, err = amqp.Dial(conf.GetRabbitUrl())
	if err != nil {
		rConn = nil
		return err
	}
	rCh, err = rConn.Channel()
	if err != nil {
		rConn = nil
		rCh = nil
		return err
	}
	return nil
}

func IsClosed() bool {
	if rConn == nil || rCh == nil {
		return true
	}
	return rConn.IsClosed()
}

func GetStatus() string {
	ConnectAndCreateQueue()
	if IsClosed() {
		return "down"
	}
	return "up"
}
