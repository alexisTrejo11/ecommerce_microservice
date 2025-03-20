package controller

import (
	"context"

	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/service"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/response"
	"github.com/gofiber/fiber/v2"
)

type EnrollmentComandController struct {
	entollmentService services.EnrollmentService
}

func NewEnrollmentComandController(entollmentService services.EnrollmentService) *EnrollmentComandController {
	return &EnrollmentComandController{
		entollmentService: entollmentService,
	}
}

func (ec *EnrollmentComandController) EnrollUserInCourse(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "enrroll_user_in_course")

	userID, err := response.GetUUIDParam(c, "user_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	courseID, err := response.GetUUIDParam(c, "course_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_course_id")
	}

	if _, err := ec.entollmentService.EnrollUserInCourse(context.Background(), userID, courseID); err != nil {
		return response.HandleApplicationError(c, err, "enrroll_user_in_course", userID.String())
	}

	logging.LogSuccess("enrroll_user_in_course", "User Successfully Enrolled", map[string]interface{}{
		"use_id": userID,
	})

	return response.Created(c, "User Successfully Enrolled", nil)
}

// Validate Course is completed
func (ec *EnrollmentComandController) CompleteCourse(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "complete_course")

	enrollentID, err := response.GetUUIDParam(c, "enrollent_id")
	if err != nil {
		return response.HandleApplicationError(c, err, "enrroll_user_in_course", enrollentID.String())
	}

	// Auth
	if err := ec.entollmentService.MarkEnrollmentComplete(context.Background(), enrollentID); err != nil {
		return response.HandleApplicationError(c, err, "complete_course", enrollentID.String())
	}

	// Generate Certifcate

	logging.LogSuccess("complete_course", "Course Succesfully Completed", map[string]interface{}{
		"enrollent_id": enrollentID,
	})

	return response.OK(c, "Course Succesfully Completed", nil)
}

// Valdiate Cancealibiltys
func (ec *EnrollmentComandController) CancellEnrollment(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, "cancell_enrollment")

	enrollentID, err := response.GetUUIDParam(c, "enrollent_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	// Auth
	if err := ec.entollmentService.CancelEnrollment(context.Background(), enrollentID); err != nil {
		return response.HandleApplicationError(c, err, "cancell_enrollment", enrollentID.String())
	}

	logging.LogSuccess("cancell_enrollment", "Enrollment Succesfully Cancelled", map[string]interface{}{
		"enrollent_id": enrollentID,
	})

	return response.OK(c, "Enrollment Succesfully Cancelled", nil)
}
