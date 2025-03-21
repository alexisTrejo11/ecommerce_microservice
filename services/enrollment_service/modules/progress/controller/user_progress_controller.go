package controller

import (
	"context"

	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
	logging "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/logger"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/middleware"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/response"
	"github.com/gofiber/fiber/v2"
)

// ProgressController handles operations related to course progress.
type ProgressController struct {
	service    services.ProgressService
	jwtManager jwt.JWTManager
}

// NewProgressController creates a new instance of the ProgressController.
func NewProgressController(service services.ProgressService, jwtManager jwt.JWTManager) *ProgressController {
	return &ProgressController{
		service:    service,
		jwtManager: jwtManager,
	}
}

// @Summary Get My Course Progress
// @Description Retrieves the progress of all lessons for the authenticated user.
// @Tags Progress
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object}  response.ApiResponse "Course progress retrieved successfully"
// @Failure 401 {object}  response.ApiResponse "Unauthorized"
// @Router /progress/me [get]
func (pc *ProgressController) GetMyCourseProgress(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return response.Unauthorized(c, err.Error(), errUnauthorized)
	}

	logging.LogIncomingRequest(c, opGetCourseProgress)

	lessons, err := pc.service.GetCourseProgress(context.Background(), userID)
	if err != nil {
		return response.HandleApplicationError(c, err, opGetCourseProgress, userID.String())
	}

	return response.OK(c, msgCourseProgressRetrieved, lessons)
}

// @Summary Mark Lesson as Complete
// @Description Marks a specific lesson as complete for the authenticated user.
// @Tags Progress
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param lesson_id path string true "Lesson ID"
// @Success 200 {object}  response.ApiResponse "Lesson marked as complete successfully"
// @Failure 400 {object}  response.ApiResponse "Invalid lesson ID"
// @Failure 401 {object}  response.ApiResponse "Unauthorized"
// @Failure 500 {object}  response.ApiResponse "Internal server error"
// @Router /progress/lesson/{lesson_id}/complete [post]
func (pc *ProgressController) MarkMyLessonComplete(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return response.Unauthorized(c, err.Error(), errUnauthorized)
	}

	logging.LogIncomingRequest(c, opMarkLessonComplete)

	lessonID, err := response.GetUUIDParam(c, "lesson_id", opMarkLessonComplete)
	if err != nil {
		return response.BadRequest(c, err.Error(), errInvalidLessonID)
	}

	err = pc.service.MarkLessonComplete(context.Background(), userID, lessonID)
	if err != nil {
		return response.HandleApplicationError(c, err, opMarkLessonComplete, lessonID.String())
	}

	logging.LogSuccess(opMarkLessonComplete, msgLessonCompleted, map[string]interface{}{
		"enrollment_id": userID,
	})

	return response.OK(c, msgLessonCompleted, nil)
}

// @Summary Mark Lesson as Incomplete
// @Description Marks a specific lesson as incomplete for the authenticated user.
// @Tags Progress
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param lesson_id path string true "Lesson ID"
// @Success 200 {object}  response.ApiResponse "Lesson marked as incomplete successfully"
// @Failure 400 {object}  response.ApiResponse "Invalid lesson ID"
// @Failure 401 {object}  response.ApiResponse "Unauthorized"
// @Failure 500 {object}  response.ApiResponse "Internal server error"
// @Router /progress/lesson/{lesson_id}/incomplete [post]
func (pc *ProgressController) MarkMyLessonIncomplete(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return response.Unauthorized(c, err.Error(), errUnauthorized)
	}

	logging.LogIncomingRequest(c, opMarkLessonIncomplete)

	lessonID, err := response.GetUUIDParam(c, "lesson_id", opMarkLessonIncomplete)
	if err != nil {
		return response.BadRequest(c, err.Error(), errInvalidLessonID)
	}

	err = pc.service.MarkLessonIncomplete(context.Background(), userID, lessonID)
	if err != nil {
		return response.HandleApplicationError(c, err, opMarkLessonIncomplete, lessonID.String())
	}

	logging.LogSuccess(opMarkLessonIncomplete, msgLessonIncompleted, map[string]interface{}{
		"enrollment_id": userID,
	})

	return response.OK(c, msgLessonIncompleted, nil)
}
