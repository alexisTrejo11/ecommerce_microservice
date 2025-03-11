package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseModel struct {
	ID              string        `gorm:"type:char(36);primaryKey"`
	Title           string        `gorm:"size:255;not null" json:"title"`
	Slug            string        `gorm:"size:255;uniqueIndex;not null" json:"slug"`
	Description     string        `gorm:"type:text" json:"description"`
	ThumbnailURL    string        `gorm:"size:512" json:"thumbnail_url"`
	Category        string        `gorm:"size:100" json:"category"`
	Level           string        `gorm:"size:50" json:"level"`
	Language        string        `gorm:"size:50" json:"language"`
	InstructorID    uuid.UUID     `gorm:"type:char(36);not null" json:"instructor_id"`
	Modules         []ModuleModel `gorm:"foreignKey:CourseID;constraint:OnDelete:CASCADE" json:"modules"`
	Tags            string        `gorm:"type:text" json:"tags"`
	Price           float64       `gorm:"type:numeric(10,2)" json:"price"`
	IsFree          bool          `json:"is_free"`
	IsPublished     bool          `json:"is_published"`
	PublishedAt     *time.Time    `json:"published_at,omitempty"`
	EnrollmentCount int           `json:"enrollment_count"`
	Rating          float64       `json:"rating"`
	ReviewCount     int           `json:"review_count"`
	CreatedAt       time.Time     `json:"created_at"`
	UpdatedAt       time.Time     `json:"updated_at"`
}

func (CourseModel) TableName() string {
	return "courses"
}

func (c *CourseModel) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return
}

type ModuleModel struct {
	ID        string        `gorm:"type:char(36);primaryKey"`
	Title     string        `gorm:"size:255;not null" json:"title"`
	Order     int           `gorm:"not null" json:"order"`
	CourseID  uuid.UUID     `gorm:"type:uuid;not null;index" json:"course_id"`
	Lessons   []LessonModel `gorm:"foreignKey:ModuleID;constraint:OnDelete:CASCADE" json:"lessons"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

func (ModuleModel) TableName() string {
	return "modules"
}

func (c *ModuleModel) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return
}

type LessonModel struct {
	ID        string          `gorm:"type:char(36);primaryKey"`
	Title     string          `gorm:"size:255;not null" json:"title"`
	VideoURL  string          `gorm:"size:512" json:"video_url"`
	Content   string          `gorm:"type:text" json:"content"`
	Duration  int             `gorm:"not null" json:"duration"` // in seconds
	Order     int             `gorm:"not null" json:"order"`
	IsPreview bool            `json:"is_preview"`
	ModuleID  uuid.UUID       `gorm:"type:uuid;not null;index" json:"module_id"`
	Resources []ResourceModel `gorm:"foreignKey:LessonID;constraint:OnDelete:CASCADE" json:"resources"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

func (LessonModel) TableName() string {
	return "lessons"
}

func (c *LessonModel) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return
}

type ResourceModel struct {
	ID        string    `gorm:"type:char(36);primaryKey"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Type      string    `gorm:"size:50;not null" json:"type"`
	URL       string    `gorm:"size:512;not null" json:"url"`
	LessonID  uuid.UUID `gorm:"type:uuid;not null;index" json:"lesson_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ResourceModel) TableName() string {
	return "resources"
}

func (c *ResourceModel) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		c.ID = uuid.New().String()
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	return
}
