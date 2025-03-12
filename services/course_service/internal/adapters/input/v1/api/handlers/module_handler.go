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

// GetModuleById godoc
// @Summary      Get Module by ID
// @Description  Retrieve a module by its unique ID.
// @Tags         Modules
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Module ID"
// @Success      200  {object}  response.ApiResponse{data=dtos.ModuleDTO} "Module successfully retrieved"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Failure      404  {object}  response.ApiResponse "Module not found"
// @Router       /v1/api/modules/{id} [get]
func (lh *ModuleHandler) GetModuleById(c *fiber.Ctx) error {
	// Log incoming request
	logging.Logger.WithFields(logrus.Fields{
		"action": "get_module_by_id",
		"method": c.Method(),
		"route":  c.Route().Path,
		"ip":     c.IP(),
		"param":  c.Params("id"),
		"user":   c.Locals("user_id"),
	}).Info("Incoming request")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "get_module_by_id",
			"error":  err.Error(),
		}).Error("Invalid Module ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	module, err := lh.useCase.GetModuleById(context.Background(), id)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":    "get_module_by_id",
			"module_id": id,
			"error":     err.Error(),
		}).Error("Module not found")
		return response.NotFound(c, "Module not found", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":    "get_module_by_id",
		"module_id": id,
	}).Info("Module successfully retrieved")

	return response.OK(c, "Module successfully retrieved", module)
}

// GetModuleByCourseId godoc
// @Summary      Get Modules by Course ID
// @Description  Retrieve modules associated with a given course ID.
// @Tags         Modules
// @Accept       json
// @Produce      json
// @Param        course_id   path      string  true  "Course ID"
// @Success      200  {object}  response.ApiResponse{data=[]dtos.ModuleDTO} "Modules successfully retrieved"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Failure      404  {object}  response.ApiResponse "Modules not found"
// @Router       /v1/api/modules/course/{course_id} [get]
func (lh *ModuleHandler) GetModuleByCourseId(c *fiber.Ctx) error {
	// Log incoming request
	logging.Logger.WithFields(logrus.Fields{
		"action": "get_module_by_course_id",
		"method": c.Method(),
		"route":  c.Route().Path,
		"ip":     c.IP(),
		"param":  c.Params("course_id"),
		"user":   c.Locals("user_id"),
	}).Info("Incoming request")

	id, err := utils.GetUUIDParam(c, "course_id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "get_module_by_course_id",
			"error":  err.Error(),
		}).Error("Invalid Course ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	modules, err := lh.useCase.GetModuleByCourseId(context.Background(), id)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":    "get_module_by_course_id",
			"course_id": id,
			"error":     err.Error(),
		}).Error("Modules not found")
		return response.NotFound(c, "Modules not found", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":    "get_module_by_course_id",
		"course_id": id,
	}).Info("Modules successfully retrieved")

	return response.OK(c, "Modules successfully retrieved", modules)
}

// CreateModule godoc
// @Summary      Create Module
// @Description  Create a new module with the provided details.
// @Tags         Modules
// @Accept       json
// @Produce      json
// @Param        module  body      dtos.ModuleInsertDTO  true  "Module data"
// @Success      201  {object}  response.ApiResponse{data=dtos.ModuleDTO} "Module successfully created"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Router       /v1/api/modules [post]
func (lh *ModuleHandler) CreateModule(c *fiber.Ctx) error {
	// Log incoming request with payload
	logging.Logger.WithFields(logrus.Fields{
		"action":  "create_module",
		"method":  c.Method(),
		"route":   c.Route().Path,
		"ip":      c.IP(),
		"user":    c.Locals("user_id"),
		"payload": c.Body(),
	}).Info("Incoming request")

	var insertDTO dtos.ModuleInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "create_module",
			"error":  err.Error(),
		}).Error("Invalid request body")
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "create_module",
			"error":  err.Error(),
		}).Error("Validation failed")
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	moduleCreated, err := lh.useCase.CreateModule(context.TODO(), insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "create_module",
			"error":  err.Error(),
		}).Error("Error creating module")
		return response.BadRequest(c, "Error creating module", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":    "create_module",
		"module_id": moduleCreated.ID,
	}).Info("Module successfully created")

	return response.Created(c, "Module successfully created", moduleCreated)
}

// UpdateModule godoc
// @Summary      Update Module
// @Description  Update an existing module by its ID with the provided details.
// @Tags         Modules
// @Accept       json
// @Produce      json
// @Param        id      path      string                true  "Module ID"
// @Param        module  body      dtos.ModuleInsertDTO  true  "Module data to update"
// @Success      200  {object}  response.ApiResponse{data=dtos.ModuleDTO} "Module successfully updated"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Router       /v1/api/modules/{id} [put]
func (lh *ModuleHandler) UpdateModule(c *fiber.Ctx) error {
	// Log incoming request with payload
	logging.Logger.WithFields(logrus.Fields{
		"action":  "update_module",
		"method":  c.Method(),
		"route":   c.Route().Path,
		"ip":      c.IP(),
		"user":    c.Locals("user_id"),
		"payload": c.Body(),
	}).Info("Incoming request")

	var insertDTO dtos.ModuleInsertDTO

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "update_module",
			"error":  err.Error(),
		}).Error("Invalid module ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	if err := c.BodyParser(&insertDTO); err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "update_module",
			"error":  err.Error(),
		}).Error("Invalid request body")
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "update_module",
			"error":  err.Error(),
		}).Error("Validation failed")
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	moduleUpdated, err := lh.useCase.UpdateModule(context.TODO(), id, insertDTO)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":    "update_module",
			"module_id": id,
			"error":     err.Error(),
		}).Error("Error updating module")
		return response.BadRequest(c, "Error updating module", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":    "update_module",
		"module_id": id,
	}).Info("Module successfully updated")

	return response.OK(c, "Module successfully updated", moduleUpdated)
}

// DeleteModule godoc
// @Summary      Delete Module
// @Description  Delete an existing module by its ID.
// @Tags         Modules
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Module ID"
// @Success      200  {object}  response.ApiResponse "Module successfully deleted"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Failure      404  {object}  response.ApiResponse "Module not found"
// @Router       /v1/api/modules/{id} [delete]
func (lh *ModuleHandler) DeleteModule(c *fiber.Ctx) error {
	// Log incoming request for deletion
	logging.Logger.WithFields(logrus.Fields{
		"action": "delete_module",
		"method": c.Method(),
		"route":  c.Route().Path,
		"ip":     c.IP(),
		"user":   c.Locals("user_id"),
		"param":  c.Params("id"),
	}).Info("Incoming request")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action": "delete_module",
			"error":  err.Error(),
		}).Error("Invalid module ID")
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	err = lh.useCase.DeleteModule(context.Background(), id)
	if err != nil {
		logging.Logger.WithFields(logrus.Fields{
			"action":    "delete_module",
			"module_id": id,
			"error":     err.Error(),
		}).Error("Module not found")
		return response.NotFound(c, "Module not found", err.Error())
	}

	logging.Logger.WithFields(logrus.Fields{
		"action":    "delete_module",
		"module_id": id,
	}).Info("Module successfully deleted")

	return response.OK(c, "Module successfully deleted", nil)
}
