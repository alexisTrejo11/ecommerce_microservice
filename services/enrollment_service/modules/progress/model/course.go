package progress

import (
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
	DesignSoftware       CourseCategory = "DESIGN_SOFTWARE"
	EngineerSoftware     CourseCategory = "ENGINEER_SOFTWARE"
	ArchitectureSoftware CourseCategory = "ARCHITECTURE_SOFTWARE"
	AI                   CourseCategory = "AI"
	Art                  CourseCategory = "ART"
	Marketing            CourseCategory = "MARKETING"
	SocialNetwork        CourseCategory = "SOCIAL_NETWORK"
	Language             CourseCategory = "LANGUAGE"
)

type Course struct {
	id           uuid.UUID
	name         string
	category     CourseCategory
	level        CourseLevel
	instructorId uuid.UUID
	thumbnailURL string
	language     string
	modules      []Module
}

func (c *Course) ID() uuid.UUID            { return c.id }
func (c *Course) Name() string             { return c.name }
func (c *Course) Category() CourseCategory { return c.category }
func (c *Course) Level() CourseLevel       { return c.level }
func (c *Course) InstructorID() uuid.UUID  { return c.instructorId }
func (c *Course) ThumbnailURL() string     { return c.thumbnailURL }
func (c *Course) Language() string         { return c.language }
func (c *Course) Modules() []Module        { return c.modules }

func NewCourse(
	name string,
	description string,
	category CourseCategory,
	level CourseLevel,
	price float64,
	isFree bool,
	instructorId uuid.UUID,
	thumbnailURL string,
	language string,
) (*Course, error) {
	c := &Course{
		id:           uuid.New(),
		name:         name,
		category:     category,
		level:        level,
		instructorId: instructorId,
		thumbnailURL: thumbnailURL,
		language:     language,
		modules:      []Module{},
	}
	return c, nil
}

type Module struct {
	ID          uuid.UUID `gorm:"primarykey"`
	CourseID    uuid.UUID `json:"course_id"`
	Title       string    `json:"title"`
	OrderNumber uint      `json:"order_number"`
	Lessons     []Lesson  `json:"lessons,omitempty"`
}

type Lesson struct {
	ID          uuid.UUID `gorm:"primarykey"`
	ModuleID    uuid.UUID `json:"module_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Duration    uint      `json:"duration_minutes"`
	OrderNumber uint      `json:"order_number"`
	ContentType string    `json:"content_type"`
}
