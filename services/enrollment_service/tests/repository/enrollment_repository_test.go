package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	enrollment "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/model"
	"github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/enrollment/repository"
	appErr "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/error"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate_Success(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewEnrollmentRepository(gormDB)

	enrollment := &enrollment.Enrollment{
		ID:             uuid.New(),
		UserID:         uuid.New(),
		CourseID:       uuid.New(),
		EnrollmentDate: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "enrollments"`).
		WithArgs(
			enrollment.ID,
			enrollment.UserID,
			enrollment.CourseID,
			enrollment.EnrollmentDate,
			sqlmock.AnyArg(), // CompletionDate (nullable)
			sqlmock.AnyArg(), // CompletionStatus (default value)
			sqlmock.AnyArg(), // Progress (default value)
			sqlmock.AnyArg(), // CertificateID (nullable)
			sqlmock.AnyArg(), // LastAccessedAt (nullable)
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Create(context.Background(), enrollment)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByID_Success(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewEnrollmentRepository(gormDB)

	id := uuid.New()
	enrollment := enrollment.Enrollment{
		ID:             id,
		UserID:         uuid.New(),
		CourseID:       uuid.New(),
		EnrollmentDate: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "course_id", "enrollment_date"}).
		AddRow(enrollment.ID, enrollment.UserID, enrollment.CourseID, enrollment.EnrollmentDate)

	mock.ExpectQuery(`SELECT \* FROM "enrollments" WHERE id = \$1 ORDER BY "enrollments"."id" LIMIT \$2`).
		WithArgs(id, 1).
		WillReturnRows(rows)

	result, err := repo.GetByID(context.Background(), id)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByID_NotFound(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewEnrollmentRepository(gormDB)

	id := uuid.New()

	mock.ExpectQuery(`SELECT \* FROM "enrollments" WHERE id = \$1 ORDER BY "enrollments"."id" LIMIT \$2`).
		WithArgs(id, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	result, err := repo.GetByID(context.Background(), id)

	assert.Error(t, err)
	assert.Equal(t, appErr.ErrNotFoundDB, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByUserAndCourse_Success(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewEnrollmentRepository(gormDB)

	userID := uuid.New()
	courseID := uuid.New()
	enrollment := enrollment.Enrollment{
		ID:             uuid.New(),
		UserID:         userID,
		CourseID:       courseID,
		EnrollmentDate: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "user_id", "course_id", "enrollment_date"}).
		AddRow(enrollment.ID, enrollment.UserID, enrollment.CourseID, enrollment.EnrollmentDate)

	mock.ExpectQuery(`SELECT \* FROM "enrollments" WHERE user_id = \$1 AND course_id = \$2 ORDER BY "enrollments"."id" LIMIT \$3`).
		WithArgs(userID, courseID, 1).
		WillReturnRows(rows)

	result, err := repo.GetByUserAndCourse(context.Background(), userID, courseID)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userID, result.UserID)
	assert.Equal(t, courseID, result.CourseID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetByUserAndCourse_NotFound(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewEnrollmentRepository(gormDB)

	userID := uuid.New()
	courseID := uuid.New()

	mock.ExpectQuery(`SELECT \* FROM "enrollments" WHERE user_id = \$1 AND course_id = \$2 ORDER BY "enrollments"."id" LIMIT \$3`).
		WithArgs(userID, courseID, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	result, err := repo.GetByUserAndCourse(context.Background(), userID, courseID)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, appErr.ErrNotFoundDB, err)
	assert.Nil(t, result)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate_Success(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewEnrollmentRepository(gormDB)

	enrollment := &enrollment.Enrollment{
		ID:             uuid.New(),
		UserID:         uuid.New(),
		CourseID:       uuid.New(),
		EnrollmentDate: time.Now(),
	}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "enrollments"`).
		WithArgs(
			enrollment.UserID,
			enrollment.CourseID,
			enrollment.EnrollmentDate,
			sqlmock.AnyArg(), // CompletionDate
			sqlmock.AnyArg(), // CompletionStatus
			sqlmock.AnyArg(), // Progress
			sqlmock.AnyArg(), // CertificateID
			sqlmock.AnyArg(), // LastAccessedAt
			enrollment.ID,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Update(context.Background(), enrollment)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// Check Deletes
func TestDelete_Success(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewEnrollmentRepository(gormDB)

	id := uuid.New()

	rows := sqlmock.NewRows([]string{"id"}).AddRow(id)

	mock.ExpectQuery(`SELECT \* FROM "enrollments" WHERE id = \$1 ORDER BY "enrollments"."id" LIMIT \$2`).
		WithArgs(id, 1).
		WillReturnRows(rows)

	mock.ExpectExec(`DELETE FROM "enrollments" WHERE id = \$1`).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Delete(context.Background(), id)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_NotFound(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewEnrollmentRepository(gormDB)

	id := uuid.New()

	mock.ExpectQuery(`SELECT \* FROM "enrollments" WHERE id = \$1 ORDER BY "enrollments"."id" LIMIT \$2`).
		WithArgs(id, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	err := repo.Delete(context.Background(), id)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, appErr.ErrNotFoundDB, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
