package repository

import (
	"log/slog"

	"github.com/highway-to-Golang/02-user-service/internal/domain"
	"github.com/highway-to-Golang/02-user-service/internal/errors"
)

type MockUserRepo struct{}

func NewMockUserRepo() *MockUserRepo {
	return &MockUserRepo{}
}

func (r *MockUserRepo) Save(user domain.User) error {
	slog.Info("MockUserRepo.Save called", "user", user)
	return nil
}

func (r *MockUserRepo) FindByID(id string) (domain.User, error) {
	slog.Info("MockUserRepo.FindByID called", "userID", id)
	return domain.User{}, errors.ErrorWithID(errors.ErrUserNotFound, id)
}

func (r *MockUserRepo) FindAll() []domain.User {
	slog.Info("MockUserRepo.FindAll called")
	return []domain.User{}
}

func (r *MockUserRepo) DeleteByID(id string) error {
	slog.Info("MockUserRepo.DeleteByID called", "userID", id)
	return errors.ErrorWithID(errors.ErrUserNotFound, id)
}
