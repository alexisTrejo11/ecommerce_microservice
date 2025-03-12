package dtos

import "github.com/google/uuid"

// ModuleDTO represents a module associated with a course in the system.
// @Description DTO that contains details of a module within a course, including its lessons.
// @SchemaExample { "id": "a6bfc1f9-0f39-4c6f-b3bc-e9d0c44f3c8a", "course_id": "1c1cdb5c-d6e4-4fb0-9755-f30b9d4fbb8c", "title": "Module 1", "order": 1, "lessons": [{"id": "f2b02b99-4789-4c30-a9b9-b574fbcbd7cd", "title": "Lesson 1", "content": "Introduction to Go"}] }
type ModuleDTO struct {
	// ID is the unique identifier for the module.
	// @example a6bfc1f9-0f39-4c6f-b3bc-e9d0c44f3c8a
	ID uuid.UUID `json:"id"`

	// CourseID is the unique identifier for the course associated with the module.
	// @example 1c1cdb5c-d6e4-4fb0-9755-f30b9d4fbb8c
	CourseID uuid.UUID `json:"course_id"`

	// Title is the name or title of the module.
	// @example Module 1
	Title string `json:"title"`

	// Order is the sequence number of the module in the course.
	// @example 1
	Order int `json:"order"`

	// Lessons is a list of lessons associated with the module.
	// @example [{"id": "f2b02b99-4789-4c30-a9b9-b574fbcbd7cd", "title": "Lesson 1", "content": "Introduction to Go"}]
	Lessons []LessonDTO `json:"lessons"`
}

// ModuleInsertDTO represents the data needed to create a new module within a course.
// @Description DTO used to insert a new module with the required fields, including lessons.
// @SchemaExample { "title": "Module 1", "course_id": "1c1cdb5c-d6e4-4fb0-9755-f30b9d4fbb8c", "order": 1, "lessons": [{"title": "Lesson 1", "content": "Introduction to Go"}] }
type ModuleInsertDTO struct {
	// Title is the name or title of the module.
	// @example Module 1
	Title string `json:"title" validate:"required"`

	// CourseID is the unique identifier for the course associated with the module.
	// @example 1c1cdb5c-d6e4-4fb0-9755-f30b9d4fbb8c
	CourseID uuid.UUID `json:"course_id" validate:"required"`

	// Order is the sequence number of the module in the course.
	// @example 1
	Order int `json:"order" validate:"min=0"`

	// Lessons is a list of lessons associated with the module.
	// @example [{"title": "Lesson 1", "content": "Introduction to Go"}]
	Lessons []LessonInsertDTO `json:"lessons" validate:"dive"`
}
