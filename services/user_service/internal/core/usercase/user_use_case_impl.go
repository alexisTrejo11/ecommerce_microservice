package usecases

import (
	"context"
	"errors"
	"fmt"

	repository "github.com/alexisTrejo11/ecommerce_microservice/internal/adapters/output"

	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/domain/entities"
	"github.com/alexisTrejo11/ecommerce_microservice/internal/core/ports/input"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUseCaseImpl(userRepository repository.UserRepository) input.UserUseCase {
	return &UserUseCaseImpl{userRepository: userRepository}
}

func (uuc *UserUseCaseImpl) GetUser(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	return uuc.userRepository.FindByID(ctx, id)
}

func (uuc *UserUseCaseImpl) UpdateUser(ctx context.Context, user *entities.User) error {
	return uuc.userRepository.Update(ctx, user)
}

func (uuc *UserUseCaseImpl) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return uuc.userRepository.Delete(ctx, id)
}

func (uuc *UserUseCaseImpl) ChangePassword(
	ctx context.Context,
	id uuid.UUID,
	currentPassword,
	newPassword string,
) error {
	user, err := uuc.userRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(currentPassword),
	); err != nil {
		return errors.New("current password is incorrect")
	}

	if err := entities.ValidatePasswordStrength(newPassword); err != nil {
		return fmt.Errorf("invalid new password: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(newPassword),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fmt.Errorf("error hashing password: %w", err)
	}

	user.PasswordHash = string(hashedPassword)
	return uuc.userRepository.Update(ctx, user)
}
