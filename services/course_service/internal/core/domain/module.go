package domain

import (
	"github.com/google/uuid"
)

type Module struct {
	ID       uuid.UUID
	Title    string
	CourseID uuid.UUID
	Order    int
	Lessons  []Lesson
}
