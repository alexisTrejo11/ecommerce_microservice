package domain

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	id        uuid.UUID
	title     string
	videoURL  string
	content   string
	moduleId  uuid.UUID
	resources []Resource
	duration  int
	order     int
	isPreview bool
	createdAt time.Time
	updatedAt time.Time
}

const maxLessonTitleLength = 100
const maxLessonDuration = 180
const maxLimitOfResources = 5

func NewLesson(
	title, videoURL, content string,
	moduleId uuid.UUID,
	duration int,
	order int,
	isPreview bool,
) (*Lesson, error) {
	now := time.Now()
	lesson := &Lesson{
		id:        uuid.New(),
		title:     title,
		videoURL:  videoURL,
		content:   content,
		moduleId:  moduleId,
		resources: []Resource{},
		duration:  duration,
		order:     order,
		isPreview: isPreview,
		createdAt: now,
		updatedAt: now,
	}

	err := lesson.validateDuration()
	if err != nil {
		return nil, err
	}
	err = lesson.validateTitle()
	if err != nil {
		return nil, err
	}

	return lesson, nil
}

func NewLessonFromModel(
	id uuid.UUID,
	title string,
	videoURL string,
	content string,
	moduleID uuid.UUID,
	duration int,
	order int,
	isPreview bool,
	createdAt, updatedAt time.Time,
) *Lesson {
	return &Lesson{
		id:        id,
		title:     title,
		videoURL:  videoURL,
		content:   content,
		moduleId:  moduleID,
		duration:  duration,
		order:     order,
		isPreview: isPreview,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (l *Lesson) ID() uuid.UUID         { return l.id }
func (l *Lesson) Title() string         { return l.title }
func (l *Lesson) VideoURL() string      { return l.videoURL }
func (l *Lesson) Content() string       { return l.content }
func (l *Lesson) ModuleID() uuid.UUID   { return l.moduleId }
func (l *Lesson) Resources() []Resource { return l.resources }
func (l *Lesson) Duration() int         { return l.duration }
func (l *Lesson) Order() int            { return l.order }
func (l *Lesson) IsPreview() bool       { return l.isPreview }
func (l *Lesson) CreatedAt() time.Time  { return l.createdAt }
func (l *Lesson) UpdatedAt() time.Time  { return l.updatedAt }

func (l *Lesson) AddResource(resource Resource) error {
	if (len(l.resources) + 1) >= maxLimitOfResources {
		return errors.New("max limit of resources per lesson reachers")
	}

	l.resources = append(l.resources, resource)
	l.updatedAt = time.Now()

	return nil
}

func (l *Lesson) UpdateContent(title, content, videoURL string, duration int, isPreview bool) error {
	if strings.TrimSpace(title) == "" {
		return errors.New("title is required")
	}
	if len(title) > maxLessonTitleLength {
		return errors.New("title exceeds maximum length")
	}
	if duration < 1 || duration > maxLessonDuration {
		return errors.New("invalid lesson duration")
	}

	l.title = title
	l.content = content
	l.videoURL = videoURL
	l.duration = duration
	l.isPreview = isPreview
	l.updatedAt = time.Now()
	return nil
}

func (l *Lesson) validateTitle() error {
	if strings.TrimSpace(l.title) == "" {
		return errors.New("lesson title is required")
	}
	if len(l.title) > maxLessonTitleLength {
		return errors.New("lesson title exceeds maximum length")
	}

	return nil
}

func (l *Lesson) validateDuration() error {
	if l.duration < 1 || l.duration > maxLessonDuration {
		return errors.New("invalid lesson duration")
	}

	return nil
}
