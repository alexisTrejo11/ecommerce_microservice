package repository

import (
	"context"
	"time"

	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/config"
	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type LessonDocument struct {
	ID          string    `bson:"_id"`
	ModuleID    string    `bson:"module_id"`
	Title       string    `bson:"title"`
	Content     string    `bson:"content"`
	Duration    uint      `bson:"duration_minutes"`
	OrderNumber uint      `bson:"order_number"`
	ContentType string    `bson:"content_type"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}

type ModuleDocument struct {
	ID          string           `bson:"_id"`
	CourseID    string           `bson:"course_id"`
	Title       string           `bson:"title"`
	OrderNumber uint             `bson:"order_number"`
	Lessons     []LessonDocument `bson:"lessons,omitempty"`
	CreatedAt   time.Time        `bson:"created_at"`
	UpdatedAt   time.Time        `bson:"updated_at"`
}

type CourseDocument struct {
	ID           string           `bson:"_id"`
	Name         string           `bson:"name"`
	Category     string           `bson:"category"`
	Level        string           `bson:"level"`
	InstructorID string           `bson:"instructor_id"`
	ThumbnailURL string           `bson:"thumbnail_url"`
	Language     string           `bson:"language"`
	Modules      []ModuleDocument `bson:"modules,omitempty"`
	CreatedAt    time.Time        `bson:"created_at"`
	UpdatedAt    time.Time        `bson:"updated_at"`
}

type CourseRepository interface {
	Save(ctx context.Context, course *progress.Course) error
	FindByID(ctx context.Context, id uuid.UUID) (*progress.Course, error)
	FindAll(ctx context.Context, limit, offset int64) ([]*progress.Course, int64, error)
	FindByInstructor(ctx context.Context, instructorID uuid.UUID) ([]*progress.Course, error)
	FindByCategory(ctx context.Context, category progress.CourseCategory) ([]*progress.Course, error)
	FindByLevel(ctx context.Context, level progress.CourseLevel) ([]*progress.Course, error)
	Update(ctx context.Context, course *progress.Course) error
	Delete(ctx context.Context, id uuid.UUID) error
	AddModule(ctx context.Context, courseID uuid.UUID, module progress.Module) error
	UpdateModule(ctx context.Context, module progress.Module) error
	DeleteModule(ctx context.Context, moduleID uuid.UUID) error
	AddLesson(ctx context.Context, moduleID uuid.UUID, lesson progress.Lesson) error
	UpdateLesson(ctx context.Context, lesson progress.Lesson) error
	DeleteLesson(ctx context.Context, lessonID uuid.UUID) error
}

type MongoDBCourseRepository struct {
	client      *mongo.Client
	dbName      string
	collections config.ProgressCollections
}

func NewMongoDBCourseRepository(client *mongo.Client, dbName string, collections config.ProgressCollections) *MongoDBCourseRepository {
	return &MongoDBCourseRepository{
		client: client,
		dbName: dbName,
		collections: config.ProgressCollections{
			CoursesColl: collections.CoursesColl,
			ModulesColl: collections.ModulesColl,
			LessonsColl: collections.LessonsColl,
		},
	}
}

func (r *MongoDBCourseRepository) toCourseDocument(course *progress.Course) CourseDocument {
	modules := make([]ModuleDocument, len(course.Modules()))
	for i, m := range course.Modules() {
		lessons := make([]LessonDocument, len(m.Lessons))
		for j, l := range m.Lessons {
			lessons[j] = LessonDocument{
				ID:          l.ID.String(),
				ModuleID:    m.ID.String(),
				Title:       l.Title,
				Content:     l.Content,
				Duration:    l.Duration,
				OrderNumber: l.OrderNumber,
				ContentType: l.ContentType,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
		}

		modules[i] = ModuleDocument{
			ID:          m.ID.String(),
			CourseID:    course.ID().String(),
			Title:       m.Title,
			OrderNumber: m.OrderNumber,
			Lessons:     lessons,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
	}

	return CourseDocument{
		ID:           course.ID().String(),
		Name:         course.Name(),
		Category:     string(course.Category()),
		Level:        string(course.Level()),
		InstructorID: course.InstructorID().String(),
		ThumbnailURL: course.ThumbnailURL(),
		Language:     course.Language(),
		Modules:      modules,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (r *MongoDBCourseRepository) toCourseEntity(doc CourseDocument) (*progress.Course, error) {
	_, err := uuid.Parse(doc.ID)
	if err != nil {
		return nil, err
	}

	instructorID, err := uuid.Parse(doc.InstructorID)
	if err != nil {
		return nil, err
	}

	course, err := progress.NewCourse(
		doc.Name,
		"",
		progress.CourseCategory(doc.Category),
		progress.CourseLevel(doc.Level),
		0,
		false,
		instructorID,
		doc.ThumbnailURL,
		doc.Language,
	)
	if err != nil {
		return nil, err
	}

	modules := make([]progress.Module, len(doc.Modules))
	for i, m := range doc.Modules {
		moduleID, err := uuid.Parse(m.ID)
		if err != nil {
			return nil, err
		}

		courseID, err := uuid.Parse(m.CourseID)
		if err != nil {
			return nil, err
		}

		lessons := make([]progress.Lesson, len(m.Lessons))
		for j, l := range m.Lessons {
			lessonID, err := uuid.Parse(l.ID)
			if err != nil {
				return nil, err
			}

			moduleID, err := uuid.Parse(l.ModuleID)
			if err != nil {
				return nil, err
			}

			lessons[j] = progress.Lesson{
				ID:          lessonID,
				ModuleID:    moduleID,
				Title:       l.Title,
				Content:     l.Content,
				Duration:    l.Duration,
				OrderNumber: l.OrderNumber,
				ContentType: l.ContentType,
			}
		}

		modules[i] = progress.Module{
			ID:          moduleID,
			CourseID:    courseID,
			Title:       m.Title,
			OrderNumber: m.OrderNumber,
			Lessons:     lessons,
		}
	}

	coursePtr := course

	return coursePtr, nil
}
