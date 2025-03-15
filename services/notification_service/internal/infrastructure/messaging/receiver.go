package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/notification-service/internal/shared/dtos"
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

func NewQueueReceiver(queueClient QueueClient, queueName string, timeout time.Duration, notificationUseCase input.NotificationUseCase) *QueueReceiver {
	return &QueueReceiver{
		queueClient:         queueClient,
		queueName:           queueName,
		timeout:             timeout,
		notificationUseCase: notificationUseCase,
	}
}

func (r *QueueReceiver) ReceiveNotification(ctx context.Context) (*dtos.NotificationMessageDTO, error) {
	resultCh := make(chan notificationResult, 1)

	go r.receiveMessage(resultCh)

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case result := <-resultCh:
		if result.err != nil {
			return nil, result.err
		}

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
		resultCh <- notificationResult{nil, "", errors.New("failed to receive message or no messages available")}
		return
	}

	var notification dtos.NotificationMessageDTO
	if err := json.Unmarshal(message, &notification); err != nil {
		resultCh <- notificationResult{nil, receiptHandle, fmt.Errorf("failed to unmarshal notification: %w", err)}
		return
	}

	if err := validateNotification(&notification); err != nil {
		resultCh <- notificationResult{nil, receiptHandle, err}
		return
	}

	resultCh <- notificationResult{&notification, receiptHandle, nil}
}

func (r *QueueReceiver) handleMessage(ctx context.Context, result notificationResult) {
	if err := r.queueClient.DeleteMessage(r.queueName, result.receipt); err != nil {
		fmt.Printf("Error deleting message from queue: %v\n", err)
	}

	notification, err := r.notificationUseCase.CreateNotification(ctx, *result.notification)
	if err != nil {
		fmt.Printf("Can't create notification: %v\n", err)
	}

	fmt.Printf("Notification Successfully Created: ID:%v\n", notification.ID)
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
