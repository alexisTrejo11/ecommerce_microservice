package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
	logging "github.com/alexisTrejo11/ecommerce_microservice/notification-service/pkg/log"
)

type NotificationReceiver interface {
	ReceiveNotification(ctx context.Context) (*dtos.NotificationMessageDTO, error)
}

type QueueReceiver struct {
	queueClient         QueueClient
	queueName           string
	notificationUseCase input.NotificationUseCase
	timeout             time.Duration
}

type QueueClient interface {
	ReceiveMessage(queueName string, timeout time.Duration) ([]byte, string, error)
	DeleteMessage(queueName string, receiptHandle string) error
}

func NewQueueReceiver(
	queueClient QueueClient,
	queueName string,
	timeout time.Duration,
	notificationUseCase input.NotificationUseCase) *QueueReceiver {
	return &QueueReceiver{
		queueClient:         queueClient,
		queueName:           queueName,
		timeout:             timeout,
		notificationUseCase: notificationUseCase,
	}
}

func (r *QueueReceiver) ReceiveNotification(ctx context.Context) (*dtos.NotificationMessageDTO, error) {
	logging.Logger.Info().Str("queue", r.queueName).Msg("Waiting for incoming task from RabbitMQ")

	resultCh := make(chan notificationResult, 1)

	go r.receiveMessage(resultCh)

	select {
	case <-ctx.Done():
		logging.Logger.Warn().Str("queue", r.queueName).Msg("Context cancelled while waiting for notification")
		return nil, ctx.Err()

	case result := <-resultCh:
		if result.err != nil {
			logging.LogError("receive_notification", "Error receiving message", map[string]interface{}{
				"error": result.err.Error(),
			})
			return nil, result.err
		}

		logging.Logger.Info().Str("queue", r.queueName).Str("notification_id", result.notification.ID).Msg("Notification message received")

		r.handleMessage(ctx, result)
		return result.notification, nil
	}
}

type notificationResult struct {
	notification *dtos.NotificationMessageDTO
	receipt      string
	err          error
}

func (r *QueueReceiver) receiveMessage(resultCh chan<- notificationResult) {
	message, receiptHandle, err := r.queueClient.ReceiveMessage(r.queueName, r.timeout)
	if err != nil || len(message) == 0 {
		logging.LogError("receive_message", "Failed to receive message from queue", map[string]interface{}{
			"error": err,
		})
		resultCh <- notificationResult{nil, "", errors.New("failed to receive message or no messages available")}
		return
	}

	logging.Logger.Info().Str("queue", r.queueName).Str("receipt_handle", receiptHandle).Msg("Raw message received from queue")

	var notification dtos.NotificationMessageDTO
	if err := json.Unmarshal(message, &notification); err != nil {
		logging.LogError("unmarshal_message", "Failed to unmarshal notification", map[string]interface{}{
			"error": err.Error(),
		})
		resultCh <- notificationResult{nil, receiptHandle, fmt.Errorf("failed to unmarshal notification: %w", err)}
		return
	}

	if err := validateNotification(&notification); err != nil {
		logging.LogError("validate_notification", "Invalid notification format", map[string]interface{}{
			"error": err.Error(),
		})
		resultCh <- notificationResult{nil, receiptHandle, err}
		return
	}

	logging.Logger.Info().Str("queue", r.queueName).Str("notification_id", notification.ID).Msg("Notification successfully parsed")

	resultCh <- notificationResult{&notification, receiptHandle, nil}
}

func (r *QueueReceiver) handleMessage(ctx context.Context, result notificationResult) {
	logging.Logger.Info().Str("queue", r.queueName).Str("notification_id", result.notification.ID).Msg("Processing received notification")

	if err := r.queueClient.DeleteMessage(r.queueName, result.receipt); err != nil {
		logging.LogError("delete_message", "Failed to delete message from queue", map[string]interface{}{
			"error": err.Error(),
			"queue": r.queueName,
		})
	} else {
		logging.Logger.Info().Str("queue", r.queueName).Str("receipt_handle", result.receipt).Msg("Message successfully deleted from queue")
	}

	notification, err := r.notificationUseCase.CreateNotification(ctx, *result.notification)
	if err != nil {
		logging.LogError("create_notification", "Failed to create notification", map[string]interface{}{
			"error":           err.Error(),
			"notification_id": result.notification.ID,
		})
		return
	}

	logging.LogSuccess("create_notification", "Notification Successfully Created", map[string]interface{}{
		"notification_id": notification.ID,
	})
}

func validateNotification(notification *dtos.NotificationMessageDTO) error {
	if notification.ID == "" {
		return errors.New("notification ID is required")
	}
	if notification.Type == "" {
		return errors.New("notification type is required")
	}
	if notification.Content == "" {
		return errors.New("notification content is required")
	}
	if notification.UserData.ID == "" {
		return errors.New("user ID is required")
	}

	switch notification.Type {
	case "email":
		if notification.UserData.Email == "" {
			return errors.New("email is required for email notifications")
		}
	case "sms":
		if notification.UserData.Phone == nil || *notification.UserData.Phone == "" {
			return errors.New("phone number is required for SMS notifications")
		}
	}

	return nil
}
