package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LessonHandler struct {
	useCase   input.LessonUseCase
	validator validator.Validate
}

func NewLessonHandler(useCase input.LessonUseCase) *LessonHandler {
	return &LessonHandler{
		useCase:   useCase,
		validator: *validator.New(),
	}
}

func (lh *LessonHandler) GetLessonById(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is obligatory"})
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	lesson, err := lh.useCase.GetLessonById(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(lesson)
}

func (lh *LessonHandler) CreateLesson(c *fiber.Ctx) error {
	var insertDTO dtos.LessonInsertDTO

	if err := c.BodyParser(&insertDTO); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := lh.validator.Struct(&insertDTO); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	lessonCreated, err := lh.useCase.CreateLesson(context.TODO(), insertDTO)
	if err != nil {
		return c.Status(400).JSON(lessonCreated)
	}

	return c.Status(201).JSON(lessonCreated)
}

func (lh *LessonHandler) UpdateLesson(c *fiber.Ctx) error {
	var insertDTO dtos.LessonInsertDTO

	idSTR := c.Params("id")
	if idSTR == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is obligatory"})
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	if err := c.BodyParser(&insertDTO); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := lh.validator.Struct(&insertDTO); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	lessonUpdated, err := lh.useCase.UpdateLesson(context.TODO(), id, insertDTO)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(201).JSON(lessonUpdated)
}

func (lh *LessonHandler) DeleteLession(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is obligatory"})
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	err = lh.useCase.DeleteLesson(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(204).JSON("deleted")
}
