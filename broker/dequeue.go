package broker

import (
	"encoding/json"
	"log"
	"time"

	"github.com/BaseMax/RabbitMQOrderGo/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

func processFirstAvalableOrder(msgs <-chan amqp.Delivery, processFirstOrder bool) *models.Order {
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
				if processFirstOrder {
					models.UpdateOrder(orderOnDB.ID, "", models.ORDER_STATUS_COMPLETED)
					m.Ack(false)
				} else {
					m.Nack(false, true)
				}
				return &orderOnDB
			}
		case <-time.After(time.Millisecond * 200):
			return nil
		}
	}
}

func DequeueFirstOrder(processFirstOrder bool) (*models.Order, error) {
	msgs, err := rCh.Consume(rQ.Name, "", false, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	order := processFirstAvalableOrder(msgs, processFirstOrder)

	rCh.Close()
	rCh, err = rConn.Channel()
	return order, err
}
