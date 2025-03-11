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
	Title    string            `json:"title" binding:"required"`
	CourseID uuid.UUID         `json:"course_id"`
	Order    int               `json:"order" binding:"required,min=0"`
	Lessons  []LessonInsertDTO `json:"lessons" binding:"dive"`
}
