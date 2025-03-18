package rabbitmq

import (
	"errors"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func (r *RabbitMQClient) ReceiveMessage(queueName string, timeout time.Duration) ([]byte, string, error) {
	msgs, err := r.channel.Consume(
		queueName,
		"",    // Consumer (vacío para que RabbitMQ asigne uno)
		false, // AutoAck (false para manejar confirmación manual)
		false, // Exclusive
		false, // NoLocal
		false, // NoWait
		nil,   // Args
	)
	if err != nil {
		return nil, "", fmt.Errorf("failed to consume messages: %w", err)
	}

	messageCh := make(chan amqp.Delivery)
	go func() {
		for msg := range msgs {
			messageCh <- msg
			return
		}
	}()

	select {
	case msg := <-messageCh:
		return msg.Body, fmt.Sprintf("%d", msg.DeliveryTag), nil
	case <-time.After(timeout):
		return nil, "", errors.New("timeout waiting for message")
	}
}
