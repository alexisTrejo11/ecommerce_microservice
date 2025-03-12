package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/response"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Lesson ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Lesson ID", "invalid id")
	}

	lesson, err := lh.useCase.GetLessonById(context.Background(), id)
	if err != nil {
		return response.NotFound(c, "Lesson not found", err.Error())
	}

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
// @Router       /lessons [post]
func (lh *LessonHandler) CreateLesson(c *fiber.Ctx) error {
	var insertDTO dtos.LessonInsertDTO

	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	lessonCreated, err := lh.useCase.CreateLesson(context.TODO(), insertDTO)
	if err != nil {
		return response.BadRequest(c, "Error creating lesson", err.Error())
	}

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
	var insertDTO dtos.LessonInsertDTO

	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Lesson ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Lesson ID", "invalid id")
	}

	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	lessonUpdated, err := lh.useCase.UpdateLesson(context.TODO(), id, insertDTO)
	if err != nil {
		return response.BadRequest(c, "Error updating lesson", err.Error())
	}

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
	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Lesson ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Lesson ID", "invalid id")
	}

	err = lh.useCase.DeleteLesson(context.Background(), id)
	if err != nil {
		return response.NotFound(c, "Lesson not found", err.Error())
	}

	return response.OK(c, "Lesson successfully deleted", nil)
}
