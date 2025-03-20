package controller

import (
	"context"

	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/middleware"
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

func (ec *EnrollmentQueryController) GetMyEnrollments(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return response.Unauthorized(c, err.Error(), "unauthorized")
	}

	logging.LogIncomingRequest(c, "get_my_enrollments")

	enrollment, _, err := ec.entollmentService.GetUserEnrollments(context.Background(), userID, 1, 10)
	if err != nil {
		return response.NotFound(c, "Enrollment Not Found", "enrollment_not_found")
	}

	logging.LogSuccess("get_my_enrollments", "User Enrollment Successfully Retrieved", map[string]interface{}{
		"user_id": userID,
	})

	return response.OK(c, "User Enrollment Successfully Retrieved", enrollment)
}

func (ec *EnrollmentQueryController) GetEnrollmentByID(c *fiber.Ctx) error {
	enrollmentID, err := response.GetUUIDParam(c, "enrollment_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_enrollment_id")
	}

	logging.LogIncomingRequest(c, "get_enrollments_by_id")

	enrollment, err := ec.entollmentService.GetEnrollmentByID(context.Background(), enrollmentID)
	if err != nil {
		return response.NotFound(c, "Enrollment Not Found", "enrollment_not_found")
	}

	logging.LogSuccess("get_enrollments_by_id", "Enrollment Successfully Retrieved", map[string]interface{}{
		"enrollment_id": enrollmentID,
	})

	return response.OK(c, "Enrollment Successfully Retrieved", enrollment)
}

func (ec *EnrollmentQueryController) GetEnrollmentByUserAndCourse(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "get_enrollment_by_user_and_course")

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

	logging.LogSuccess("get_enrollment_by_user_and_course", "Enrollments Successfully Retrieved", map[string]interface{}{
		"course_id": courseID,
		"user_id":   userID,
	})

	return response.OK(c, "Enrollments Successfully Retrieved", enrollment)
}

func (ec *EnrollmentQueryController) GetCourseEnrollments(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "get_course_enrollments")

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

	logging.LogSuccess("get_enrollments_by_id", "Enrollment Successfully Retrieved", map[string]interface{}{
		"course_id": courseID,
	})

	return response.OK(c, "Course Enrollment Successfully Retrieved", enrollments)
}
