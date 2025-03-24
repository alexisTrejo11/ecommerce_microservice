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

type LessonReceiver interface {
	ReceiveCreateLesson(ctx context.Context) error
	ReceiveUpdateLesson(ctx context.Context) error
	ReceiveDeleteLesson(ctx context.Context) error
}

type lessonResult struct {
	Lesson  *repository.LessonDocument
	receipt string
	err     error
}

type LessonQueueReceiver struct {
	queueClient      QueueClient
	queueName        string
	timeout          int64
	lessonRepository repository.CourseRepository
}

func NewLessonQueueReceiver(
	queueClient QueueClient,
	queueName string,
	timeout int64,
	lessonRepository repository.CourseRepository,
) *LessonQueueReceiver {
	return &LessonQueueReceiver{
		queueClient:      queueClient,
		queueName:        queueName,
		timeout:          timeout,
		lessonRepository: lessonRepository,
	}
}

func (r *LessonQueueReceiver) ReceiveLesson(ctx context.Context) error {
	logging.Logger.Info().Str("queue", r.queueName).Msg("Waiting for incoming lesson task from RabbitMQ")

	resultCh := make(chan lessonResult, 1)

	go r.receiveMessage(resultCh)

	select {
	case <-ctx.Done():
		logging.Logger.Warn().Str("queue", r.queueName).Msg("Context cancelled while waiting for lesson")
		ctx.Err()

	case result := <-resultCh:
		if result.err != nil {
			logging.LogError("receive_lesson", "Error receiving message", map[string]interface{}{
				"error": result.err.Error(),
			})
			return result.err
		}

		logging.Logger.Info().Str("queue", r.queueName).Str("lesson_id", result.Lesson.ID).Msg("lesson message received")

		r.handleMessage(ctx, result, "create")
		return nil
	}

	return nil
}

func (r *LessonQueueReceiver) receiveMessage(resultCh chan<- lessonResult) {
	message, receiptHandle, err := r.queueClient.ReceiveMessage(r.queueName, time.Duration(r.timeout))
	if err != nil || len(message) == 0 {
		logging.LogError("receive_message", "Failed to receive message from queue", map[string]interface{}{
			"error": err,
		})
		resultCh <- lessonResult{nil, "", errors.New("failed to receive message or no messages available")}
		return
	}

	logging.Logger.Info().Str("queue", r.queueName).Str("receipt_handle", receiptHandle).Msg("Raw message received from queue")

	var lesson repository.LessonDocument
	if err := json.Unmarshal(message, &lesson); err != nil {
		logging.LogError("unmarshal_message", "Failed to unmarshal lesson", map[string]interface{}{
			"error": err.Error(),
		})
		resultCh <- lessonResult{nil, receiptHandle, fmt.Errorf("failed to unmarshal lesson: %w", err)}
		return
	}

	if err := validateLesson(&lesson); err != nil {
		logging.LogError("validate_lesson", "Invalid lesson format", map[string]interface{}{
			"error": err.Error(),
		})
		resultCh <- lessonResult{nil, receiptHandle, err}
		return
	}

	logging.Logger.Info().Str("queue", r.queueName).Str("lesson_id", lesson.ID).Msg("lesson successfully parsed")

	resultCh <- lessonResult{&lesson, receiptHandle, nil}
}

func (r *LessonQueueReceiver) handleMessage(ctx context.Context, result lessonResult, action string) {
	logging.Logger.Info().Str("queue", r.queueName).Str("lesson_id", result.Lesson.ID).Msg("Processing received lesson")

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
		err := r.lessonRepository.AddLesson(ctx, uuid.MustParse(result.Lesson.CourseID), *result.Lesson)
		if err != nil {
			logging.LogError("create_persisted", "Failed to create lesson", map[string]interface{}{
				"error":     err.Error(),
				"lesson_id": result.Lesson.ID,
			})
			return
		}
	case "DELETE":
		err := r.lessonRepository.DeleteLesson(ctx, uuid.MustParse(result.Lesson.ID))
		if err != nil {
			logging.LogError("lesson_deleted", "Failed to delete lesson", map[string]interface{}{
				"error":     err.Error(),
				"lesson_id": result.Lesson.ID,
			})
			return
		}
	}

	logging.LogSuccess("process_lesson", "Lesson Successfully Processed", map[string]interface{}{
		"lesson_id": result.Lesson.ID,
		"action":    action,
	})
}

func validateLesson(lesson *repository.LessonDocument) error {
	_, err := uuid.Parse(lesson.ID)
	if err != nil {
		return errors.New("lesson ID is invalid")
	}

	if lesson.Title == "" {
		return errors.New("lesson title is required")
	}

	if lesson.ModuleID == "" {
		return errors.New("module ID is required")
	}

	if lesson.OrderNumber < 1 {
		return errors.New("lesson order must be greater than zero")
	}

	if lesson.ContentType == "" {
		return errors.New("lesson type is required")
	}

	if lesson.Content == "" && lesson.ContentType != "quiz" {
		fmt.Println("Warning: Content URL is missing for non-quiz lesson.")
	}

	return nil
}
