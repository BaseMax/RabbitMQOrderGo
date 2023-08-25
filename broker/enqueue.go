package broker

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/BaseMax/RabbitMQOrderGo/models"
)

func EnqueueOrderToRabbit(order models.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = rCh.PublishWithContext(context.Background(), "", rQ.Name, false, false, amqp.Publishing{
		ContentType: "json/application",
		Body:        body,
	})
	return err
}
