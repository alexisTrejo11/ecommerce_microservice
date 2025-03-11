package dtos

import "github.com/google/uuid"

type ResourceDTO struct {
	ID       uuid.UUID `json:"id"`
	LessonID uuid.UUID `json:"lesson_id"`
	Title    string    `json:"title"`
	Type     string    `json:"type"`
	URL      string    `json:"url"`
}

type ResourceInsertDTO struct {
	Title    string    `json:"title" binding:"required"`
	LessonID uuid.UUID `json:"lesson_id" binding:"required"`
	Type     string    `json:"type" binding:"required,oneof=PDF SLIDES LINK CODE OTHER"`
	URL      string    `json:"url" binding:"required,url"`
}
