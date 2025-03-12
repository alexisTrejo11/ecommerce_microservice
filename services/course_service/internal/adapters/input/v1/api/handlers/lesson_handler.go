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

	return response.OK(c, "Lesson Successfully Retrieved", lesson)
}

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

	return response.OK[any](c, "Lesson successfully deleted", nil)
}
