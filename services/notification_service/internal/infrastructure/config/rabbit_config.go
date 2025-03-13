package config

import (
	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Connection, error) {
	return amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
}
