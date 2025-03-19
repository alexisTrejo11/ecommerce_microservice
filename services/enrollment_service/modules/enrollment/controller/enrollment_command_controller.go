package controller

import (
	"context"

	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/service"
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
	userID, err := response.GetUUIDParam(c, "user_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	courseID, err := response.GetUUIDParam(c, "course_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_course_id")
	}

	if _, err := ec.entollmentService.EnrollUserInCourse(context.Background(), userID, courseID); err != nil {
		return response.BadRequest(c, err.Error(), "error")
	}

	return response.Created(c, "User Successfully Enrolled", nil)
}

func (ec *EnrollmentComandController) CompleteCourse(c *fiber.Ctx) error {
	enrollentID, err := response.GetUUIDParam(c, "enrollent_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	// Auth
	if err := ec.entollmentService.MarkEnrollmentComplete(context.Background(), enrollentID); err != nil {
		return response.BadRequest(c, err.Error(), "error")
	}

	return response.OK(c, "User Successfully Enrolled", nil)
}

// Soft Cancel // Valdiate Cancealibiltys
func (ec *EnrollmentComandController) CancellEnrollment(c *fiber.Ctx) error {
	enrollentID, err := response.GetUUIDParam(c, "enrollent_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	// Auth
	if err := ec.entollmentService.CancelEnrollment(context.Background(), enrollentID); err != nil {
		return response.BadRequest(c, err.Error(), "error")
	}

	return response.OK(c, "Enrollment Succesfully Cancelled", nil)
}
