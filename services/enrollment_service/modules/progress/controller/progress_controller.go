package controller

import (
	"context"

	services "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/service"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/jwt"
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

// Group By Course
func (pc *ProgressController) GetMyCourseProgress(c *fiber.Ctx) error {
	userID, err := pc.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	lessons, err := pc.service.GetCourseProgress(context.Background(), userID)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid")
	}

	return response.OK(c, "User Course Progress Succesfully Retrieved", lessons)
}

func (pc *ProgressController) MarkLessonComplete(c *fiber.Ctx) error {
	userID, err := pc.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	lessonID, err := response.GetUUIDParam(c, "lesson_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_lesson_id")
	}

	err = pc.service.MarkLessonComplete(context.Background(), userID, lessonID)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid")
	}

	return response.OK(c, "Lession Successfully Mark As Completed", nil)
}

func (pc *ProgressController) MarkLessonIncomplete(c *fiber.Ctx) error {
	userID, err := pc.jwtManager.GetUserIDFromToken(c)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_user_id")
	}

	lessonID, err := response.GetUUIDParam(c, "lesson_id")
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid_lesson_id")
	}

	err = pc.service.MarkLessonIncomplete(context.Background(), userID, lessonID)
	if err != nil {
		return response.BadRequest(c, err.Error(), "invalid")
	}

	return response.OK(c, "Lession Successfully Mark As Incompleted", nil)
}
