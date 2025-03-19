package mapper

import (
	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/dtos"
)

func ToCompletedLessonDTO(lesson progress.CompletedLesson) dtos.CompletedLessonDTO {
	return dtos.CompletedLessonDTO{
		ID:           lesson.ID,
		EnrollmentID: lesson.EnrollmentID,
		LessonID:     lesson.LessonID,
		CompletedAt:  fromTimePointer(lesson.CompletedAt),
	}
}

func ToCompletedLesson(lessonDTO dtos.CompletedLessonDTO) progress.CompletedLesson {
	return progress.CompletedLesson{
		ID:           lessonDTO.ID,
		EnrollmentID: lessonDTO.EnrollmentID,
		LessonID:     lessonDTO.LessonID,
		IsCompleted:  !lessonDTO.CompletedAt.IsZero(),
		CompletedAt:  toTimePointer(lessonDTO.CompletedAt),
	}
}

func ToCompletedLessonDTOs(lessons []progress.CompletedLesson) []dtos.CompletedLessonDTO {
	var dtos []dtos.CompletedLessonDTO
	for _, lesson := range lessons {
		dtos = append(dtos, ToCompletedLessonDTO(lesson))
	}
	return dtos
}

func ToCompletedLessons(lessonDTOs []dtos.CompletedLessonDTO) []progress.CompletedLesson {
	var lessons []progress.CompletedLesson
	for _, lessonDTO := range lessonDTOs {
		lessons = append(lessons, ToCompletedLesson(lessonDTO))
	}
	return lessons
}
