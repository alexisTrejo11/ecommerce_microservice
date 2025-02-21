package usecases

import (
	"context"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/errors"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/output"
	"github.com/google/uuid"
)

type SessionUserCaseImpl struct {
	sessionRepository output.SessionRepository
}

func NewSessionUserCase(sessionRepository output.SessionRepository) input.SessionUseCase {
	return &SessionUserCaseImpl{sessionRepository: sessionRepository}
}

func (suc *SessionUserCaseImpl) GetUserSessions(ctx context.Context, userID uuid.UUID) ([]*entities.Session, error) {
	sessions, err := suc.sessionRepository.FindAllByUserID(ctx, userID)
	if len(sessions) == 0 {
		return make([]*entities.Session, 0), nil
	}

	if err != nil {
		return nil, err
	}

	return sessions, nil
}

func (suc *SessionUserCaseImpl) DeleteSession(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	session, err := suc.sessionRepository.FindByID(ctx, id)
	if session != nil {
		if session.UserID != userID {
			return errors.ErrForbbiden
		}
	}

	if err != nil {
		return err
	}

	if err := suc.sessionRepository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
