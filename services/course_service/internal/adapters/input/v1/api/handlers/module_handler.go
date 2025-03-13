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
	logging.LogIncomingRequest(c, "get_module_by_id")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("get_module_by_id", "Invalid module ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	module, err := lh.useCase.GetModuleById(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "get_module_by_id", id.String())
	}

	logging.LogSuccess("get_module_by_id", "Module successfully retrieved", map[string]interface{}{
		"module_id": id,
	})

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
func (lh *ModuleHandler) GetModulesByCourseId(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "get_modules_course_by_id")

	id, err := utils.GetUUIDParam(c, "course_id")
	if err != nil {
		logging.LogError("update_course", "Invalid course ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "Invalid course ID")
	}

	modules, err := lh.useCase.GetModuleByCourseId(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "get_modules_course_by_id", id.String())
	}

	logging.LogSuccess("get_modules_course_by_id", "Modules successfully retrieved", map[string]interface{}{
		"course_id": id,
	})

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
	logging.LogIncomingRequest(c, "create_module")

	var insertDTO dtos.ModuleInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		logging.LogError("create_module", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.LogError("create_module", "invalid request data", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	moduleCreated, err := lh.useCase.CreateModule(context.TODO(), insertDTO)
	if err != nil {
		return response.HandleApplicationError(c, err, "create_module", moduleCreated.ID.String())
	}

	logging.LogSuccess("create_module", "Module successfully created", map[string]interface{}{
		"module_id": moduleCreated.ID,
	})

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
	logging.LogIncomingRequest(c, "update_module")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("update_module", "Invalid module ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	var insertDTO dtos.ModuleInsertDTO

	if err := c.BodyParser(&insertDTO); err != nil {
		logging.LogError("update_module", "can't parse body request", map[string]interface{}{
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

	moduleUpdated, err := lh.useCase.UpdateModule(context.TODO(), id, insertDTO)
	if err != nil {
		return response.HandleApplicationError(c, err, "update_module", id.String())
	}

	logging.LogSuccess("update_module", "Module successfully updated", map[string]interface{}{
		"module_id": id,
	})

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
	logging.LogIncomingRequest(c, "delete_module")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("delete_module", "Invalid module ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "Invalid module ID")
	}

	err = lh.useCase.DeleteModule(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "delete_module", id.String())
	}

	logging.LogSuccess("delete_module", "Module successfully deleted", map[string]interface{}{
		"course_id": id,
	})

	return response.OK(c, "Module successfully deleted", nil)
}
