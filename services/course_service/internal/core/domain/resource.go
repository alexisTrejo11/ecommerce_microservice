package domain

import (
	"strings"
	"time"

	customErrors "github.com/alexisTrejo11/ecommerce_microservice/course-service/internal/core/application/errors"
	"github.com/google/uuid"
)

type Resource struct {
	id        uuid.UUID
	lessonID  uuid.UUID
	title     string
	resType   ResourceType
	url       string
	createdAt time.Time
	updatedAt time.Time
}

type ResourceType string

const (
	PDF    ResourceType = "PDF"
	SLIDES ResourceType = "SLIDES"
	LINK   ResourceType = "LINK"
	CODE   ResourceType = "CODE"
	OTHER  ResourceType = "OTHER"
)

var allowedResourceTypes = []ResourceType{PDF, SLIDES, LINK, CODE, OTHER}

func NewResource(
	lessonID uuid.UUID,
	title string,
	resType ResourceType,
	url string,
) (*Resource, error) {
	if strings.TrimSpace(title) == "" {
		return nil, customErrors.ErrResourceTitleRequired
	}
	if strings.TrimSpace(url) == "" {
		return nil, customErrors.ErrResourceURLRequired
	}
	if !isValidResourceType(resType) {
		return nil, customErrors.ErrResourceInvalidType
	}

	now := time.Now()
	return &Resource{
		id:        uuid.New(),
		lessonID:  lessonID,
		title:     title,
		resType:   resType,
		url:       url,
		createdAt: now,
		updatedAt: now,
	}, nil
}

func NewResourceFromModel(
	id uuid.UUID,
	lessonId uuid.UUID,
	title string,
	resType ResourceType,
	url string,
	createdAt time.Time,
	updatedAt time.Time,
) *Resource {
	return &Resource{
		id:        id,
		lessonID:  lessonId,
		title:     title,
		resType:   resType,
		url:       url,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (r *Resource) ID() uuid.UUID        { return r.id }
func (r *Resource) LessonID() uuid.UUID  { return r.lessonID }
func (r *Resource) Title() string        { return r.title }
func (r *Resource) Type() ResourceType   { return r.resType }
func (r *Resource) URL() string          { return r.url }
func (r *Resource) CreatedAt() time.Time { return r.createdAt }
func (r *Resource) UpdatedAt() time.Time { return r.updatedAt }

func (r *Resource) UpdateInfo(newTitle, newURL string, newType ResourceType) error {
	if strings.TrimSpace(newTitle) == "" {
		return customErrors.ErrResourceTitleRequired
	}

	if newURL == "" {
		return customErrors.ErrResourceURLRequired
	}

	if !isValidResourceType(newType) {
		return customErrors.ErrResourceInvalidType
	}

	r.title = newTitle
	r.url = newURL
	r.resType = newType
	r.updatedAt = time.Now()

	return nil
}

func isValidResourceType(rt ResourceType) bool {
	for _, t := range allowedResourceTypes {
		if t == rt {
			return true
		}
	}
	return false
}
