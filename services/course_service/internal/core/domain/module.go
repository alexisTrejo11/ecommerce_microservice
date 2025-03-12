package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

const maxLessonsPerModule = 50

type Module struct {
	id        uuid.UUID
	title     string
	courseID  uuid.UUID
	order     int
	lessons   []Lesson
	createdAt time.Time
	updatedAt time.Time
}

func NewModule(title string, courseID uuid.UUID, order int) (*Module, error) {
	title = strings.TrimSpace(title)
	if title == "" || len(title) > 100 {
		return nil, errors.New("title is required and must be less than 100 characters")
	}

	return &Module{
		id:        uuid.New(),
		title:     title,
		courseID:  courseID,
		order:     order,
		lessons:   []Lesson{},
		createdAt: time.Now(),
		updatedAt: time.Now(),
	}, nil
}

func NewModuleFromModel(
	id uuid.UUID,
	title string,
	courseID uuid.UUID,
	order int,
	lessons []Lesson,
) (*Module, error) {
	if len(title) < 3 || len(title) > 100 {
		return nil, errors.New("title must be between 3 and 100 characters")
	}

	if order < 0 {
		return nil, errors.New("order must be non-negative")
	}

	if len(lessons) > maxLessonsPerModule {
		return nil, fmt.Errorf("a module cannot have more than %d lessons", maxLessonsPerModule)
	}

	return &Module{
		id:       id,
		title:    title,
		courseID: courseID,
		order:    order,
		lessons:  lessons,
	}, nil
}

func (m *Module) ID() uuid.UUID        { return m.id }
func (m *Module) Title() string        { return m.title }
func (m *Module) CourseID() uuid.UUID  { return m.courseID }
func (m *Module) Order() int           { return m.order }
func (m *Module) Lessons() []Lesson    { return m.lessons }
func (m *Module) CreatedAt() time.Time { return m.createdAt }
func (m *Module) UpdatedAt() time.Time { return m.updatedAt }

func (m *Module) Update(title string, order int) error {
	if err := m.updateOrder(order); err != nil {
		return err
	}

	if err := m.updateTitle(title); err != nil {
		return err
	}

	return nil
}

func (m *Module) AddLesson(lesson Lesson) error {
	if len(m.lessons) >= maxLessonsPerModule {
		return errors.New("maximum number of lessons reached for this module")
	}
	m.lessons = append(m.lessons, lesson)
	m.updatedAt = time.Now()
	return nil
}

func (m *Module) SetLessons(lessons []Lesson) error {
	if len(lessons) > maxLessonsPerModule {
		return errors.New("cannot assign more than 50 lessons to a module")
	}
	m.lessons = lessons
	m.updatedAt = time.Now()
	return nil
}

func (m *Module) updateTitle(newTitle string) error {
	newTitle = strings.TrimSpace(newTitle)
	if newTitle == "" || len(newTitle) > 100 {
		return errors.New("title must be non-empty and under 100 characters")
	}
	m.title = newTitle
	m.updatedAt = time.Now()
	return nil
}

func (m *Module) updateOrder(newOrder int) error {
	if newOrder < 0 {
		return errors.New("order must be non-negative")
	}

	m.order = newOrder
	m.updatedAt = time.Now()

	return nil
}
