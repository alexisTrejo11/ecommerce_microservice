package input

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/shared/dtos"
	"github.com/google/uuid"
)

type LessonUseCase interface {
	GetLessonById(ctx context.Context, id uuid.UUID) (*dtos.LessonDTO, error)
	CreateLesson(ctx context.Context, dto dtos.LessonInsertDTO) (*dtos.LessonDTO, error)
	UpdateLesson(ctx context.Context, id uuid.UUID, dto dtos.LessonInsertDTO) (*dtos.LessonDTO, error)
	DeleteLesson(ctx context.Context, id uuid.UUID) error
}
