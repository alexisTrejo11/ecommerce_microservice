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
	"github.com/sirupsen/logrus"
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
	// Log incoming request
	logging.Logger.WithFields(logrus.Fields{
		"action": "get_resource_by_id",
		"method": c.Method(),
		"route":  c.Route().Path,
		"ip":     c.IP(),
		"param":  c.Params("id"),
		"user":   c.Locals("user_id"),
	}).Info("Incoming request")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "get_resource_by_id",
			"error":  err.Error(),
		}).Error("Invalid resource ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	resource, err := lh.useCase.GetResourceById(context.Background(), id)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":      "get_resource_by_id",
			"resource_id": id,
			"error":       err.Error(),
		}).Error("Resource not found")
		return response.NotFound(c, "Resource not found", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":      "get_resource_by_id",
		"resource_id": id,
	}).Info("Resource successfully retrieved")

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
	// Log incoming request
	logging.Logger.WithFields(logrus.Fields{
		"action": "get_resource_by_lesson_id",
		"method": c.Method(),
		"route":  c.Route().Path,
		"ip":     c.IP(),
		"param":  c.Params("lesson_id"),
		"user":   c.Locals("user_id"),
	}).Info("Incoming request")

	lessonId, err := utils.GetUUIDParam(c, "lesson_id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "get_resource_by_lesson_id",
			"error":  err.Error(),
		}).Error("Invalid lesson ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	resource, err := lh.useCase.GetResourcesByLessonId(context.Background(), lessonId)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":    "get_resource_by_lesson_id",
			"lesson_id": lessonId,
			"error":     err.Error(),
		}).Error("Resource not found")
		return response.NotFound(c, "Resource not found", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":    "get_resource_by_lesson_id",
		"lesson_id": lessonId,
	}).Info("Resource successfully retrieved")

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
// @Router       /v1/api/resources [post]
func (lh *ResourceHandler) CreateResource(c *fiber.Ctx) error {
	// Log incoming request with payload
	logging.Logger.WithFields(logrus.Fields{
		"action":  "create_resource",
		"method":  c.Method(),
		"route":   c.Route().Path,
		"ip":      c.IP(),
		"user":    c.Locals("user_id"),
		"payload": c.Body(),
	}).Info("Incoming request")

	var insertDTO dtos.ResourceInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "create_resource",
			"error":  err.Error(),
		}).Error("Invalid request body")
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "create_resource",
			"error":  err.Error(),
		}).Error("Validation failed")
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	resourceCreated, err := lh.useCase.CreateResource(context.TODO(), insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "create_resource",
			"error":  err.Error(),
		}).Error("Error creating resource")
		return response.BadRequest(c, "Error creating resource", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":      "create_resource",
		"resource_id": resourceCreated.ID,
	}).Info("Resource successfully created")

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
	// Log incoming request with payload
	logging.Logger.WithFields(logrus.Fields{
		"action":  "update_resource",
		"method":  c.Method(),
		"route":   c.Route().Path,
		"ip":      c.IP(),
		"user":    c.Locals("user_id"),
		"payload": c.Body(),
	}).Info("Incoming request")

	var insertDTO dtos.ResourceInsertDTO

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "update_resource",
			"error":  err.Error(),
		}).Error("Invalid resource ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	if err := c.BodyParser(&insertDTO); err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "update_resource",
			"error":  err.Error(),
		}).Error("Invalid request body")
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "update_resource",
			"error":  err.Error(),
		}).Error("Validation failed")
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	resourceUpdated, err := lh.useCase.UpdateResource(context.TODO(), id, insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":      "update_resource",
			"resource_id": id,
			"error":       err.Error(),
		}).Error("Error updating resource")
		return response.BadRequest(c, "Error updating resource", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":      "update_resource",
		"resource_id": id,
	}).Info("Resource successfully updated")

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
	// Log incoming request for deletion
	logging.Logger.WithFields(logrus.Fields{
		"action": "delete_resource",
		"method": c.Method(),
		"route":  c.Route().Path,
		"ip":     c.IP(),
		"user":   c.Locals("user_id"),
		"param":  c.Params("id"),
	}).Info("Incoming request")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "delete_resource",
			"error":  err.Error(),
		}).Error("Invalid resource ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	err = lh.useCase.DeleteResource(context.Background(), id)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":      "delete_resource",
			"resource_id": id,
			"error":       err.Error(),
		}).Error("Resource not found")
		return response.NotFound(c, "Resource not found", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":      "delete_resource",
		"resource_id": id,
	}).Info("Resource successfully deleted")

	return response.OK(c, "Resource successfully deleted", nil)
}
