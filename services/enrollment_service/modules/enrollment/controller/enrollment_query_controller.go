package controller

import (
	"context"

	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/response"
	"github.com/gofiber/fiber/v2"
)

type EnrollmentQueryController struct {
	entollmentService services.EnrollmentService
	jwtManager        jwt.JWTManager
}

func NewEnrollmentQueryController(entollmentService services.EnrollmentService, jwtManager jwt.JWTManager) *EnrollmentQueryController {
	return &EnrollmentQueryController{
		entollmentService: entollmentService,
		jwtManager:        jwtManager,
	}
}

func (ec *EnrollmentQueryController) GetUserEnrollments(c *fiber.Ctx) error {
	userID, err := ec.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	enrollment, _, err := ec.entollmentService.GetUserEnrollments(context.Background(), userID, 1, 10)
	if err != nil {
		return response.NotFound(c, "Enrollment Not Found", "enrollment_not_found")
	}

	return response.OK(c, "Enrollment Successfully Retrieved", enrollment)
}

func (ec *EnrollmentQueryController) GetEnrollmentByID(c *fiber.Ctx) error {
	enrollmentID, err := response.GetUUIDParam(c, "enrollment_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_enrollment_id")
	}

	enrollment, err := ec.entollmentService.GetEnrollmentByID(context.Background(), enrollmentID)
	if err != nil {
		return response.NotFound(c, "Enrollment Not Found", "enrollment_not_found")
	}

	return response.OK(c, "Enrollment Successfully Retrieved", enrollment)
}

func (ec *EnrollmentQueryController) GetEnrollmentByUserAndCourse(c *fiber.Ctx) error {
	userID, err := response.GetUUIDParam(c, "user_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	courseID, err := response.GetUUIDParam(c, "course_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_course_id")
	}

	enrollment, err := ec.entollmentService.GetEnrollmentByUserAndCourse(context.Background(), userID, courseID)
	if err != nil {
		return response.NotFound(c, "Enrollment Not Found", "enrollment_not_found")
	}

	return response.OK(c, "Enrollment Successfully Retrieved", enrollment)
}

func (ec *EnrollmentQueryController) GetCourseEnrollments(c *fiber.Ctx) error {
	courseID, err := response.GetUUIDParam(c, "course_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_course_id")
	}

	enrollments, _, err := ec.entollmentService.GetCourseEnrollments(context.Background(), courseID, 1, 10)
	if err != nil {
		return response.NotFound(c, "Enrollment Not Found", "enrollment_not_found")
	}

	if len(enrollments) == 0 {
		return response.OK(c, "No Enrollments Found for this Course", enrollments)
	}

	return response.OK(c, "Course Enrollment Successfully Retrieved", enrollments)
}
