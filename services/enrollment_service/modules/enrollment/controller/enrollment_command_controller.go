package controller

import (
	"context"

	c_services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/service"
	e_services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/service"
	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/service"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/response"
	"github.com/gofiber/fiber/v2"
)

// EnrollmentComandController handles enrollment-related write operations
type EnrollmentComandController struct {
	enrollmentService  e_services.EnrollmentService
	progressService    services.ProgressService
	certifcate_service c_services.CertificateService
}

// NewEnrollmentComandController creates a new controller instance
func NewEnrollmentComandController(
	enrollmentService e_services.EnrollmentService,
	certifcate_service c_services.CertificateService,
	progressService services.ProgressService,
) *EnrollmentComandController {
	return &EnrollmentComandController{
		enrollmentService:  enrollmentService,
		certifcate_service: certifcate_service,
		progressService:    progressService,
	}
}

// CompleteMyCourse godoc
// @Summary      Mark course as completed
// @Description  Complete a course and generate certificate for the enrollment
// @Tags         Enrollment Management
// @Accept       json
// @Produce      json
// @Param        enrollent_id   path      string  true  "Enrollment ID"
// @Success      200  {object}  response.ApiResponse{data=dtos.CertificateDTO} "Course completed successfully"
// @Failure      400  {object}  response.ApiResponse "Invalid enrollment ID"
// @Failure      404  {object}  response.ApiResponse "Enrollment not found"
// @Failure      500  {object}  response.ApiResponse "Internal server error"
// @Router       /v1/api/enrollments/{enrollent_id}/complete [post]
func (ec *EnrollmentComandController) CompleteMyCourse(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, KeyCompleteCourse)

	enrollentID, err := response.GetUUIDParam(c, "enrollent_id", KeyCompleteCourse)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyCompleteCourse, enrollentID.String())
	}

	if err := ec.enrollmentService.MarkEnrollmentComplete(context.Background(), enrollentID); err != nil {
		return response.HandleApplicationError(c, err, KeyCompleteCourse, enrollentID.String())
	}

	certificate, err := ec.certifcate_service.GenerateCertificate(context.Background(), enrollentID)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyCompleteCourse, enrollentID.String())
	}

	logging.LogSuccess(KeyCompleteCourse, MsgCourseCompleted, map[string]interface{}{
		"enrollent_id": enrollentID,
	})

	return response.OK(c, MsgCourseCompleted, certificate)
}

// CancellMyEnrollment godoc
// @Summary      Cancel enrollment
// @Description  Cancel an existing enrollment for a user
// @Tags         Enrollment Management
// @Accept       json
// @Produce      json
// @Param        user_id        path      string  true  "User ID"
// @Param        enrollent_id   path      string  true  "Enrollment ID"
// @Success      200  {object}  response.ApiResponse "Enrollment cancelled successfully"
// @Failure      400  {object}  response.ApiResponse "Invalid user/enrollment ID"
// @Failure      404  {object}  response.ApiResponse "Enrollment not found"
// @Failure      500  {object}  response.ApiResponse "Internal server error"
// @Router       /v1/api/users/{user_id}/enrollments/{enrollent_id} [delete]
func (ec *EnrollmentComandController) CancellMyEnrollment(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, KeyCancelEnrollment)

	enrollentID, err := response.GetUUIDParam(c, "enrollent_id", KeyCancelEnrollment)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidEnrollmentID)
	}

	userID, err := response.GetUUIDParam(c, "user_id", KeyCancelEnrollment)
	if err != nil {
		logging.LogError(KeyCancelEnrollment, "Invalid user ID", map[string]interface{}{
			"error": err.Error(),
		})
		return response.BadRequest(c, err.Error(), MsgInvalidUserID)
	}

	if err := ec.enrollmentService.CancelEnrollment(context.Background(), userID, enrollentID); err != nil {
		return response.HandleApplicationError(c, err, KeyCancelEnrollment, enrollentID.String())
	}

	logging.LogSuccess(KeyCancelEnrollment, MsgEnrollmentCancelled, map[string]interface{}{
		"enrollent_id": enrollentID,
	})

	return response.OK(c, MsgEnrollmentCancelled, nil)
}

// EnrollUserInCourse godoc
// @Summary      Enroll user in course
// @Description  Create a new enrollment and track record for a user-course combination
// @Tags         Enrollment Management
// @Accept       json
// @Produce      json
// @Param        user_id    path      string  true  "User ID"
// @Param        course_id  path      string  true  "Course ID"
// @Success      201  {object}  response.ApiResponse "Enrollment created successfully"
// @Failure      400  {object}  response.ApiResponse "Invalid user/course ID"
// @Failure      409  {object}  response.ApiResponse "User already enrolled"
// @Failure      500  {object}  response.ApiResponse "Internal server error"
// @Router       /v1/api/users/{user_id}/courses/{course_id}/enroll [post]
func (ec *EnrollmentComandController) EnrollUserInCourse(c *fiber.Ctx) error {
	logging.LogIncomingRequest(c, KeyEnrollUserInCourse)

	userID, err := response.GetUUIDParam(c, "user_id", KeyEnrollUserInCourse)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidUserID)
	}

	courseID, err := response.GetUUIDParam(c, "course_id", KeyEnrollUserInCourse)
	if err != nil {
		return response.BadRequest(c, err.Error(), MsgInvalidCourseID)
	}

	enrollment, err := ec.enrollmentService.EnrollUserInCourse(context.Background(), userID, courseID)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyEnrollUserInCourse, userID.String())
	}

	err = ec.progressService.CreateCourseTrackRecord(context.TODO(), enrollment.ID)
	if err != nil {
		return response.HandleApplicationError(c, err, KeyEnrollUserInCourse, userID.String())
	}

	logging.LogSuccess(KeyEnrollUserInCourse, MsgUserEnrolled, map[string]interface{}{
		"user_id": userID,
	})

	return response.Created(c, MsgUserEnrolled, nil)
}
