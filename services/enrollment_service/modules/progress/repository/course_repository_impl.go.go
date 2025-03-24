package repository

import (
	"context"
	"errors"
	"time"

	progress "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/progress/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *MongoDBCourseRepository) Save(ctx context.Context, course *progress.Course) error {
	doc := r.ToCourseDocument(course)

	filter := bson.M{"_id": doc.ID}
	opts := options.Replace().SetUpsert(true)

	_, err := r.collections.CoursesColl.ReplaceOne(ctx, filter, doc, opts)
	return err
}

func (r *MongoDBCourseRepository) FindByID(ctx context.Context, id uuid.UUID) (*progress.Course, error) {
	filter := bson.M{"_id": id.String()}

	var doc CourseDocument
	if err := r.collections.CoursesColl.FindOne(ctx, filter).Decode(&doc); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("course not found")
		}
		return nil, err
	}

	return r.ToCourseEntity(doc)
}

func (r *MongoDBCourseRepository) FindAll(ctx context.Context, limit, offset int64) ([]*progress.Course, int64, error) {

	total, err := r.collections.CoursesColl.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	opts := options.Find().
		SetLimit(limit).
		SetSkip(offset).
		SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collections.CoursesColl.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var documents []CourseDocument
	if err = cursor.All(ctx, &documents); err != nil {
		return nil, 0, err
	}

	courses := make([]*progress.Course, len(documents))
	for i, doc := range documents {
		course, err := r.ToCourseEntity(doc)
		if err != nil {
			return nil, 0, err
		}
		courses[i] = course
	}

	return courses, total, nil
}

func (r *MongoDBCourseRepository) FindByInstructor(ctx context.Context, instructorID uuid.UUID) ([]*progress.Course, error) {
	filter := bson.M{"instructor_id": instructorID.String()}

	cursor, err := r.collections.CoursesColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var documents []CourseDocument
	if err = cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	courses := make([]*progress.Course, len(documents))
	for i, doc := range documents {
		course, err := r.ToCourseEntity(doc)
		if err != nil {
			return nil, err
		}
		courses[i] = course
	}

	return courses, nil
}

func (r *MongoDBCourseRepository) FindByCategory(ctx context.Context, category progress.CourseCategory) ([]*progress.Course, error) {
	filter := bson.M{"category": string(category)}

	cursor, err := r.collections.CoursesColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var documents []CourseDocument
	if err = cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	courses := make([]*progress.Course, len(documents))
	for i, doc := range documents {
		course, err := r.ToCourseEntity(doc)
		if err != nil {
			return nil, err
		}
		courses[i] = course
	}

	return courses, nil
}

func (r *MongoDBCourseRepository) FindByLevel(ctx context.Context, level progress.CourseLevel) ([]*progress.Course, error) {
	filter := bson.M{"level": string(level)}

	cursor, err := r.collections.CoursesColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var documents []CourseDocument
	if err = cursor.All(ctx, &documents); err != nil {
		return nil, err
	}

	courses := make([]*progress.Course, len(documents))
	for i, doc := range documents {
		course, err := r.ToCourseEntity(doc)
		if err != nil {
			return nil, err
		}
		courses[i] = course
	}

	return courses, nil
}

func (r *MongoDBCourseRepository) Update(ctx context.Context, course *progress.Course) error {
	doc := r.ToCourseDocument(course)
	doc.UpdatedAt = time.Now()

	filter := bson.M{"_id": doc.ID}
	update := bson.M{"$set": doc}

	_, err := r.collections.CoursesColl.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoDBCourseRepository) Delete(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id.String()}

	_, err := r.collections.ModulesColl.DeleteMany(ctx, bson.M{"course_id": id.String()})
	if err != nil {
		return err
	}

	_, err = r.collections.CoursesColl.DeleteOne(ctx, filter)
	return err
}

func (r *MongoDBCourseRepository) AddModule(ctx context.Context, courseID uuid.UUID, module ModuleDocument) error {
	_, err := r.collections.ModulesColl.InsertOne(ctx, module)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": courseID.String()}
	update := bson.M{"$push": bson.M{"modules": module}}

	_, err = r.collections.CoursesColl.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoDBCourseRepository) UpdateModule(ctx context.Context, module progress.Module) error {
	filter := bson.M{"_id": module.ID.String()}

	update := bson.M{
		"$set": bson.M{
			"title":        module.Title,
			"order_number": module.OrderNumber,
			"updated_at":   time.Now(),
		},
	}

	_, err := r.collections.ModulesColl.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoDBCourseRepository) DeleteModule(ctx context.Context, moduleID uuid.UUID) error {
	filter := bson.M{"_id": moduleID.String()}

	var moduleDoc ModuleDocument
	if err := r.collections.ModulesColl.FindOne(ctx, filter).Decode(&moduleDoc); err != nil {
		return err
	}

	_, err := r.collections.LessonsColl.DeleteMany(ctx, bson.M{"module_id": moduleID.String()})
	if err != nil {
		return err
	}

	_, err = r.collections.ModulesColl.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	courseFilter := bson.M{"_id": moduleDoc.CourseID}
	courseUpdate := bson.M{"$pull": bson.M{"modules": bson.M{"_id": moduleID.String()}}}

	_, err = r.collections.CoursesColl.UpdateOne(ctx, courseFilter, courseUpdate)
	return err
}

func (r *MongoDBCourseRepository) AddLesson(ctx context.Context, moduleID uuid.UUID, lesson LessonDocument) error {
	_, err := r.collections.LessonsColl.InsertOne(ctx, lesson)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": moduleID.String()}
	update := bson.M{"$push": bson.M{"lessons": lesson}}

	_, err = r.collections.ModulesColl.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoDBCourseRepository) UpdateLesson(ctx context.Context, lesson progress.Lesson) error {
	filter := bson.M{"_id": lesson.ID.String()}

	update := bson.M{
		"$set": bson.M{
			"title":        lesson.Title,
			"content":      lesson.Content,
			"duration":     lesson.Duration,
			"order_number": lesson.OrderNumber,
			"content_type": lesson.ContentType,
			"updated_at":   time.Now(),
		},
	}

	_, err := r.collections.LessonsColl.UpdateOne(ctx, filter, update)
	return err
}

func (r *MongoDBCourseRepository) DeleteLesson(ctx context.Context, lessonID uuid.UUID) error {
	filter := bson.M{"_id": lessonID.String()}

	var lessonDoc LessonDocument
	if err := r.collections.LessonsColl.FindOne(ctx, filter).Decode(&lessonDoc); err != nil {
		return err
	}

	_, err := r.collections.LessonsColl.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	moduleFilter := bson.M{"_id": lessonDoc.ModuleID}
	moduleUpdate := bson.M{"$pull": bson.M{"lessons": bson.M{"_id": lessonID.String()}}}

	_, err = r.collections.ModulesColl.UpdateOne(ctx, moduleFilter, moduleUpdate)
	return err
}

func (r *MongoDBCourseRepository) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}
