package messages

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ interface {
	Publish(ctx context.Context, exchange, routingKey string, body []byte) error
	PublishJSON(ctx context.Context, exchange, routingKey string, data interface{}) error
	DeclareQueue(name string) (amqp091.Queue, error)
	ConsumeMessages(queueName string) (<-chan amqp091.Delivery, error)
	Close()
}
