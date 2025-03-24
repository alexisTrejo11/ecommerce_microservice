package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/repository"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
	"github.com/google/uuid"
)

type ModuleReceiver interface {
	ReceiveCreatemodule(ctx context.Context) error
	ReceiveUpdatemodule(ctx context.Context) error
	ReceiveDeletemodule(ctx context.Context) error
}

type moduleResult struct {
	module  *repository.ModuleDocument
	receipt string
	err     error
}

type moduleQueueReceiver struct {
	queueClient      QueueClient
	queueName        string
	timeout          int64
	moduleRepository repository.CourseRepository
}

func NewModuleQueueReceiver(
	queueClient QueueClient,
	queueName string,
	timeout int64,
	moduleRepository repository.CourseRepository,
) *moduleQueueReceiver {
	return &moduleQueueReceiver{
		queueClient:      queueClient,
		queueName:        queueName,
		timeout:          timeout,
		moduleRepository: moduleRepository,
	}
}

func (r *moduleQueueReceiver) ReceiveModule(ctx context.Context) error {
	logging.Logger.Info().Str("queue", r.queueName).Msg("Waiting for incoming module task from RabbitMQ")

	resultCh := make(chan moduleResult, 1)

	go r.receiveMessage(resultCh)

	select {
	case <-ctx.Done():
		logging.Logger.Warn().Str("queue", r.queueName).Msg("Context cancelled while waiting for module")
		ctx.Err()

	case result := <-resultCh:
		if result.err != nil {
			logging.LogError("receive_module", "Error receiving message", map[string]interface{}{
				"error": result.err.Error(),
			})
			return result.err
		}

		logging.Logger.Info().Str("queue", r.queueName).Str("module_id", result.module.ID).Msg("module message received")

		r.handleMessage(ctx, result, "create")
		return nil
	}

	return nil
}

func (r *moduleQueueReceiver) receiveMessage(resultCh chan<- moduleResult) {
	message, receiptHandle, err := r.queueClient.ReceiveMessage(r.queueName, time.Duration(r.timeout))
	if err != nil || len(message) == 0 {
		logging.LogError("receive_message", "Failed to receive message from queue", map[string]interface{}{
			"error": err,
		})
		resultCh <- moduleResult{nil, "", errors.New("failed to receive message or no messages available")}
		return
	}

	logging.Logger.Info().Str("queue", r.queueName).Str("receipt_handle", receiptHandle).Msg("Raw message received from queue")

	var module repository.ModuleDocument
	if err := json.Unmarshal(message, &module); err != nil {
		logging.LogError("unmarshal_message", "Failed to unmarshal module", map[string]interface{}{
			"error": err.Error(),
		})
		resultCh <- moduleResult{nil, receiptHandle, fmt.Errorf("failed to unmarshal module: %w", err)}
		return
	}

	if err := validatemodule(&module); err != nil {
		logging.LogError("validate_module", "Invalid module format", map[string]interface{}{
			"error": err.Error(),
		})
		resultCh <- moduleResult{nil, receiptHandle, err}
		return
	}

	logging.Logger.Info().Str("queue", r.queueName).Str("module_id", module.ID).Msg("module successfully parsed")

	resultCh <- moduleResult{&module, receiptHandle, nil}
}

func (r *moduleQueueReceiver) handleMessage(ctx context.Context, result moduleResult, action string) {
	logging.Logger.Info().Str("queue", r.queueName).Str("module_id", result.module.ID).Msg("Processing received module")

	if err := r.queueClient.DeleteMessage(r.queueName, result.receipt); err != nil {
		logging.LogError("delete_message", "Failed to delete message from queue", map[string]interface{}{
			"error": err.Error(),
			"queue": r.queueName,
		})
	} else {
		logging.Logger.Info().Str("queue", r.queueName).Str("receipt_handle", result.receipt).Msg("Message successfully deleted from queue")
	}

	switch action {
	case "CREATE", "UPDATE":
		err := r.moduleRepository.AddModule(ctx, uuid.MustParse(result.module.CourseID), *result.module)
		if err != nil {
			logging.LogError("create_persisted", "Failed to create module", map[string]interface{}{
				"error":     err.Error(),
				"module_id": result.module.ID,
			})
			return
		}
	case "DELETE":
		err := r.moduleRepository.DeleteModule(ctx, uuid.MustParse(result.module.ID))
		if err != nil {
			logging.LogError("module_deleted", "Failed to delete module", map[string]interface{}{
				"error":     err.Error(),
				"module_id": result.module.ID,
			})
			return
		}
	}

	logging.LogSuccess("process_module", "module Successfully Processed", map[string]interface{}{
		"module_id": result.module.ID,
		"action":    action,
	})
}

func validatemodule(module *repository.ModuleDocument) error {
	_, err := uuid.Parse(module.ID)
	if err != nil {
		return errors.New("module ID is invalid")
	}

	if module.Title == "" {
		return errors.New("module title is required")
	}

	if module.OrderNumber < 1 {
		return errors.New("module order must be greater than zero")
	}

	return nil
}
