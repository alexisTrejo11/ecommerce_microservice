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

type ResourceHandler struct {
	useCase   input.ResourceUseCase
	validator *validator.Validate
}

func NewResourceHandler(useCase input.ResourceUseCase) *ResourceHandler {
	return &ResourceHandler{
		useCase:   useCase,
		validator: validator.New(),
	}
}

func (lh *ResourceHandler) GetResourceById(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Resource ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Resource ID", "invalid id")
	}

	resource, err := lh.useCase.GetResourceById(context.Background(), id)
	if err != nil {
		return response.NotFound(c, "Resource not found", err.Error())
	}

	return response.OK(c, "Resource successfully retrieved", resource)
}

func (lh *ResourceHandler) GetResourceByLessonId(c *fiber.Ctx) error {
	lessonIdSTR := c.Params("lesson_id")
	if lessonIdSTR == "" {
		return response.BadRequest(c, "Lesson ID is mandatory", "id is obligatory")
	}

	lessonId, err := uuid.Parse(lessonIdSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Lesson ID", "invalid lesson id")
	}

	resource, err := lh.useCase.GetResourcesByLessonId(context.Background(), lessonId)
	if err != nil {
		return response.NotFound(c, "Resource not found", err.Error())
	}

	return response.OK(c, "Resource successfully retrieved", resource)
}

func (lh *ResourceHandler) CreateResource(c *fiber.Ctx) error {
	var insertDTO dtos.ResourceInsertDTO

	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	resourceCreated, err := lh.useCase.CreateResource(context.TODO(), insertDTO)
	if err != nil {
		return response.BadRequest(c, "Error creating resource", err.Error())
	}

	return response.Created(c, "Resource successfully created", resourceCreated)
}

func (lh *ResourceHandler) UpdateResource(c *fiber.Ctx) error {
	var insertDTO dtos.ResourceInsertDTO

	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Resource ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Resource ID", "invalid id")
	}

	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	resourceUpdated, err := lh.useCase.UpdateResource(context.TODO(), id, insertDTO)
	if err != nil {
		return response.BadRequest(c, "Error updating resource", err.Error())
	}

	return response.OK(c, "Resource successfully updated", resourceUpdated)
}

func (lh *ResourceHandler) DeleteResource(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Resource ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Resource ID", "invalid id")
	}

	err = lh.useCase.DeleteResource(context.Background(), id)
	if err != nil {
		return response.NotFound(c, "Resource not found", err.Error())
	}

	return response.OK[any](c, "Resource successfully deleted", nil)
}
