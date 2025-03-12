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

	return response.OK(c, "Course successfully deleted", nil)
}
