package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/response"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/utils"
	logging "github.com/alexisTrejo11/ecommerce_microservice/course-service/pkg/log"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type LessonHandler struct {
	useCase   input.LessonUseCase
	validator *validator.Validate
}

func NewLessonHandler(useCase input.LessonUseCase) *LessonHandler {
	return &LessonHandler{
		useCase:   useCase,
		validator: validator.New(),
	}
}

// GetLessonById godoc
// @Summary      Get Lesson by ID
// @Description  Retrieve a lesson by its unique ID.
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Lesson ID"
// @Success      200  {object}  response.ApiResponse{data=dtos.LessonDTO} "Lesson successfully retrieved"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Failure      404  {object}  response.ApiResponse "Lesson not found"
// @Router       /v1/api/lessons/{id} [get]
func (lh *LessonHandler) GetLessonById(c *fiber.Ctx) error {
	// Log incoming request
	logging.Logger.WithFields(logrus.Fields{
		"action": "get_lesson_by_id",
		"method": c.Method(),
		"route":  c.Route().Path,
		"ip":     c.IP(),
		"param":  c.Params("id"),
		"user":   c.Locals("user_id"),
	}).Info("Incoming request")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "get_lesson_by_id",
			"error":  err.Error(),
		}).Error("Invalid lesson ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	lesson, err := lh.useCase.GetLessonById(context.Background(), id)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":    "get_lesson_by_id",
			"lesson_id": id,
			"error":     err.Error(),
		}).Error("Lesson not found")
		return response.NotFound(c, "Lesson not found", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":    "get_lesson_by_id",
		"lesson_id": id,
	}).Info("Lesson successfully retrieved")

	return response.OK(c, "Lesson successfully retrieved", lesson)
}

// CreateLesson godoc
// @Summary      Create a new Lesson
// @Description  Create a new lesson with the provided details.
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Param        lesson  body      dtos.LessonInsertDTO  true  "Lesson to create"
// @Success      201  {object}  response.ApiResponse{data=dtos.LessonDTO} "Lesson successfully created"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Router       /v1/api/lessons [post]
func (lh *LessonHandler) CreateLesson(c *fiber.Ctx) error {
	// Log incoming request with payload
	logging.Logger.WithFields(logrus.Fields{
		"action":  "create_lesson",
		"method":  c.Method(),
		"route":   c.Route().Path,
		"ip":      c.IP(),
		"user":    c.Locals("user_id"),
		"payload": c.Body(),
	}).Info("Incoming request")

	var insertDTO dtos.LessonInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "create_lesson",
			"error":  err.Error(),
		}).Error("Invalid request body")
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "create_lesson",
			"error":  err.Error(),
		}).Error("Validation failed")
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	lessonCreated, err := lh.useCase.CreateLesson(context.TODO(), insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "create_lesson",
			"error":  err.Error(),
		}).Error("Error creating lesson")
		return response.BadRequest(c, "Error creating lesson", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":    "create_lesson",
		"lesson_id": lessonCreated.ID,
	}).Info("Lesson successfully created")

	return response.Created(c, "Lesson successfully created", lessonCreated)
}

// UpdateLesson godoc
// @Summary      Update an existing Lesson
// @Description  Update the details of an existing lesson identified by its ID.
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Param        id      path      string               true  "Lesson ID"
// @Param        lesson  body      dtos.LessonInsertDTO true  "Lesson data to update"
// @Success      200     {object}  response.ApiResponse{data=dtos.LessonDTO} "Lesson successfully updated"
// @Failure      400     {object}  response.ApiResponse "Bad Request"
// @Router       /v1/api/lessons/{id} [put]
func (lh *LessonHandler) UpdateLesson(c *fiber.Ctx) error {
	// Log incoming request with payload
	logging.Logger.WithFields(logrus.Fields{
		"action":  "update_lesson",
		"method":  c.Method(),
		"route":   c.Route().Path,
		"ip":      c.IP(),
		"user":    c.Locals("user_id"),
		"payload": c.Body(),
	}).Info("Incoming request")

	var insertDTO dtos.LessonInsertDTO

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "update_lesson",
			"error":  err.Error(),
		}).Error("Invalid lesson ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	if err := c.BodyParser(&insertDTO); err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "update_lesson",
			"error":  err.Error(),
		}).Error("Invalid request body")
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "update_lesson",
			"error":  err.Error(),
		}).Error("Validation failed")
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	lessonUpdated, err := lh.useCase.UpdateLesson(context.TODO(), id, insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":    "update_lesson",
			"lesson_id": id,
			"error":     err.Error(),
		}).Error("Error updating lesson")
		return response.BadRequest(c, "Error updating lesson", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":    "update_lesson",
		"lesson_id": id,
	}).Info("Lesson successfully updated")

	return response.OK(c, "Lesson successfully updated", lessonUpdated)
}

// DeleteLesson godoc
// @Summary      Delete a Lesson
// @Description  Delete an existing lesson identified by its ID.
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Lesson ID"
// @Success      200  {object}  response.ApiResponse "Lesson successfully deleted"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Failure      404  {object}  response.ApiResponse "Lesson not found"
// @Router       /v1/api/lessons/{id} [delete]
func (lh *LessonHandler) DeleteLesson(c *fiber.Ctx) error {
	// Log incoming request for deletion
	logging.Logger.WithFields(logrus.Fields{
		"action": "delete_lesson",
		"method": c.Method(),
		"route":  c.Route().Path,
		"ip":     c.IP(),
		"user":   c.Locals("user_id"),
		"param":  c.Params("id"),
	}).Info("Incoming request")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "delete_lesson",
			"error":  err.Error(),
		}).Error("Invalid lesson ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	err = lh.useCase.DeleteLesson(context.Background(), id)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":    "delete_lesson",
			"lesson_id": id,
			"error":     err.Error(),
		}).Error("Lesson not found")
		return response.NotFound(c, "Lesson not found", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":    "delete_lesson",
		"lesson_id": id,
	}).Info("Lesson successfully deleted")

	return response.OK(c, "Lesson successfully deleted", nil)
}
