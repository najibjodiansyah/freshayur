package rabbitmq

import (
	"encoding/json"

	"github.com/najibjodiansyah/freshayur/notification/internal/core"
	"github.com/streadway/amqp"
)

type Consumer interface {
	ConsumeMessage(qName string, req core.NotificationEntity) error
}

type consume struct {
	*amqp.Connection
}

func NewConcumer(conn *amqp.Connection) Consumer {
	return &consume{
		Connection: conn,
	}
}

func (c *consume) ConsumeMessage(qName string, req core.NotificationEntity) error {
	ch, err := c.Connection.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	msgs, err := ch.Consume(
		qName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		// Process the message
		_ = msg // Placeholder for message processing logic

		if err := json.Unmarshal(msg.Body, &req); err != nil {
			continue
		}
	}

	return nil
}
