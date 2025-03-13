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
	logging.LogIncomingRequest(c, "get_lesson_by_id")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("get_lesson_by_id", "Invalid lesson ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	lesson, err := lh.useCase.GetLessonById(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "delete_course", id.String())
	}

	logging.LogSuccess("get_lesson_by_id", "Lesson successfully updated", map[string]interface{}{
		"lesson_id": id,
	})

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
	logging.LogIncomingRequest(c, "create_lesson")

	var insertDTO dtos.LessonInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		logging.LogError("create_lesson", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.LogError("create_lesson", "invalid request data", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	lessonCreated, err := lh.useCase.CreateLesson(context.TODO(), insertDTO)
	if err != nil {
		return response.HandleApplicationError(c, err, "create_course", lessonCreated.ID.String())
	}

	logging.LogSuccess("update_lesson", "Lesson successfully updated", map[string]interface{}{
		"lesson_id": lessonCreated.ID,
	})

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
	logging.LogIncomingRequest(c, "update_lesson")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("update_course", "Invalid lesson ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	var insertDTO dtos.LessonInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		logging.LogError("update_lesson", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.LogError("update_lesson", "invalid request data", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	lessonUpdated, err := lh.useCase.UpdateLesson(context.TODO(), id, insertDTO)
	if err != nil {
		return response.HandleApplicationError(c, err, "update_course", id.String())
	}

	logging.LogSuccess("update_lesson", "Lesson successfully updated", map[string]interface{}{
		"lesson_id": id,
	})

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
	logging.LogIncomingRequest(c, "delete_lesson")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("delete_course", "Invalid lesson ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	err = lh.useCase.DeleteLesson(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "delete_lesson", id.String())
	}

	logging.LogSuccess("delete_lesson", "Lesson successfully deleted", map[string]interface{}{
		"lesson_id": id,
	})

	return response.OK(c, "Lesson successfully deleted", nil)
}
