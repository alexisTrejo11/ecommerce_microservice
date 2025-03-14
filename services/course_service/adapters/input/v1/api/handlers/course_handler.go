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

// GetCourseById godoc
// @Summary      Get Course by ID
// @Description  Retrieve a course by its unique ID.
// @Tags         Courses
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Course ID"
// @Success      200  {object}  response.ApiResponse{data=dtos.CourseDTO} "Course successfully retrieved"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Failure      404  {object}  response.ApiResponse "Course not found"
// @Router       /v1/api/courses/{id} [get]
func (lh *CourseHandler) GetCourseById(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "get_course_by_id")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("get_course_by_id", "invalid course ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid course ID")
	}

	course, err := lh.useCase.GetCourseById(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "get_course_by_id", id.String())
	}

	logging.LogSuccess("get_course_by_id", "Course successfully retrieved", map[string]interface{}{
		"course_id": course.ID,
	})

	return response.OK(c, "Course Successfully Retrieved", course)
}

// CreateCourse godoc
// @Summary      Create a new Course
// @Description  Create a course with the provided information.
// @Tags         Courses
// @Accept       json
// @Produce      json
// @Param        course  body      dtos.CourseInsertDTO  true  "Course information"
// @Success      201     {object}  response.ApiResponse{data=dtos.CourseDTO} "Course successfully created"
// @Failure      400     {object}  response.ApiResponse "Bad Request"
// @Router       /v1/api/courses [post]
func (lh *CourseHandler) CreateCourse(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "create_course")

	var insertDTO dtos.CourseInsertDTO
	if err := c.BodyParser(&insertDTO); err != nil {
		logging.LogError("create_course", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	courseCreated, err := lh.useCase.CreateCourse(context.TODO(), insertDTO)
	if err != nil {
		return response.HandleApplicationError(c, err, "create_course", courseCreated.ID)
	}

	logging.LogSuccess("create_course", "Course successfully created", map[string]interface{}{
		"course_id": courseCreated.ID,
	})

	return response.Created(c, "Course successfully created", courseCreated)
}

// UpdateCourse godoc
// @Summary      Update an existing Course
// @Description  Update course details using its ID.
// @Tags         Courses
// @Accept       json
// @Produce      json
// @Param        id      path      string                true  "Course ID"
// @Param        course  body      dtos.CourseInsertDTO  true  "Course information to update"
// @Success      200     {object}  response.ApiResponse{data=dtos.CourseDTO} "Course successfully updated"
// @Failure      400     {object}  response.ApiResponse "Bad Request"
// @Router       /v1/api/courses/{id} [put]
func (lh *CourseHandler) UpdateCourse(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "update_course")

	var insertDTO dtos.CourseInsertDTO
	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("update_course", "Invalid course ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	if err := c.BodyParser(&insertDTO); err != nil {
		logging.LogError("update_course", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	errorsMap, err := utils.ValidateStruct(lh.validator, &insertDTO)
	if err != nil {
		logging.LogError("update_course", "can't parse body request", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, "Validation failed", errorsMap)
	}

	CourseUpdated, err := lh.useCase.UpdateCourse(context.TODO(), id, insertDTO)
	if err != nil {
		return response.HandleApplicationError(c, err, "update_course", id.String())
	}

	logging.LogSuccess("update_course", "Course successfully deleted", map[string]interface{}{
		"course_id": id,
	})

	return response.OK(c, "Course successfully updated", CourseUpdated)
}

// DeleteCourse godoc
// @Summary      Delete a Course
// @Description  Delete an existing course by its unique ID.
// @Tags         Courses
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Course ID"
// @Success      200  {object}  response.ApiResponse "Course successfully deleted"
// @Failure      400  {object}  response.ApiResponse "Bad Request"
// @Failure      404  {object}  response.ApiResponse "Course not found"
// @Router       /v1/api/courses/{id} [delete]
func (lh *CourseHandler) DeleteCourse(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "delete_course")

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		logging.LogError("delete_course", "Invalid course ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), "invalid id")
	}

	err = lh.useCase.DeleteCourse(context.Background(), id)
	if err != nil {
		return response.HandleApplicationError(c, err, "delete_course", id.String())
	}

	logging.LogSuccess("delete_course", "Course successfully deleted", map[string]interface{}{
		"course_id": id,
	})

	return response.OK(c, "Course successfully deleted", nil)
}
