package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CourseHandler struct {
	useCase   input.CourseUseCase
	validator *validator.Validate
}

func NewCourseHandler(useCase input.CourseUseCase) *CourseHandler {
	return &CourseHandler{
		useCase:   useCase,
		validator: validator.New(),
	}
}

func (lh *CourseHandler) GetCourseById(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is obligatory"})
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	Course, err := lh.useCase.GetCourseById(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(Course)
}

func (lh *CourseHandler) CreateCourse(c *fiber.Ctx) error {
	var insertDTO dtos.CourseInsertDTO

	if err := c.BodyParser(&insertDTO); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errorsMap,
		})
	}

	CourseCreated, err := lh.useCase.CreateCourse(context.TODO(), insertDTO)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(201).JSON(CourseCreated)
}

func (lh *CourseHandler) UpdateCourse(c *fiber.Ctx) error {
	var insertDTO dtos.CourseInsertDTO

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

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"errors":  errorsMap,
		})
	}

	CourseUpdated, err := lh.useCase.UpdateCourse(context.TODO(), id, insertDTO)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(201).JSON(CourseUpdated)
}

func (lh *CourseHandler) DeleteCourse(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is obligatory"})
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	err = lh.useCase.DeleteCourse(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(204).JSON("deleted")
}
