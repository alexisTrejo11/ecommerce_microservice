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
		return response.Unauthorized(c, err.Error(), MsgUnauthorized)
	}

	logging.LogIncomingRequest(c, KeyGetMyEnrollments)

	enrollment, _, err := ec.entollmentService.GetUserEnrollments(context.Background(), userID, 1, 10)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyGetEnrollmentByUserAndCourse, userID.String())
	}

	logging.LogSuccess(KeyGetMyEnrollments, MsgUserEnrollmentRetrieved, map[string]interface{}{
		"user_id": userID,
	})

	return response.OK(c, MsgUserEnrollmentRetrieved, enrollment)
}

func (ec *EnrollmentQueryController) GetEnrollmentByID(c *fiber.Ctx) error {
	enrollmentID, err := response.GetUUIDParam(c, "enrollment_id", KeyGetEnrollmentByID)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidEnrollmentID)
	}

	logging.LogIncomingRequest(c, KeyGetEnrollmentByID)

	enrollment, err := ec.entollmentService.GetEnrollmentByID(context.Background(), enrollmentID)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyGetEnrollmentByUserAndCourse, enrollmentID.String())
	}

	logging.LogSuccess(KeyGetEnrollmentByID, MsgEnrollmentRetrieved, map[string]interface{}{
		"enrollment_id": enrollmentID,
	})

	return response.OK(c, MsgEnrollmentRetrieved, enrollment)
}

func (ec *EnrollmentQueryController) GetEnrollmentByUserAndCourse(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, KeyGetEnrollmentByUserAndCourse)

	userID, err := response.GetUUIDParam(c, "user_id", KeyGetEnrollmentByUserAndCourse)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidUserID)
	}

	courseID, err := response.GetUUIDParam(c, "course_id", KeyGetEnrollmentByUserAndCourse)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidCourseID)
	}

	enrollment, err := ec.entollmentService.GetEnrollmentByUserAndCourse(context.Background(), userID, courseID)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyGetEnrollmentByUserAndCourse, courseID.String())
	}

	logging.LogSuccess(KeyGetEnrollmentByUserAndCourse, MsgEnrollmentsRetrieved, map[string]interface{}{
		"course_id": courseID,
		"user_id":   userID,
	})

	return response.OK(c, MsgEnrollmentsRetrieved, enrollment)
}

func (ec *EnrollmentQueryController) GetCourseEnrollments(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, KeyGetCourseEnrollments)

	courseID, err := response.GetUUIDParam(c, "course_id", KeyGetCourseEnrollments)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidCourseID)
	}

	enrollments, _, err := ec.entollmentService.GetCourseEnrollments(context.Background(), courseID, 1, 10)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyGetEnrollmentByUserAndCourse, courseID.String())
	}

	if len(enrollments) == 0 {
		return response.OK(c, MsgNoEnrollmentsFoundForCourse, enrollments)
	}

	logging.LogSuccess(KeyGetCourseEnrollments, MsgCourseEnrollmentRetrieved, map[string]interface{}{
		"course_id": courseID,
	})

	return response.OK(c, MsgCourseEnrollmentRetrieved, enrollments)
}
