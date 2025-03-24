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

type CourseReceiver interface {
	ReceiveCreateCourse(ctx context.Context) error
	ReceiveUpdateCourse(ctx context.Context) error
	ReceiveDeleteCourse(ctx context.Context) error
}

type courseResult struct {
	Course  *repository.CourseDocument
	receipt string
	err     error
}

func NewCourseQueueReceiver(
	queueClient QueueClient,
	queueName string,
	timeout time.Duration,
	courseRepository repository.CourseRepository) *CourseQueueReceiver {
	return &CourseQueueReceiver{
		queueClient:      queueClient,
		queueName:        queueName,
		timeout:          timeout,
		courseRepository: courseRepository,
	}
}

func (r *CourseQueueReceiver) ReceiveCourse(ctx context.Context) error {
	logging.Logger.Info().Str("queue", r.queueName).Msg("Waiting for incoming task from RabbitMQ")

	resultCh := make(chan courseResult, 1)

	go r.receiveMessage(resultCh)

	select {
	case <-ctx.Done():
		logging.Logger.Warn().Str("queue", r.queueName).Msg("Context cancelled while waiting for course")
		ctx.Err()

	case result := <-resultCh:
		if result.err != nil {
			logging.LogError("receive_course", "Error receiving message", map[string]interface{}{
				"error": result.err.Error(),
			})
			return result.err
		}

		logging.Logger.Info().Str("queue", r.queueName).Str("course_id", result.Course.ID).Msg("course message received")

		r.handleMessage(ctx, result, "create")
		return nil
	}

	return nil
}

func (r *CourseQueueReceiver) receiveMessage(resultCh chan<- courseResult) {
	message, receiptHandle, err := r.queueClient.ReceiveMessage(r.queueName, r.timeout)
	if err != nil || len(message) == 0 {
		logging.LogError("receive_message", "Failed to receive message from queue", map[string]interface{}{
			"error": err,
		})
		resultCh <- courseResult{nil, "", errors.New("failed to receive message or no messages available")}
		return
	}

	logging.Logger.Info().Str("queue", r.queueName).Str("receipt_handle", receiptHandle).Msg("Raw message received from queue")

	var course repository.CourseDocument
	if err := json.Unmarshal(message, &course); err != nil {
		logging.LogError("unmarshal_message", "Failed to unmarshal course", map[string]interface{}{
			"error": err.Error(),
		})
		resultCh <- courseResult{nil, receiptHandle, fmt.Errorf("failed to unmarshal course: %w", err)}
		return
	}

	if err := validateCourse(&course); err != nil {
		logging.LogError("validate_course", "Invalid course format", map[string]interface{}{
			"error": err.Error(),
		})
		resultCh <- courseResult{nil, receiptHandle, err}
		return
	}

	logging.Logger.Info().Str("queue", r.queueName).Str("course_id", course.ID).Msg("course successfully parsed")

	resultCh <- courseResult{&course, receiptHandle, nil}
}

func (r *CourseQueueReceiver) handleMessage(ctx context.Context, result courseResult, action string) {
	logging.Logger.Info().Str("queue", r.queueName).Str("course_id", result.Course.ID).Msg("Processing received course")

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

		doc, _ := r.courseRepository.ToCourseEntity(*result.Course)
		err := r.courseRepository.Save(ctx, doc)
		if err != nil {
			logging.LogError("create_perisisted", "Failed to create course", map[string]interface{}{
				"error":     err.Error(),
				"course_id": result.Course.ID,
			})
			return
		}
	case "DELETE":
		err := r.courseRepository.Delete(ctx, uuid.MustParse(result.Course.ID))
		if err != nil {
			logging.LogError("course_deleted", "Failed to create course", map[string]interface{}{
				"error":     err.Error(),
				"course_id": result.Course.ID,
			})
			return
		}

	}

	logging.LogSuccess("create_course", "course Successfully Created", map[string]interface{}{
		"course_id": result.Course.ID,
	})
}

func validateCourse(course *repository.CourseDocument) error {
	_, err := uuid.Parse(course.ID)
	if err != nil {
		return errors.New("course ID is invalid")
	}

	if course.Name == "" {
		return errors.New("course name is required")
	}

	if course.Category == "" {
		return errors.New("course category is required")
	}

	if course.Level == "" {
		return errors.New("course level is required")
	}

	if course.InstructorID == "" {
		return errors.New("instructor ID is required")
	}

	if course.Language == "" {
		return errors.New("course language is required")
	}

	if course.ThumbnailURL == "" {
		fmt.Println("Warning: Thumbnail URL is missing but not required.")
	}

	if len(course.Modules) == 0 {
		fmt.Println("Warning: No modules found for the course.")
	}

	return nil
}
