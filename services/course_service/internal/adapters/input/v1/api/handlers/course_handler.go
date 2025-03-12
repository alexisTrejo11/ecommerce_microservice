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
		return response.BadRequest(c, "Course ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Course ID", "invalid id")
	}

	course, err := lh.useCase.GetCourseById(context.Background(), id)
	if err != nil {
		return response.NotFound(c, "Course not found", err.Error())
	}

	return response.OK(c, "Course Successfully Retrieved", course)
}

func (lh *CourseHandler) CreateCourse(c *fiber.Ctx) error {
	var insertDTO dtos.CourseInsertDTO

	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	CourseCreated, err := lh.useCase.CreateCourse(context.TODO(), insertDTO)
	if err != nil {
		return response.BadRequest(c, "Error creating course", err.Error())
	}

	return response.Created(c, "Course successfully created", CourseCreated)
}

func (lh *CourseHandler) UpdateCourse(c *fiber.Ctx) error {
	var insertDTO dtos.CourseInsertDTO

	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Course ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Course ID", "invalid id")
	}

	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	CourseUpdated, err := lh.useCase.UpdateCourse(context.TODO(), id, insertDTO)
	if err != nil {
		return response.BadRequest(c, "Error updating course", err.Error())
	}

	return response.OK(c, "Course successfully updated", CourseUpdated)
}

func (lh *CourseHandler) DeleteCourse(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Course ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Course ID", "invalid id")
	}

	err = lh.useCase.DeleteCourse(context.Background(), id)
	if err != nil {
		return response.NotFound(c, "Course not found", err.Error())
	}

	return response.OK[any](c, "Course successfully deleted", nil)
}
