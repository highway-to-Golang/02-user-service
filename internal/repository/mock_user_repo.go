package repository

import (
	"fmt"
	"log/slog"

	"user-service/internal/domain"
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
	return domain.User{}, fmt.Errorf("user with ID %s not found", id)
}

func (r *MockUserRepo) FindAll() []domain.User {
	slog.Info("MockUserRepo.FindAll called")
	return []domain.User{}
}

func (r *MockUserRepo) DeleteByID(id string) error {
	slog.Info("MockUserRepo.DeleteByID called", "userID", id)
	return fmt.Errorf("user with ID %s not found", id)
}
