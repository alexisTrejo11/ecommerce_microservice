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

// ResourceHandler handles Resource-related endpoints.
type ResourceHandler struct {
	useCase   input.ResourceUseCase
	validator *validator.Validate
}

// NewResourceHandler creates a new ResourceHandler.
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
	logging.LogIncomingRequest(c, "get_resource_by_id")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("get_resource_by_id", "invalid resource ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid resource ID")
	}

	resource, err := lh.useCase.GetResourceById(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "get_resource_by_id", id.String())
	}

	logging.LogSuccess("get_resource_by_id", "Module successfully deleted", map[string]interface{}{
		"resource_id": id,
	})

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
func (lh *ResourceHandler) GetResourcesByLessonId(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "get_resources_by_lesson_id")

	lessonId, err := utils.GetUUIDParam(c, "lesson_id")
	if err != nil {
		logging.LogError("delete_module", "invalid lesson ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid lesson ID")
	}

	resources, err := lh.useCase.GetResourcesByLessonId(context.Background(), lessonId)
	if err != nil {
		return response.HandleApplicationError(c, err, "get_resources_by_lesson_id", lessonId.String())
	}

	logging.LogSuccess("get_resources_by_lesson_id", "Resources successfully retrieved", map[string]interface{}{
		"lesson_id": lessonId,
	})

	return response.OK(c, "Resources successfully retrieved", resources)
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
// @Router       /v1/api/resources [post]
func (lh *ResourceHandler) CreateResource(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "create_resource")

	var insertDTO dtos.ResourceInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		logging.LogError("create_resource", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.LogError("update_module", "invalid request data", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	newResource, err := lh.useCase.CreateResource(context.TODO(), insertDTO)
	if err != nil {
		return response.HandleApplicationError(c, err, "create_resource", newResource.ID.String())
	}

	logging.LogSuccess("create_resource", "Resource successfully created", map[string]interface{}{
		"resource_id": newResource.ID,
	})

	return response.Created(c, "Resource successfully created", newResource)
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
	logging.LogIncomingRequest(c, "update_resource")

	var insertDTO dtos.ResourceInsertDTO

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("update_resource", "invalid resouce ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid resouce ID")
	}

	if err := c.BodyParser(&insertDTO); err != nil {
		logging.LogError("update_resource", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.LogError("update_resource", "invalid request data", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	resourceUpdated, err := lh.useCase.UpdateResource(context.TODO(), id, insertDTO)
	if err != nil {
		return response.HandleApplicationError(c, err, "update_resource", resourceUpdated.ID.String())
	}

	logging.LogSuccess("update_resource", "Resource successfully updated", map[string]interface{}{
		"resource_id": resourceUpdated.ID,
	})

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
	logging.LogIncomingRequest(c, "delete_resource")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("delete_resource", "Invalid resource ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "Invalid resource ID")
	}

	err = lh.useCase.DeleteResource(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "delete_resource", id.String())
	}

	logging.LogSuccess("delete_resource", "Resource successfully updated", map[string]interface{}{
		"resource_id": id.String(),
	})

	return response.OK(c, "Resource successfully deleted", nil)
}
