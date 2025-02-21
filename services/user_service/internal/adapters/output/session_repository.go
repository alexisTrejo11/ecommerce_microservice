package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	models "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output/persistence/mysql/mappers"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
)

type SessionRepositoryImpl struct {
	db             *gorm.DB
	sessionMappers mappers.SessionMappers
}

func NewSessionRepository(db *gorm.DB) output.SessionRepository {
	return &SessionRepositoryImpl{
		db: db,
	}
}

func (r *SessionRepositoryImpl) Create(ctx context.Context, session *entities.Session) error {
	sessionModel := r.sessionMappers.EntityToModel(*session)

	if err := r.db.WithContext(ctx).Create(&sessionModel).Error; err != nil {
		return fmt.Errorf("error creating session: %w", err)
	}

	return nil
}

func (r *SessionRepositoryImpl) FindByID(ctx context.Context, id uuid.UUID) (*entities.Session, error) {
	var sessionModel models.SessionModel

	if err := r.db.WithContext(ctx).First(&sessionModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error finding session by ID: %w", err)
	}

	session := r.sessionMappers.ModelToDomain(sessionModel)
	return session, nil
}

func (r *SessionRepositoryImpl) FindByRefreshToken(ctx context.Context, token string) (*entities.Session, error) {
	var sessionModel models.SessionModel

	if err := r.db.WithContext(ctx).First(&sessionModel, "refresh_token = ?", token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("error finding session by refresh token: %w", err)
	}

	session := r.sessionMappers.ModelToDomain(sessionModel)
	return session, nil
}

func (r *SessionRepositoryImpl) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]*entities.Session, error) {
	var sessionModels []models.SessionModel

	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&sessionModels).Error; err != nil {
		return nil, fmt.Errorf("error finding sessions by user ID: %w", err)
	}

	sessions := make([]*entities.Session, 0, len(sessionModels))
	for _, model := range sessionModels {
		session := r.sessionMappers.ModelToDomain(model)
		sessions = append(sessions, session)
	}

	return sessions, nil
}

func (r *SessionRepositoryImpl) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.SessionModel{}).Error; err != nil {
		return fmt.Errorf("error deleting session: %w", err)
	}
	return nil
}

func (r *SessionRepositoryImpl) DeleteAllByUserID(ctx context.Context, userID uuid.UUID) error {
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&models.SessionModel{}).Error; err != nil {
		return fmt.Errorf("error deleting sessions by user ID: %w", err)
	}
	return nil
}

func (r *SessionRepositoryImpl) DeleteExpired(ctx context.Context) error {
	now := time.Now()

	if err := r.db.WithContext(ctx).Where("expires_at < ?", now).Delete(&models.SessionModel{}).Error; err != nil {
		return fmt.Errorf("error deleting expired sessions: %w", err)
	}
	return nil
}
