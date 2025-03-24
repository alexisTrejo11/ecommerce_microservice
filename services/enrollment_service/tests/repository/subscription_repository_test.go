package repository_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	subscription "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/model"
	su_repository "github.com/alexisTrejo11/ecommerce_microservice/enrollment-service/modules/subscription/repository"
)

func TestGetByID(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := su_repository.NewSubscriptionRepository(gormDB)

	ctx := context.Background()
	subscriptionID := uuid.New()
	foundSub := &subscription.Subscription{
		ID:        subscriptionID,
		UserID:    uuid.New(),
		PlanName:  "Basic",
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 1, 0),
		Status:    "ACTIVE",
	}

	t.Run("GetByID_Success", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "user_id", "plan_name", "start_date", "end_date", "status"}).
			AddRow(foundSub.ID, foundSub.UserID, foundSub.PlanName, foundSub.StartDate, foundSub.EndDate, foundSub.Status)

		mock.ExpectQuery(`SELECT \* FROM "subscriptions" WHERE id = \$1 AND "subscriptions"."deleted_at" IS NULL ORDER BY "subscriptions"."id" LIMIT \$2`).
			WithArgs(subscriptionID, 1).
			WillReturnRows(rows)

		sub, err := repo.GetByID(ctx, subscriptionID)
		assert.NoError(t, err)
		assert.Equal(t, foundSub, sub)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("GetByID_NotFound", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "subscriptions" WHERE id = \$1 AND "subscriptions"."deleted_at" IS NULL ORDER BY "subscriptions"."id" LIMIT \$2`).
			WithArgs(subscriptionID, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		sub, err := repo.GetByID(ctx, subscriptionID)
		assert.Nil(t, sub)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
func TestSave(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := su_repository.NewSubscriptionRepository(gormDB)

	ctx := context.Background()
	newSub := &subscription.Subscription{
		UserID:           uuid.New(),
		PlanName:         "Basic",
		StartDate:        time.Now(),
		EndDate:          time.Now().AddDate(0, 1, 0),
		Status:           "ACTIVE",
		SubscriptionType: "MONTHLY",
		PaymentID:        uuid.New(),
	}

	t.Run("Save_Insert_Success", func(t *testing.T) {

		newSub.ID = uuid.Nil

		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "subscriptions"`).
			WithArgs(
				sqlmock.AnyArg(),
				newSub.UserID,
				newSub.PlanName,
				newSub.StartDate,
				newSub.EndDate,
				newSub.Status,
				newSub.SubscriptionType,
				newSub.PaymentID,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				nil,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Save(ctx, newSub)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Save_Update_Success", func(t *testing.T) {

		newSub.ID = uuid.New()

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "subscriptions"`).
			WithArgs(
				newSub.UserID,
				newSub.PlanName,
				newSub.StartDate,
				newSub.EndDate,
				newSub.Status,
				newSub.SubscriptionType,
				newSub.PaymentID,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				nil,
				newSub.ID,
			).
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		err := repo.Save(ctx, newSub)
		assert.NoError(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("Save_DB_Error", func(t *testing.T) {
		newSub.ID = uuid.Nil

		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "subscriptions"`).
			WithArgs(
				sqlmock.AnyArg(),
				newSub.UserID,
				newSub.PlanName,
				newSub.StartDate,
				newSub.EndDate,
				newSub.Status,
				newSub.SubscriptionType,
				newSub.PaymentID,
				sqlmock.AnyArg(),
				sqlmock.AnyArg(),
				nil,
			).
			WillReturnError(errors.New("db error"))
		mock.ExpectRollback()

		err := repo.Save(ctx, newSub)
		assert.Error(t, err)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestSoftDelete(t *testing.T) {
	gormDB, mock := setupMockDB(t)
	repo := su_repository.NewSubscriptionRepository(gormDB)
	ctx := context.Background()
	subscriptionID := uuid.New()

	/*
		t.Run("SoftDelete_Success", func(t *testing.T) {
			mock.ExpectQuery(`SELECT \* FROM "subscriptions" WHERE id = \$1 AND "subscriptions"."deleted_at" IS NULL ORDER BY "subscriptions"."id" LIMIT 1`).
				WithArgs(subscriptionID).
				WillReturnRows(sqlmock.NewRows([]string{"id", "deleted_at"}).
					AddRow(subscriptionID, nil))

			mock.ExpectExec(`UPDATE "subscriptions" SET "deleted_at" = \$1 WHERE "id" = \$2 AND "deleted_at" IS NULL`).
				WithArgs(sqlmock.AnyArg(), subscriptionID).
				WillReturnResult(sqlmock.NewResult(1, 1))

			err := repo.SoftDelete(ctx, subscriptionID)
			assert.NoError(t, err)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	*/
	t.Run("SoftDelete_NotFound", func(t *testing.T) {
		mock.ExpectQuery(`SELECT \* FROM "subscriptions" WHERE id = \$1 AND "subscriptions"."deleted_at" IS NULL ORDER BY "subscriptions"."id" LIMIT \$[12]`).
			WithArgs(subscriptionID, 1).
			WillReturnError(gorm.ErrRecordNotFound)

		err := repo.SoftDelete(ctx, subscriptionID)
		assert.Error(t, err)
		assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
