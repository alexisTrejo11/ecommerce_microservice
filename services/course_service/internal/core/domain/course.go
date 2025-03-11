package domain

import (
	"regexp"
	"strings"
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
	AI                   CourseCategory = "AI"
	Art                  CourseCategory = "ART"
	Marketing            CourseCategory = "MARKETING"
	SocialNetwork        CourseCategory = "SOCIAL_NETWORK"
	Language             CourseCategory = "LANGUAGE"
)

type Course struct {
	Id              uuid.UUID
	Name            string
	Slug            string
	Description     string
	Category        CourseCategory
	Level           CourseLevel
	Price           float64
	IsFree          bool
	Rating          int
	InstructorId    uuid.UUID
	ThumbnailURL    string
	Language        string
	ReviewCount     int
	EnrollmentCount int
	Tags            []string
	PublishedAt     *time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Modules         []Module
}

func (c *Course) GenerateSlug() {
	slug := strings.ToLower(c.Name)

	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")

	reg := regexp.MustCompile("[^a-z0-9-]+")
	slug = reg.ReplaceAllString(slug, "")

	slug = regexp.MustCompile("-+").ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	c.Slug = slug
}

func (c *Course) ValidateLanguage() bool {
	availableLanguages := []string{"ENGLISH", "SPANISH", "FRENCH", "ITALIAN", "PORTUGUESE"}

	for _, lang := range availableLanguages {
		if strings.EqualFold(c.Language, lang) {
			return true
		}
	}
	return false
}
