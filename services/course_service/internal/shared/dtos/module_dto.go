package dtos

import "github.com/google/uuid"

type ModuleDTO struct {
	ID       uuid.UUID   `json:"id"`
	CourseID uuid.UUID   `json:"course_id"`
	Title    string      `json:"title"`
	Order    int         `json:"order"`
	Lessons  []LessonDTO `json:"lessons"`
}

type ModuleInsertDTO struct {
	Title    string            `json:"title" validate:"required"`
	CourseID uuid.UUID         `json:"course_id" validate:"required"`
	Order    int               `json:"order" validate:"min=0"`
	Lessons  []LessonInsertDTO `json:"lessons" validate:"dive"`
}
