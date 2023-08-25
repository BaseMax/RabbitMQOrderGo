package broker

import (
	"encoding/json"
	"log"
	"time"

	"github.com/BaseMax/RabbitMQOrderGo/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

func browseLastAvalableOrder(msgs <-chan amqp.Delivery) *models.Order {
	for {
		var orderOnQ models.Order

		select {
		case m := <-msgs:
			err := json.Unmarshal(m.Body, &orderOnQ)
			if err != nil {
				log.Println("dequeue unmarshal:", err)
			}

			orderOnDB, err := models.GetOrderById(orderOnQ.ID)
			if err != nil {
				log.Println("dequeue last order:", err)
				m.Ack(false)
			}

			if orderOnDB.Status != models.ORDER_STATUS_PROCESSING {
				m.Ack(false)
			} else {
				m.Nack(false, true)
				return &orderOnDB
			}
		case <-time.After(time.Millisecond * 200):
			return nil
		}
	}
}

func DequeueLastOrder() (*models.Order, error) {
	msgs, err := rCh.Consume(rQ.Name, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	order := browseLastAvalableOrder(msgs)

	rCh.Close()
	rCh, err = rConn.Channel()
	return order, err
}
