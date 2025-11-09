package rabbitmq

import (
	"fmt"

	"github.com/najibjodiansyah/freshayur/notification/internal"
	"github.com/streadway/amqp"
)

func NewConnection(cfg internal.Config) (*amqp.Connection, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
