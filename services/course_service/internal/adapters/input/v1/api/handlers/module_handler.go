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

type ModuleHandler struct {
	useCase   input.ModuleUseCase
	validator *validator.Validate
}

func NewModuleHandler(useCase input.ModuleUseCase) *ModuleHandler {
	return &ModuleHandler{
		useCase:   useCase,
		validator: validator.New(),
	}
}

func (lh *ModuleHandler) GetModuleById(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Module ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Module ID", "invalid id")
	}

	module, err := lh.useCase.GetModuleById(context.Background(), id)
	if err != nil {
		return response.NotFound(c, "Module not found", err.Error())
	}

	return response.OK(c, "Module Successfully Retrieved", module)
}

func (lh *ModuleHandler) GetModuleByCourseId(c *fiber.Ctx) error {
	idSTR := c.Params("course_id")
	if idSTR == "" {
		return response.BadRequest(c, "Course ID is mandatory", "course id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Course ID", "invalid id")
	}

	courses, err := lh.useCase.GetModuleByCourseId(context.Background(), id)
	if err != nil {
		return response.NotFound(c, "Modules not found", err.Error())
	}

	return response.OK(c, "Modules Successfully Retrieved", courses)
}

func (lh *ModuleHandler) CreateModule(c *fiber.Ctx) error {
	var insertDTO dtos.ModuleInsertDTO

	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	moduleCreated, err := lh.useCase.CreateModule(context.TODO(), insertDTO)
	if err != nil {
		return response.BadRequest(c, "Error creating module", err.Error())
	}

	return response.Created(c, "Module successfully created", moduleCreated)
}

func (lh *ModuleHandler) UpdateModule(c *fiber.Ctx) error {
	var insertDTO dtos.ModuleInsertDTO

	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Module ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Module ID", "invalid id")
	}

	if err := c.BodyParser(&insertDTO); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	moduleUpdated, err := lh.useCase.UpdateModule(context.TODO(), id, insertDTO)
	if err != nil {
		return response.BadRequest(c, "Error updating module", err.Error())
	}

	return response.OK(c, "Module successfully updated", moduleUpdated)
}

func (lh *ModuleHandler) DeleteModule(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return response.BadRequest(c, "Module ID is mandatory", "id is obligatory")
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return response.BadRequest(c, "Invalid Module ID", "invalid id")
	}

	err = lh.useCase.DeleteModule(context.Background(), id)
	if err != nil {
		return response.NotFound(c, "Module not found", err.Error())
	}

	return response.OK[any](c, "Module successfully deleted", nil)
}
