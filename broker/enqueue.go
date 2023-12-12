package broker

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/BaseMax/RabbitMQOrderGo/models"
)

func EnqueueToRabbit[T models.Order | models.Refund](msg T, queueName string) error {
	var err error

	_, err = rCh.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = rCh.PublishWithContext(context.Background(), "", queueName, false, false, amqp.Publishing{
		ContentType: "json/application",
		Body:        body,
	})
	return err
}
