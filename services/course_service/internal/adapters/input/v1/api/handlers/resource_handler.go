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

// GetResourceById godoc
// @Summary      Get Resource by ID
// @Description  Retrieve a resource by its unique ID.
// @Tags         Resources
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Resource ID"
// @Success      200  {object}  response.ApiResponse{data=dtos.ResourceDTO} "Resource successfully retrieved"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Failure      404  {object}  response.ApiResponse "Resource not found"
// @Router       /v1/api/resources/{id} [get]
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

// GetResourceByLessonId godoc
// @Summary      Get Resources by Lesson ID
// @Description  Retrieve resources associated with a specific lesson.
// @Tags         Resources
// @Accept       json
// @Produce      json
// @Param        lesson_id   path      string  true  "Lesson ID"
// @Success      200  {object}  response.ApiResponse{data=[]dtos.ResourceDTO} "Resources successfully retrieved"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Failure      404  {object}  response.ApiResponse "Resources not found"
// @Router       /v1/api/lesson/{lesson_id} [get]
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

// CreateResource godoc
// @Summary      Create a new Resource
// @Description  Create a new resource for a lesson with the provided details.
// @Tags         Resources
// @Accept       json
// @Produce      json
// @Param        resource  body      dtos.ResourceInsertDTO  true  "Resource to create"
// @Success      201  {object}  response.ApiResponse{data=dtos.ResourceDTO} "Resource successfully created"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Router       /resources [post]
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

// UpdateResource godoc
// @Summary      Update an existing Resource
// @Description  Update the details of an existing resource identified by its ID.
// @Tags         Resources
// @Accept       json
// @Produce      json
// @Param        id        path      string                 true  "Resource ID"
// @Param        resource  body      dtos.ResourceInsertDTO true  "Resource data to update"
// @Success      200  {object}  response.ApiResponse{data=dtos.ResourceDTO} "Resource successfully updated"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Router       /v1/api/resources/{id} [put]
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

// DeleteResource godoc
// @Summary      Delete a Resource
// @Description  Delete an existing resource identified by its ID.
// @Tags         Resources
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Resource ID"
// @Success      200  {object}  response.ApiResponse "Resource successfully deleted"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Failure      404  {object}  response.ApiResponse "Resource not found"
// @Router       /v1/api/resources/{id} [delete]
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

	return response.OK(c, "Resource successfully deleted", nil)
}
