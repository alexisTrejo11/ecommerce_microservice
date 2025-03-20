package config

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

func InitMongoClient() *mongo.Client {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		logger.Fatal().Msg("MONGO_URI is not defined")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Fatal().Err(err).Msg("Error while connecting to MongoDB")
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Fatal().Err(err).Msg("MongoDB doesn't respond")
	}

	logger.Info().Msg("Successfully connected to MongoDB")
	return client
}

type ProgressCollections struct {
	CoursesColl *mongo.Collection
	ModulesColl *mongo.Collection
	LessonsColl *mongo.Collection
}

func createIndexes(ctx context.Context, collection *mongo.Collection, indexes []mongo.IndexModel) error {
	_, err := collection.Indexes().CreateMany(ctx, indexes)
	return err
}

func CreateCollections(client *mongo.Client, dbName string) (*ProgressCollections, error) {
	coursesColl := client.Database(dbName).Collection("courses")
	modulesColl := client.Database(dbName).Collection("modules")
	lessonsColl := client.Database(dbName).Collection("lessons")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	indexModels := []mongo.IndexModel{
		{Keys: bson.D{{Key: "instructor_id", Value: 1}}, Options: options.Index().SetName("idx_instructor_id")},
		{Keys: bson.D{{Key: "category", Value: 1}}, Options: options.Index().SetName("idx_category")},
		{Keys: bson.D{{Key: "level", Value: 1}}, Options: options.Index().SetName("idx_level")},
	}

	moduleIndexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "course_id", Value: 1}}, Options: options.Index().SetName("idx_course_id")},
	}

	lessonIndexes := []mongo.IndexModel{
		{Keys: bson.D{{Key: "module_id", Value: 1}}, Options: options.Index().SetName("idx_module_id")},
	}

	if err := createIndexes(ctx, coursesColl, indexModels); err != nil {
		return nil, fmt.Errorf("failed to create indexes for courses collection: %w", err)
	}

	if err := createIndexes(ctx, modulesColl, moduleIndexes); err != nil {
		return nil, fmt.Errorf("failed to create indexes for modules collection: %w", err)
	}

	if err := createIndexes(ctx, lessonsColl, lessonIndexes); err != nil {
		return nil, fmt.Errorf("failed to create indexes for lessons collection: %w", err)
	}

	logger.Info().Msg("Collections and indexes created successfully")
	return &ProgressCollections{
		CoursesColl: coursesColl,
		ModulesColl: modulesColl,
		LessonsColl: lessonsColl,
	}, nil
}
