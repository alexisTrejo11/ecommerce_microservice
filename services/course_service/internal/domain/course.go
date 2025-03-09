package domain

import (
	"time"

	"github.com/google/uuid"
)

type CourseLevel string
type CourseCategory string

const (
	Beginner     CourseLevel = "BEGINNER"
	Intermediate CourseLevel = "INTERMEDIATE"
	Advanced     CourseLevel = "ADVANCED"
)

const (
	Programming          CourseCategory = "PROGRAMMING"
	DesignSoftware       CourseCategory = "DESGIN_SOFTWARE"
	EngineerSoftaware    CourseCategory = "ENGINEER_SOFTWARE"
	ArchitectureSoftware CourseCategory = "ARCHITECTURE_SOFTWARE"
	AI                   CourseCategory = "ARCHITECTURE_SOFTWARE"
	Art                  CourseCategory = "ART"
	Marketing            CourseCategory = "MARKETING"
	SocialNetwork        CourseCategory = "SOCIAL_NETWORK"
	Language             CourseCategory = "Language"
)

type Course struct {
	Id              uuid.UUID
	Name            string
	Description     string
	Category        string
	Level           CourseLevel
	Price           float64
	IsFree          bool
	Rating          int
	InstructorId    uuid.UUID
	ReviewCount     int
	EnrollmentCount int
	PublishedAt     *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Lesson          string
}

type Instructor struct {
	FirstName string
	LastName  string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
