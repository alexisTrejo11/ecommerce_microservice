package domain

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/adapters/output/models"
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

var availableLanguages = []string{"ENGLISH", "SPANISH", "FRENCH", "ITALIAN", "PORTUGUESE"}

type Course struct {
	id              uuid.UUID
	name            string
	slug            string
	description     string
	category        CourseCategory
	level           CourseLevel
	price           float64
	isFree          bool
	rating          float64
	instructorId    uuid.UUID
	thumbnailURL    string
	language        string
	reviewCount     int
	enrollmentCount int
	tags            []string
	publishedAt     *time.Time
	createdAt       time.Time
	updatedAt       time.Time
	modules         []Module
}

func (c *Course) ID() uuid.UUID            { return c.id }
func (c *Course) Name() string             { return c.name }
func (c *Course) Slug() string             { return c.slug }
func (c *Course) Description() string      { return c.description }
func (c *Course) Category() CourseCategory { return c.category }
func (c *Course) Level() CourseLevel       { return c.level }
func (c *Course) Price() float64           { return c.price }
func (c *Course) IsFree() bool             { return c.isFree }
func (c *Course) InstructorID() uuid.UUID  { return c.instructorId }
func (c *Course) ThumbnailURL() string     { return c.thumbnailURL }
func (c *Course) Language() string         { return c.language }
func (c *Course) Rating() float64          { return c.rating }
func (c *Course) ReviewCount() int         { return c.reviewCount }
func (c *Course) Tags() []string           { return c.tags }
func (c *Course) Modules() []Module        { return c.modules }
func (c *Course) PublishedAt() *time.Time  { return c.publishedAt }
func (c *Course) EnrollmentCount() int     { return c.enrollmentCount }
func (c *Course) CreatedAt() time.Time     { return c.createdAt }
func (c *Course) UpdatedAt() time.Time     { return c.updatedAt }

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
	tags []string,
) (*Course, error) {

	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name is required")
	}

	c := &Course{
		id:              uuid.New(),
		name:            name,
		description:     description,
		category:        category,
		level:           level,
		price:           price,
		isFree:          isFree,
		instructorId:    instructorId,
		thumbnailURL:    thumbnailURL,
		language:        language,
		tags:            tags,
		reviewCount:     0,
		enrollmentCount: 0,
		rating:          0,
		publishedAt:     nil,
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
		modules:         []Module{},
	}

	c.generateSlug()

	if !c.validateLanguage() {
		return nil, errors.New("invalid language")
	}

	return c, nil
}

func NewCourseFromModel(model models.CourseModel) *Course {
	c := &Course{
		id:              model.ID,
		name:            model.Title,
		description:     model.Description,
		category:        CourseCategory(model.Category),
		level:           CourseLevel(model.Level),
		price:           model.Price,
		isFree:          model.IsFree,
		rating:          model.Rating,
		slug:            model.Slug,
		instructorId:    model.InstructorID,
		thumbnailURL:    model.ThumbnailURL,
		language:        model.Language,
		reviewCount:     model.ReviewCount,
		enrollmentCount: model.EnrollmentCount,
		tags:            model.Tags,
		publishedAt:     model.PublishedAt,
		createdAt:       model.CreatedAt,
		updatedAt:       model.UpdatedAt,
		modules:         []Module{},
	}
	return c
}

func (c *Course) UpdateInfo(
	name string,
	description string,
	category CourseCategory,
	level CourseLevel,
	price float64,
	isFree bool,
	thumbnailURL string,
	language string,
	tags []string,
) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}

	c.name = name
	c.description = description
	c.category = category
	c.level = level
	c.price = price
	c.isFree = isFree
	c.thumbnailURL = thumbnailURL
	c.language = language
	c.tags = tags
	c.updatedAt = time.Now()

	c.generateSlug()

	if !c.validateLanguage() {
		return errors.New("invalid language")
	}

	return nil
}

func (c *Course) generateSlug() {
	slug := strings.ToLower(c.name)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	reg := regexp.MustCompile("[^a-z0-9-]+")
	slug = reg.ReplaceAllString(slug, "")
	slug = regexp.MustCompile("-+").ReplaceAllString(slug, "-")
	slug = strings.Trim(slug, "-")
	c.slug = slug
}

func (c *Course) validateLanguage() bool {
	for _, lang := range availableLanguages {
		if strings.EqualFold(c.language, lang) {
			return true
		}
	}
	return false
}

func (c *Course) Publish() error {
	if c.publishedAt != nil {
		return errors.New("course already published")
	}
	now := time.Now()
	c.publishedAt = &now
	c.updatedAt = now
	return nil
}

func (c *Course) UpdateRating(newRating float64) {
	c.rating = newRating
	c.updatedAt = time.Now()
}

func (c *Course) EnrollStudent() {
	c.enrollmentCount++
	c.updatedAt = time.Now()
}

func (c *Course) AddModule(module Module) {
	c.modules = append(c.modules, module)
	c.updatedAt = time.Now()
}

func (c *Course) SetModules(modules []Module) {
	c.modules = modules
}
