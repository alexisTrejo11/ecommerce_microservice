package rabbitmq

import (
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/repository"
)

type CourseQueueReceiver struct {
	queueClient      QueueClient
	queueName        string
	courseRepository repository.CourseRepository
	timeout          time.Duration
}

type QueueClient interface {
	ReceiveMessage(queueName string, timeout time.Duration) ([]byte, string, error)
	DeleteMessage(queueName string, receiptHandle string) error
}
