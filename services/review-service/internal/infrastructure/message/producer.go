package rabbitmq

import (
	"fmt"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/rating-service/internal/infrastructure/config"
)

func (r *RabbitMQClient) PublishCourseRatingUpdate(courseID string, newRating float64) error {
	message := map[string]interface{}{
		"course_id":  courseID,
		"rating":     newRating,
		"updated_at": time.Now().Format(time.RFC3339),
	}

	return r.PublishMessage(config.QueueUpdateCourseRating, message)
}

func (r *RabbitMQClient) DeleteMessage(queueName string, receiptHandle string) error {
	deliveryTag, err := parseDeliveryTag(receiptHandle)
	if err != nil {
		return fmt.Errorf("invalid receipt handle: %w", err)
	}
	return r.channel.Ack(deliveryTag, false)
}
