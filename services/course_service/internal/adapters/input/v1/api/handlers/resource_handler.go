package handlers

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ResourceHandler struct {
	useCase   input.ResourceUseCase
	validator validator.Validate
}

func NewResourceHandler(useCase input.ResourceUseCase) *ResourceHandler {
	return &ResourceHandler{
		useCase:   useCase,
		validator: *validator.New(),
	}
}

func (lh *ResourceHandler) GetResourceById(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is obligatory"})
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	Resource, err := lh.useCase.GetResourceById(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(Resource)
}

func (lh *ResourceHandler) CreateResource(c *fiber.Ctx) error {
	var insertDTO dtos.ResourceInsertDTO

	if err := c.BodyParser(&insertDTO); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := lh.validator.Struct(&insertDTO); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	ResourceCreated, err := lh.useCase.CreateResource(context.TODO(), insertDTO)
	if err != nil {
		return c.Status(400).JSON(ResourceCreated)
	}

	return c.Status(201).JSON(ResourceCreated)
}

func (lh *ResourceHandler) UpdateResource(c *fiber.Ctx) error {
	var insertDTO dtos.ResourceInsertDTO

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

	ResourceUpdated, err := lh.useCase.UpdateResource(context.TODO(), id, insertDTO)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(201).JSON(ResourceUpdated)
}

func (lh *ResourceHandler) DeleteResource(c *fiber.Ctx) error {
	idSTR := c.Params("id")
	if idSTR == "" {
		return c.Status(400).JSON(fiber.Map{"error": "id is obligatory"})
	}

	id, err := uuid.Parse(idSTR)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	err = lh.useCase.DeleteResource(context.Background(), id)
	if err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(204).JSON("deleted")
}
