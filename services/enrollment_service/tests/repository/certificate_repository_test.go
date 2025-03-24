package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	certificate "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/model"
	repository "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/certificate/repository"
	appErr "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/shared/error"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sqlmock: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("failed to open gorm DB: %v", err)
	}

	return gormDB, mock
}

func TestCreate(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewCertificateRepository(gormDB)

	ctx := context.Background()
	validCert := &certificate.Certificate{
		ID:             uuid.New(),
		EnrollmentID:   uuid.New(),
		IssuedAt:       time.Now(),
		CertificateURL: "http://example.com/cert",
		ExpiresAt:      time.Now().AddDate(1, 0, 0),
	}

	t.Run("Create_Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "certificates"`).
			WithArgs(validCert.ID, validCert.EnrollmentID, validCert.IssuedAt, validCert.CertificateURL, validCert.ExpiresAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Create(ctx, validCert)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Create_DB_Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "certificates"`).
			WithArgs(validCert.ID, validCert.EnrollmentID, validCert.IssuedAt, validCert.CertificateURL, validCert.ExpiresAt).
			WillReturnError(errors.New("db error"))
		mock.ExpectRollback()

		err := repo.Create(ctx, validCert)
		assert.Equal(t, appErr.ErrDB, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetByEnrollment(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewCertificateRepository(gormDB)

	ctx := context.Background()
	enrollmentID := uuid.New()
	foundCert := &certificate.Certificate{
		ID:             uuid.New(),
		EnrollmentID:   enrollmentID,
		IssuedAt:       time.Now(),
		CertificateURL: "http://example.com/cert",
	}

	t.Run("GetByEnrollment_Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "enrollment_id", "issued_at", "certificate_url", "expires_at"}).
			AddRow(foundCert.ID, foundCert.EnrollmentID, foundCert.IssuedAt, foundCert.CertificateURL, foundCert.ExpiresAt)

		mock.ExpectQuery(`SELECT \* FROM "certificates" WHERE enrollment_id = \$1 ORDER BY "certificates"."id" LIMIT \$2`).
			WithArgs(enrollmentID, 1).
			WillReturnRows(rows)

		cert, err := repo.GetByEnrollment(ctx, enrollmentID)
		assert.NoError(t, err)
		assert.Equal(t, foundCert, cert)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetByEnrollment_NotFound", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "certificates" WHERE enrollment_id = \$1 ORDER BY "certificates"."id" LIMIT \$2`).
			WithArgs(enrollmentID, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		cert, err := repo.GetByEnrollment(ctx, enrollmentID)
		assert.Nil(t, cert)
		assert.Equal(t, appErr.ErrCertificateNotFoundDB, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetByEnrollment_DB_Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "certificates" WHERE enrollment_id = \$1`).
			WithArgs(enrollmentID, 1).
			WillReturnError(errors.New("db error"))

		cert, err := repo.GetByEnrollment(ctx, enrollmentID)
		assert.Nil(t, cert)
		assert.Equal(t, appErr.ErrDB, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestGetByUser(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewCertificateRepository(gormDB)

	ctx := context.Background()
	userID := uuid.New()
	foundCerts := []certificate.Certificate{
		{ID: uuid.New(), EnrollmentID: uuid.New(), IssuedAt: time.Now()},
		{ID: uuid.New(), EnrollmentID: uuid.New(), IssuedAt: time.Now()},
	}

	t.Run("GetByUser_Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "enrollment_id", "issued_at", "certificate_url", "expires_at"})
		for _, cert := range foundCerts {
			rows.AddRow(cert.ID, cert.EnrollmentID, cert.IssuedAt, cert.CertificateURL, cert.ExpiresAt)
		}

		mock.ExpectQuery(`SELECT \* FROM "certificates" WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnRows(rows)

		certs, err := repo.GetByUser(ctx, userID)
		assert.NoError(t, err)
		assert.Equal(t, &foundCerts, certs)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetByUser_NoCertificates", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "certificates" WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnRows(sqlmock.NewRows([]string{"id", "enrollment_id", "issued_at", "certificate_url", "expires_at"}))

		certs, err := repo.GetByUser(ctx, userID)
		assert.NoError(t, err)
		assert.Empty(t, *certs)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetByUser_DB_Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "certificates" WHERE user_id = \$1`).
			WithArgs(userID).
			WillReturnError(errors.New("db error"))

		certs, err := repo.GetByUser(ctx, userID)
		assert.Nil(t, certs)
		assert.Equal(t, appErr.ErrDB, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := repository.NewCertificateRepository(gormDB)

	ctx := context.Background()
	cert := &certificate.Certificate{
		ID:             uuid.MustParse("90c158b8-cc75-4db2-b433-451f63c86ab0"),
		EnrollmentID:   uuid.MustParse("534ffde3-8cb8-43bf-8e05-a0ae0707df58"),
		IssuedAt:       time.Date(2025, 3, 24, 11, 4, 32, 0, time.UTC),
		CertificateURL: "http://example.com/cert",
		ExpiresAt:      time.Time{},
	}

	t.Run("Update_Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "certificates"`).
			WithArgs(
				cert.EnrollmentID,   // $1
				cert.IssuedAt,       // $2
				cert.CertificateURL, // $3
				cert.ExpiresAt,      // $4
				cert.ID,             // $5
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Update(ctx, cert)
		assert.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Update_DB_Error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "certificates"`).
			WithArgs(
				cert.EnrollmentID,   // $1
				cert.IssuedAt,       // $2
				cert.CertificateURL, // $3
				cert.ExpiresAt,      // $4
				cert.ID,             // $5
			).
			WillReturnError(errors.New("db error"))
		mock.ExpectRollback()

		err := repo.Update(ctx, cert)
		assert.Equal(t, appErr.ErrDB, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
