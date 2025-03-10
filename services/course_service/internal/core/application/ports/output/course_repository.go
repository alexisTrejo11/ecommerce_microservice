package output

import (
	"github.com/google/uuid"
)

type CourseRepository interface {
	GetCourseById(id string)
	GetCoursesByCategory(category string)
	GetCoursesByInstructorId(instructorId string)
	CourseSearch()
	CreateCourse()
	UpdateCourse()
	DeleteCourse(id uuid.UUID)
}
