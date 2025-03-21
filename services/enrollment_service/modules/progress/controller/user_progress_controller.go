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

type ProgressController struct {
	service    services.ProgressService
	jwtManager jwt.JWTManager
}

func NewProgressController(service services.ProgressService, jwtManager jwt.JWTManager) *ProgressController {
	return &ProgressController{
		service:    service,
		jwtManager: jwtManager,
	}
}

// GetMyCourseProgress retrieves course progress for the authenticated user
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

// MarkMyLessonComplete marks a lesson as complete for the authenticated user
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

// MarkMyLessonIncomplete marks a lesson as incomplete for the authenticated user
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
